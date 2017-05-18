package jsm_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/sfazilyesil/jsm"
	"github.com/xeipuuv/gojsonschema"
	"os"
)

var _ = Describe("JSONSchemaMatcher", func() {

	Describe("When actual & expected both are string and loaderType is nil", func() {
		Context("and when JSON resource matches the given schema", func() {
			var (
				matchJSONSchema types.GomegaMatcher
				jsonDoc         = `{"id":1, "name":"product 1"}`
			)
			BeforeEach(func() {
				jsonSchema := `{
			    "type": "object",
			    "required": [
				"id",
				"name"
			    ],
			    "properties": {
				"id": {
				    "type": "integer"
				},
				"name": {
				    "type": "string"
				}
			    }
			}`
				matchJSONSchema = jsm.MatchJSONSchema(jsonSchema)
			})

			It("Should match", func() {
				Expect(jsonDoc).To(matchJSONSchema)
			})
			It("Should return true for Match method", func() {
				success, err := matchJSONSchema.Match(jsonDoc)
				Expect(success).To(BeTrue())
				Expect(err).To(BeNil())
			})
		})

		Context("and when JSON resource does not match the given schema", func() {
			var (
				matchJSONSchema types.GomegaMatcher
				jsonDoc         = `{"ref":1, "name":"product 1"}`
			)
			BeforeEach(func() {
				jsonSchema := `{
			    "type": "object",
			    "required": [
				"id",
				"name"
			    ],
			    "properties": {
				"id": {
				    "type": "integer"
				},
				"name": {
				    "type": "string"
				}
			    }
			}`
				matchJSONSchema = jsm.MatchJSONSchema(jsonSchema)
			})

			It("Should not to match", func() {
				Expect(jsonDoc).NotTo(matchJSONSchema)
			})
			It("Should return false for Match method", func() {
				success, err := matchJSONSchema.Match(jsonDoc)
				Expect(success).To(BeFalse())
				Expect(err).To(BeNil())
			})
			It("Should return a message containing validation error messages for FailureMessage method", func() {
				matchJSONSchema.Match(jsonDoc)
				Expect(matchJSONSchema.FailureMessage(jsonDoc)).To(ContainSubstring("id: id is required"))
			})
		})
	})

	Describe("When expected is JSONLoader and loaderType is nil", func() {
		Context("and when actual matches expected", func() {
			var (
				matchJSONSchema types.GomegaMatcher
				jsonDoc         = `{"id":1, "name":"product 1"}`
			)
			BeforeEach(func() {
				wd, err := os.Getwd()
				if err != nil {
					panic(err)
				}
				jsonSchema := gojsonschema.NewReferenceLoader("file://" + wd + "/schema.json")
				matchJSONSchema = jsm.MatchJSONSchema(jsonSchema)
			})

			It("Should match", func() {
				Expect(jsonDoc).To(matchJSONSchema)
			})
			It("Should return true for Match method", func() {
				success, err := matchJSONSchema.Match(jsonDoc)
				Expect(success).To(BeTrue())
				Expect(err).To(BeNil())
			})
		})
	})

})
