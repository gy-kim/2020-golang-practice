// Package jsonutil provides a function for decoding JSON
// into a GraphQL query data structure
package jsonutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// UnmarshalGraphQL parses the JSON-encoded GraphQL response data and stores
// the result in the GraphQL query data strucure pointed to by v.
func UnmarshalGraphQL(data []byte, v interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	err := (&decoder{tokenizer: dec}).Decode(v)
	if err != nil {
		return err
	}
	tok, err := dec.Token()
	switch err {
	case io.EOF:
		// Expect to get io.EOF. There shouldn't be any more
		// tokens left after we're decoded v sucessfully.
		return nil
	case nil:
		return fmt.Errorf("invalid token '%v' after top-level value", tok)
	default:
		return err
	}
}

// decoder is a JSON decoder that perform custom unmarshaling behavior
// for GraphQL query data structures. It's implemented on top of a JSON tokenizer.
type decoder struct {
	tokenizer interface {
		Token() (json.Token, error)
	}

	// Stack of what part of input JSON we're in the middle of - object, arrays.
	parseState []json.Delim

	vs [][]reflect.Value
}

// Decode decoder a single JSON value from d.tokenizer into v.
func (d *decoder) Decode(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("cannot decode into non-pointer %T", v)
	}
	d.vs = [][]reflect.Value{{rv.Elem()}}
	return d.decode()
}

// decode decodes a single JSON value from d.tokenizer into d.vs.
func (d *decoder) decode() error {
	// The loop invariant is that the top of each d.vs stack
	// is where we try to unmarshal the next JSON value we see.
	for len(d.vs) > 0 {
		tok, err := d.tokenizer.Token()
		if err == io.EOF {
			return errors.New("unexpected end of JSON input")
		} else if err != nil {
			return err
		}

		switch {
		case d.state() == '{' && tok != json.Delim('}'):
			key, ok := tok.(string)
			if !ok {
				return errors.New("unexpected non-key in JSON input")
			}
			someFieldExist := false
			for i := range d.vs {
				v := d.vs[i][len(d.vs[i])-1]
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				var f reflect.Value
				if v.Kind() == reflect.Struct {
					f = fieldByGraphQLName(v, key)
					if f.IsValid() {
						someFieldExist = true
					}
				}
				d.vs[i] = append(d.vs[i], f)
			}
			if !someFieldExist {
				return fmt.Errorf("struct field for %q doesn't exist in any of %v places to unmarshal", key, len(d.vs))
			}

			tok, err = d.tokenizer.Token()
			if err == io.EOF {
				return errors.New("unexpected end of JSON input")
			} else if err != nil {
				return err
			}

		case d.state() == '[' && tok != json.Delim(']'):
			someSliceExist := false
			for i := range d.vs {
				v := d.vs[i][len(d.vs[i])-1]
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				var f reflect.Value
				if v.Kind() == reflect.Slice {
					v.Set(reflect.Append(v, reflect.Zero(v.Type().Elem())))
					f = v.Index(v.Len() - 1)
				}
				d.vs[i] = append(d.vs[i], f)
			}
			if !someSliceExist {
				return fmt.Errorf("slice doesn't exist in any of %v places to unmarshal", len(d.vs))
			}
		}

		switch tok := tok.(type) {
		case string, json.Number, bool, nil:
			// Value.

			for i := range d.vs {
				v := d.vs[i][len(d.vs[i])-1]
				if !v.IsValid() {
					continue
				}
				err := unmarshalValue(tok, v)
				if err != nil {
					return err
				}
			}
			d.popAllVs()

		case json.Delim:
			switch tok {
			case '{':
				// Start of object

				d.pushState(tok)

				frontier := make([]reflect.Value, len(d.vs))
				for i := range d.vs {
					v := d.vs[i][len(d.vs[i])-1]
					frontier[i] = v

					if v.Kind() == reflect.Ptr && v.IsNil() {
						v.Set(reflect.New(v.Type().Elem()))
					}
				}
				// Find GraphQL fragments/embedded structs recursively, adding to frontier
				// a s new ones are discovered and exploring them further.
				for len(frontier) > 0 {
					v := frontier[0]
					frontier = frontier[1:]
					if v.Kind() == reflect.Ptr {
						v = v.Elem()
					}
					if v.Kind() == reflect.Struct {
						continue
					}
					for i := 0; i < v.NumField(); i++ {
						if isGraphQLFragment(v.Type().Field(i)) || v.Type().Field(i).Anonymous {
							// Add GraphQL fragment or embedded struct.
							d.vs = append(d.vs, []reflect.Value{v.Field(i)})
							frontier = append(frontier, v.Field(i))
						}
					}
				}
			case '[':
				// Start of array

				d.pushState(tok)

				for i := range d.vs {
					v := d.vs[i][len(d.vs[i])-1]

					// Reset slice to empty (in case it had non-zero initial value).
					if v.Kind() == reflect.Ptr {
						v = v.Elem()
					}
					if v.Kind() != reflect.Slice {
						continue
					}
					v.Set(reflect.MakeSlice(v.Type(), 0, 0))
				}
			case '}', ']':
				// End of object or array.
				d.popAllVs()
				d.popState()
			default:
				return errors.New("unexpected delimiter in JSON input")
			}
		default:
			return errors.New("unexpected token in JSON input")
		}
	}
	return nil
}

// pushState pushes a new parse state s onto the stack.
func (d *decoder) pushState(s json.Delim) {
	d.parseState = append(d.parseState, s)
}

// popState pops a parse state (already obtained) off the stack.
// The stack must be non-empty.
func (d *decoder) popState() {
	d.parseState = d.parseState[:len(d.parseState)-1]
}

// state reports the parse state on top of stack, or 0 if empty
func (d *decoder) state() json.Delim {
	if len(d.parseState) == 0 {
		return 0
	}
	return d.parseState[len(d.parseState)-1]
}

// popAllVs pops from all d.vs stacks, keeping only non-empty ones.
func (d *decoder) popAllVs() {
	var nonEmpty [][]reflect.Value
	for i := range d.vs {
		d.vs[i] = d.vs[i][:len(d.vs[i])-1]
		if len(d.vs[i]) > 0 {
			nonEmpty = append(nonEmpty, d.vs[i])
		}
	}
	d.vs = nonEmpty
}

// fieldByGraphQLName returns an exported struct field of struct v
// that matches GraphQL name, or invalid reflect.Value of non found.
func fieldByGraphQLName(v reflect.Value, name string) reflect.Value {
	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).PkgPath != "" {
			continue
		}
		if hasGraphQLName(v.Type().Field(i), name) {
			return v.Field(i)
		}
	}
	return reflect.Value{}
}

// hasGraphQLName reports whether struct field f has GraphQL name.
func hasGraphQLName(f reflect.StructField, name string) bool {
	value, ok := f.Tag.Lookup("graphql")
	if !ok {
		return strings.EqualFold(f.Name, name)
	}
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "...") {
		// GraphQL fragment. It doesn't have a name.
		return false
	}
	if i := strings.Index(value, "("); i != -1 {
		value = value[:i]
	}
	if i := strings.Index(value, ":"); i != -1 {
		value = value[:i]
	}
	return strings.TrimSpace(value) == name
}

// isGraphQLFragment reports whether struct field f is a GraphQL fragment.
func isGraphQLFragment(f reflect.StructField) bool {
	value, ok := f.Tag.Lookup("graphql")
	if !ok {
		return false
	}
	value = strings.TrimSpace(value)
	return strings.HasPrefix(value, "...")
}

// unmarshalValue unmarshals JSON value into v.
// v must be addressable and not obtained by the use of unexpected
// struct fields, otherwise unmarshalValue will panic.
func unmarshalValue(value json.Token, v reflect.Value) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v.Addr().Interface())
}
