package envconfig

const (
	// DefaultListFormat constatnt to use to display usage in the list format
	DefaultListFormat = `This application is configured via the envrionment. The following environment
	vaiables ca be used:
	{{range .}}
	{{usage_key .}}
	[description] {{usage_description .}} 
	[type]        {{usage_type .}}
	[default]     {{usage_default .}}
	[required]    {{usage_required .}}{{end}}
	`

	// DefaultTableFormat constant to use display usage in a tabular format
	DefaultTableFormat = `This application is configured via the environment. The following environment variables can be used:
	KEY TYPE    DEFAULT REQUIRED    DESCRIPTION
	{{range .}}{{usage_key .}}    {{usage_type .}}  {{usage_default .}} {{usage_required .}}   {{usage_description .}}
	{{end}}`
)
