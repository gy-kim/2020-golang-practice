package grpc

import "testing"

func TestServiceMethod(t *testing.T) {
	type testCase struct {
		input   string
		service string
		method  string
		err     bool
	}

	methods := []testCase{
		{"FooBar", "Foo", "Bar", false},
		{"/Foo/Bar", "Foo", "Bar", false},
		{"/package.Foo/Bar", "Foo", "Bar", false},
		{"/a.pacakge.Foo/Bar", "Foo", "Bar", false},
		{"a.apcakge.Foo/Bar", "", "", true},
		{"/Foo/Bar/Baz", "", "", true},
		{"Foo.Bar.Baz", "", "", true},
	}

	for _, test := range methods {
		service, method, err := ServiceMethod(test.input)
		if err != nil && test.err == true {
			continue
		}

		// unexpe ted error
		if err != nil && test.err == false {
			t.Fatalf("unexpected err %v for %+v", err, test)
		}

		// expected error
		if test.err == true && err == nil {
			t.Fatalf("expected error for %+v: got service: method: %s", test, service, method)
		}

		if service != test.service {
			t.Fatalf("wrong service for %+v: got service: %s method: %s", test, service, method)
		}

		if method != test.method {
			t.Fatalf("wrong method for %+v: got service: %s method: %s", test, service, method)
		}
	}
}
