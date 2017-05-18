package jsm

import (
	"fmt"
	"github.com/onsi/gomega/types"
	"github.com/xeipuuv/gojsonschema"
)

type LoaderType int

const (
	StringLoader LoaderType = iota
	ReferenceLoader
)

// MatchJSONSchema is a matcher that check if json document matches
// the given JSON Schema draft v0.4.
func MatchJSONSchema(i interface{}, loaderType ...LoaderType) types.GomegaMatcher {
	return &matchJSONSchemaMatcher{
		schemaLoader: getLoader(i, loaderType...),
	}
}

type matchJSONSchemaMatcher struct {
	schemaLoader     gojsonschema.JSONLoader
	validationErrMsg string
}

func (m *matchJSONSchemaMatcher) Match(actual interface{}) (success bool, err error) {
	documentLoader := getLoader(actual)

	result, err := gojsonschema.Validate(m.schemaLoader, documentLoader)
	if err != nil {
		return false, fmt.Errorf("Failed to validate JSON: %s", err.Error())
	}

	if !result.Valid() {
		for _, desc := range result.Errors() {
			m.validationErrMsg += fmt.Sprintf("- %s\n", desc)
		}
	}
	return result.Valid(), nil

}

func (m *matchJSONSchemaMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto conform the JSON schemaLoader of\n\t%#v\nSee the errors:\n%s", actual, m.schemaLoader, m.validationErrMsg)
}

func (m *matchJSONSchemaMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nnot to conform the JSON schemaLoader of\n\t%#v", actual, m.schemaLoader)
}

func getLoader(i interface{}, loaderType ...LoaderType) gojsonschema.JSONLoader {
	jsonLoader, ok := i.(gojsonschema.JSONLoader)
	if ok {
		return jsonLoader
	}

	source, ok := i.(string)
	if !ok {
		panic("MatchJSONSchema expects a string or gojsonschema.JSONLoader")
	}

	loader := StringLoader
	if len(loaderType) > 0 {
		loader = loaderType[0]
	}

	switch loader {
	case StringLoader:
		jsonLoader = gojsonschema.NewStringLoader(source)
	case ReferenceLoader:
		jsonLoader = gojsonschema.NewReferenceLoader(source)
	}

	return jsonLoader
}
