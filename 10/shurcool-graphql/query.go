package graphql

import (
	"encoding/json"
	"reflect"
)

var jsonUnmarshaler = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
