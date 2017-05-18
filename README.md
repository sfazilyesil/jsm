# jsm
Gomega Matcher for JSON Schema Draft v4 - go language. 

### Motivation
Providing a Gomega Matcher originally intended to use with [Ginkgo](http://onsi.github.io/ginkgo) golang BDD test 
framework when validating a JSON document against a [JSON schema](http://json-schema.org/) in tests.

<br/>

### Dependencies
- It depends on the [gojsonschema](https://github.com/xeipuuv/gojsonschema) for validating JSON documents.
- It implements GomegaMatcher interface so depends on [Gomega](https://github.com/onsi/gomega)

<br/>

### Installation
```bash
go get github.com/onsi/gomega/
go get github.com/xeipuuv/gojsonschema
go get github.com/sfazilyesil/jsm
```

<br/>

### Usage
It supports all the loader types that gojsonschema provides.
<br /><br/>

- **Basic Usage Template**
```go
. import "github.com/sfazilyesil/jsm"

Expect(jsonDoc).To(MatchJSONSchema(jsonSchema))
Expect(jsonDocLoader).To(MatchJSONSchema(jsonSchemaLoader))
```

- **String Inputs**
```go
It("should match with the schema", func(){
   jsonDoc := `{"id":1, "name":"product 1"}`
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
   Expect(jsonDoc).To(MatchJSONSchema(jsonSchema))
})
```

- **Reference (URI) Inputs**
```go
It("should match with the schema", func(){
   jsonDoc := `{"id":1, "name":"product 1"}`
   jsonSchema := "file://[project-directory]/lovely-schema.json"
   Expect(jsonDoc).To(MatchJSONSchema(jsonSchema, ReferenceLoader))
})
```

- **Loader Inputs**
<br/>
Asume there is a *schema.json* file exists in the project test directory.
We create a reference loader which loads that schema file and then 
give it to the matcher.

```go
It("should match with the schema", func(){
   jsonDoc := `{"id":1, "name":"product 1"}`
   
   wd, err := os.Getwd() // current directory
   if err != nil {
   		panic(err)
   }
   schemaLoader := gojsonschema.NewReferenceLoader("file://" + wd + "/schema.json")
   
   Expect(jsonDoc).To(MatchJSONSchema(schemaLoader))
})
```

<br/>
---


### Links

- https://github.com/onsi/gomega
- http://onsi.github.io/ginkgo
- https://github.com/xeipuuv/gojsonschema
- http://json-schema.org
