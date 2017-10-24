package options

import "github.com/pkg/errors"

//acronym_style

//AcronymStyle class is "acronym_style" parameter
type AcronymStyle int

const (
	//StyleUnknown is unknown style
	StyleUnknown AcronymStyle = iota
	//StyleMSNaming is "MS naming" style
	StyleMSNaming
	//StyleGuidelines is "guidelines" style
	StyleGuidelines
	//StyleCamelStrict is "camel strict" style
	StyleCamelStrict
	//StyleLiteral is "literal" style
	StyleLiteral
)

var styleMap = map[AcronymStyle]string{
	StyleMSNaming:    "MS naming",
	StyleGuidelines:  "guidelines",
	StyleCamelStrict: "camel strict",
	StyleLiteral:     "literal",
}

//NewAcronymStyleOption returns Casing instance
func NewAcronymStyleOption(value string) (AcronymStyle, error) {
	for s, v := range styleMap {
		if v == value {
			return s, nil
		}
	}
	return StyleUnknown, errors.Wrap(ErrOption, "error in options.NewAcronymStyleOption() function")
}

//Key returns key string
func (s AcronymStyle) Key() string {
	return "acronym_style"
}

//Value returns key string
func (s AcronymStyle) Value() string {
	return s.getAcronymStyleValue()
}

func (s AcronymStyle) getAcronymStyleValue() string {
	if str, ok := styleMap[s]; ok {
		return str
	}
	return ""
}
