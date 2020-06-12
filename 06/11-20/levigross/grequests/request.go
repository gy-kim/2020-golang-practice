package grequests

// RequestOptions is the location that of where the data
type RequestOptions struct {
	
	// Data is a mpa of key values that will eventually convert into the 
	// body if a POST request.
	Data map[string]string

	// Params is a map
	Params map[string]string

	// QueryStruct is a struct that encapsulates a set of
	QueryStruct interface{}
}