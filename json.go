package toolbox

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

//IsCompleteJSON returns true if supplied represent complete JSON
func IsCompleteJSON(candidate string) bool {
	candidate = strings.TrimSpace(candidate)
	if candidate == "" {
		return false
	}
	curlyStart := strings.Count(candidate, "{")
	curlyEnd := strings.Count(candidate, "}")
	squareStart := strings.Count(candidate, "[")
	squareEnd := strings.Count(candidate, "]")
	if !(curlyStart == curlyEnd && squareStart == squareEnd) {
		return false
	}
	var aMap = make(map[string]interface{})
	err := jsonDecoderFactory{}.Create(strings.NewReader(candidate)).Decode(&aMap)
	return err == nil
}

//IsNewLineDelimitedJSON returns true if supplied content is multi line delimited JSON
func IsNewLineDelimitedJSON(candidate string) bool {
	candidate = strings.TrimSpace(candidate)
	if candidate == "" {
		return false
	}
	lines := strings.Split(candidate, "\n")
	if len(lines) == 1 {
		return false
	}
	return IsCompleteJSON(lines[0]) && IsCompleteJSON(lines[1])
}

//JSONToMap converts JSON source into map
func JSONToMap(source interface{}) (map[string]interface{}, error) {
	var reader io.Reader
	switch value := source.(type) {
	case io.Reader:
		reader = value
	case []byte:
		reader = bytes.NewReader(value)
	case string:
		reader = strings.NewReader(value)
	default:
		return nil, fmt.Errorf("unsupported type: %T", source)
	}
	var result = make(map[string]interface{})
	err := jsonDecoderFactory{}.Create(reader).Decode(&result)
	return result, err
}

//AsJSONText converts data structure int text JSON
func AsJSONText(source interface{}) (string, error) {
	if IsStruct(source) || IsMap(source) || IsSlice(source) {
		buf := new(bytes.Buffer)
		err := NewJSONEncoderFactory().Create(buf).Encode(source)
		return buf.String(), err
	}
	return "", fmt.Errorf("unsupported type: %T", source)
}