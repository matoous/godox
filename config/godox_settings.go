package config

type GoDoxSettings struct {
	Format      bool
	Keywords    []string          `mapstructure:"keywords"`
	FormatRules []GoDoxFormatRule `mapstructure:"format-rules"`
}

type GoDoxFormatRule struct {
	Keyword           string
	RegularExpression string
}
