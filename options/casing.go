package options

//Casing class is "casing" parameter
type Casing int

const (
	//CaseUnknown is unknown casing
	CaseUnknown Casing = iota
	//CaseCamel is casing "camel"
	CaseCamel
	//CasePascal is casing "pascal"
	CasePascal
	//CaseUnderscore is casing "lower underscore"
	CaseUnderscore
	//CaseUpperUnderscore is casing "upper underscore"
	CaseUpperUnderscore
	//CaseHyphen is casing "hyphen"
	CaseHyphen
)

var casingMap = map[Casing]string{
	CaseCamel:           "camel",
	CasePascal:          "pascal",
	CaseUnderscore:      "lower underscore",
	CaseUpperUnderscore: "upper underscore",
	CaseHyphen:          "hyphen",
}

//NewCasingOption returns Casing instance
func NewCasingOption(value string) (Casing, bool) {
	for c, v := range casingMap {
		if v == value {
			return c, true
		}
	}
	return CaseUnknown, false
}

//Key returns key string
func (c Casing) Key() string {
	return "casing"
}

//Value returns key string
func (c Casing) Value() string {
	return c.getCasingValue()
}

func (c Casing) getCasingValue() string {
	if str, ok := casingMap[c]; ok {
		return str
	}
	return ""
}
