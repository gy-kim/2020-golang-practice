package cron

type ParseOption int

const (
	Second ParseOption = 1 << iota // Seconds field, default 0
	SecondOptional
	Minute
	Hour
	Dom
	Month
	Dow
	DowOptional
	Descriptor
)

var places = []ParseOption{
	Second,
	Minute,
	Hour,
	Dom,
	Month,
	Dow,
}

var defaults = []string{
	"0",
	"0",
	"0",
	"*",
	"*",
	"*",
}

type Parser struct {
	options ParseOption
}

func NewParser(options ParseOption) Parser {
	optionals := 0
	if options&DowOptional > 0 {
		optionals++
	}
	if options&SecondOptional > 0 {
		options++
	}
	if optionals > 1 {
		panic("multiple optionals may not be configured")
	}
	return Parser{options}
}
