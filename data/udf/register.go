package udf

import "github.com/viant/toolbox/data"

//Register registers defined in this package UDFs
func Register(aMap data.Map) {
	aMap.Put("AsInt", AsInt)
	aMap.Put("AsString", AsString)
	aMap.Put("AsFloat", AsFloat)
	aMap.Put("AsBool", AsBool)
	aMap.Put("AsMap", AsMap)
	aMap.Put("AsData", AsData)
	aMap.Put("AsCollection", AsCollection)
	aMap.Put("Join", Join)
	aMap.Put("Keys", Keys)
	aMap.Put("Values", Values)
	aMap.Put("Length", Length)
	aMap.Put("Len", Length)
	aMap.Put("IndexOf", IndexOf)
	aMap.Put("FormatTime", FormatTime)
	aMap.Put("QueryEscape", QueryEscape)
}