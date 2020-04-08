package meta_schema


const MetaSchemaJSON = "{\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"$id\":\"https://raw.githubusercontent.com/open-rpc/meta-schema/master/schema.json\",\"title\":\"openrpcDocument\",\"type\":\"object\",\"required\":[\"info\",\"methods\",\"openrpc\"],\"additionalProperties\":false,\"patternProperties\":{\"^x-\":{}},\"properties\":{\"openrpc\":{\"title\":\"openrpc\",\"type\":\"string\",\"enum\":[\"1.2.4\",\"1.2.3\",\"1.2.2\",\"1.2.1\",\"1.2.0\",\"1.1.12\",\"1.1.11\",\"1.1.10\",\"1.1.9\",\"1.1.8\",\"1.1.7\",\"1.1.6\",\"1.1.5\",\"1.1.4\",\"1.1.3\",\"1.1.2\",\"1.1.1\",\"1.1.0\",\"1.0.0\",\"1.0.0-rc1\",\"1.0.0-rc0\"]},\"info\":{\"$ref\":\"#/definitions/infoObject\"},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"},\"servers\":{\"title\":\"servers\",\"type\":\"array\",\"additionalItems\":false,\"items\":{\"$ref\":\"#/definitions/serverObject\"}},\"methods\":{\"title\":\"methods\",\"type\":\"array\",\"additionalItems\":false,\"items\":{\"$ref\":\"#/definitions/methodObject\"}},\"components\":{\"title\":\"components\",\"type\":\"object\",\"properties\":{\"schemas\":{\"title\":\"schemaComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/schema\"}}},\"links\":{\"title\":\"linkComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/linkObject\"}}},\"errors\":{\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/errorObject\"}}},\"examples\":{\"title\":\"exampleComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/exampleObject\"}}},\"examplePairings\":{\"title\":\"examplePairingComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/examplePairingObject\"}}},\"contentDescriptors\":{\"title\":\"contentDescriptorComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/contentDescriptorObject\"}}},\"tags\":{\"title\":\"tagComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/tagObject\"}}}}}},\"definitions\":{\"schema\":{\"$ref\":\"https://json-schema.org/draft-07/schema#\"},\"referenceObject\":{\"title\":\"referenceObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"$ref\"],\"properties\":{\"$ref\":{\"title\":\"ref\",\"type\":\"string\",\"format\":\"uri-reference\"}}},\"errorObject\":{\"title\":\"errorObject\",\"type\":\"object\",\"description\":\"Defines an application level error.\",\"additionalProperties\":false,\"required\":[\"code\",\"message\"],\"properties\":{\"code\":{\"title\":\"errorObjectCode\",\"description\":\"A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.\",\"type\":\"integer\"},\"message\":{\"title\":\"errorObjectMessage\",\"description\":\"A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.\",\"type\":\"string\"},\"data\":{\"title\":\"errorObjectData\",\"description\":\"A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).\"}}},\"licenseObject\":{\"title\":\"licenseObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"licenseObjectName\",\"type\":\"string\"},\"url\":{\"title\":\"licenseObjectUrl\",\"type\":\"string\"}},\"patternProperties\":{\"^x-\":{}}},\"contactObject\":{\"title\":\"contactObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"contactObjectName\",\"type\":\"string\"},\"email\":{\"title\":\"contactObjectEmail\",\"type\":\"string\"},\"url\":{\"title\":\"contactObjectUrl\",\"type\":\"string\"}},\"patternProperties\":{\"^x-\":{}}},\"infoObject\":{\"title\":\"infoObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"title\",\"version\"],\"properties\":{\"title\":{\"title\":\"infoObjectProperties\",\"type\":\"string\"},\"description\":{\"title\":\"infoObjectDescription\",\"type\":\"string\"},\"termsOfService\":{\"title\":\"infoObjectTermsOfService\",\"type\":\"string\",\"format\":\"uri\"},\"version\":{\"title\":\"infoObjectVersion\",\"type\":\"string\"},\"contact\":{\"$ref\":\"#/definitions/contactObject\"},\"license\":{\"$ref\":\"#/definitions/licenseObject\"}},\"patternProperties\":{\"^x-\":{}}},\"serverObject\":{\"title\":\"serverObject\",\"type\":\"object\",\"required\":[\"url\"],\"additionalProperties\":false,\"properties\":{\"url\":{\"title\":\"serverObjectUrl\",\"type\":\"string\",\"format\":\"uri\"},\"name\":{\"title\":\"serverObjectName\",\"type\":\"string\"},\"description\":{\"title\":\"serverObjectDescription\",\"type\":\"string\"},\"summary\":{\"title\":\"serverObjectSummary\",\"type\":\"string\"},\"variables\":{\"title\":\"serverObjectVariables\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"title\":\"serverObjectVariable\",\"type\":\"object\",\"required\":[\"default\"],\"properties\":{\"default\":{\"title\":\"serverObjectVariableDefault\",\"type\":\"string\"},\"description\":{\"title\":\"serverObjectVariableDescription\",\"type\":\"string\"},\"enum\":{\"title\":\"serverObjectVariableEnum\",\"type\":\"array\",\"items\":{\"title\":\"serverObjectVariableEnumItem\",\"type\":\"string\"}}}}}}},\"patternProperties\":{\"^x-\":{}}},\"linkObject\":{\"title\":\"linkObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"linkObjectName\",\"type\":\"string\",\"minLength\":1},\"summary\":{\"title\":\"linkObjectSummary\",\"type\":\"string\"},\"method\":{\"title\":\"linkObjectMethod\",\"type\":\"string\"},\"description\":{\"title\":\"linkObjectDescription\",\"type\":\"string\"},\"params\":{\"title\":\"linkObjectParams\"},\"server\":{\"title\":\"linkObjectServer\",\"$ref\":\"#/definitions/serverObject\"}},\"patternProperties\":{\"^x-\":{}}},\"externalDocumentationObject\":{\"title\":\"externalDocumentationObject\",\"type\":\"object\",\"additionalProperties\":false,\"description\":\"information about external documentation\",\"required\":[\"url\"],\"properties\":{\"description\":{\"title\":\"externalDocumentationObjectDescription\",\"type\":\"string\"},\"url\":{\"title\":\"externalDocumentationObjectUrl\",\"type\":\"string\",\"format\":\"uri\"}},\"patternProperties\":{\"^x-\":{}}},\"methodObject\":{\"title\":\"methodObject\",\"type\":\"object\",\"required\":[\"name\",\"result\",\"params\"],\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"methodObjectName\",\"description\":\"The cannonical name for the method. The name MUST be unique within the methods array.\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"methodObjectDescription\",\"description\":\"A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.\",\"type\":\"string\"},\"summary\":{\"title\":\"methodObjectSummary\",\"description\":\"A short summary of what the method does.\",\"type\":\"string\"},\"servers\":{\"title\":\"servers\",\"type\":\"array\",\"additionalItems\":false,\"items\":{\"$ref\":\"#/definitions/serverObject\"}},\"tags\":{\"title\":\"methodObjectTags\",\"type\":\"array\",\"items\":{\"oneOf\":[{\"$ref\":\"#/definitions/tagObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"paramStructure\":{\"title\":\"methodObjectParamStructure\",\"type\":\"string\",\"description\":\"Format the server expects the params. Defaults to 'by-position'.\",\"enum\":[\"by-position\",\"by-name\",\"either\"],\"default\":\"by-position\"},\"params\":{\"title\":\"methodObjectParams\",\"type\":\"array\",\"items\":{\"oneOf\":[{\"$ref\":\"#/definitions/contentDescriptorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"result\":{\"title\":\"methodObjectResult\",\"oneOf\":[{\"$ref\":\"#/definitions/contentDescriptorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"errors\":{\"title\":\"methodObjectErrors\",\"description\":\"Defines an application level error.\",\"type\":\"array\",\"items\":{\"oneOf\":[{\"$ref\":\"#/definitions/errorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"links\":{\"title\":\"methodObjectLinks\",\"type\":\"array\",\"items\":{\"oneOf\":[{\"$ref\":\"#/definitions/linkObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"examples\":{\"title\":\"methodObjectExamples\",\"type\":\"array\",\"items\":{\"oneOf\":[{\"$ref\":\"#/definitions/examplePairingObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"deprecated\":{\"title\":\"methodObjectDeprecated\",\"type\":\"boolean\",\"default\":false},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"}},\"patternProperties\":{\"^x-\":{}}},\"tagObject\":{\"title\":\"tagObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"name\"],\"properties\":{\"name\":{\"title\":\"tagObjectName\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"tagObjectDescription\",\"type\":\"string\"},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"}},\"patternProperties\":{\"^x-\":{}}},\"exampleObject\":{\"title\":\"exampleObject\",\"type\":\"object\",\"required\":[\"name\",\"value\"],\"properties\":{\"summary\":{\"title\":\"exampleObjectSummary\",\"type\":\"string\"},\"value\":{\"title\":\"exampleObjectValue\"},\"description\":{\"title\":\"exampleObjectDescription\",\"type\":\"string\"},\"name\":{\"title\":\"exampleObjectName\",\"type\":\"string\",\"minLength\":1}},\"patternProperties\":{\"^x-\":{}}},\"examplePairingObject\":{\"title\":\"examplePairingObject\",\"type\":\"object\",\"required\":[\"name\",\"params\",\"result\"],\"properties\":{\"name\":{\"title\":\"examplePairingObjectName\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"examplePairingObjectDescription\",\"type\":\"string\"},\"params\":{\"title\":\"examplePairingObjectParams\",\"type\":\"array\",\"items\":{\"oneOf\":[{\"$ref\":\"#/definitions/exampleObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"result\":{\"title\":\"examplePairingObjectresult\",\"oneOf\":[{\"$ref\":\"#/definitions/exampleObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}}},\"contentDescriptorObject\":{\"title\":\"contentDescriptorObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"name\",\"schema\"],\"properties\":{\"name\":{\"title\":\"contentDescriptorObjectName\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"contentDescriptorObjectDescription\",\"type\":\"string\"},\"summary\":{\"title\":\"contentDescriptorObjectSummary\",\"type\":\"string\"},\"schema\":{\"title\":\"contentDescriptorObjectSchema\",\"$ref\":\"#/definitions/schema\"},\"required\":{\"title\":\"contentDescriptorObjectRequired\",\"type\":\"boolean\",\"default\":false},\"deprecated\":{\"title\":\"contentDescriptorObjectDeprecated\",\"type\":\"boolean\",\"default\":false}},\"patternProperties\":{\"^x-\":{}}}}}"

type Openrpc string
const (
	OpenrpcEnum0 Openrpc = "1.2.4"
	OpenrpcEnum1 Openrpc = "1.2.3"
	OpenrpcEnum2 Openrpc = "1.2.2"
	OpenrpcEnum3 Openrpc = "1.2.1"
	OpenrpcEnum4 Openrpc = "1.2.0"
	OpenrpcEnum5 Openrpc = "1.1.12"
	OpenrpcEnum6 Openrpc = "1.1.11"
	OpenrpcEnum7 Openrpc = "1.1.10"
	OpenrpcEnum8 Openrpc = "1.1.9"
	OpenrpcEnum9 Openrpc = "1.1.8"
	OpenrpcEnum10 Openrpc = "1.1.7"
	OpenrpcEnum11 Openrpc = "1.1.6"
	OpenrpcEnum12 Openrpc = "1.1.5"
	OpenrpcEnum13 Openrpc = "1.1.4"
	OpenrpcEnum14 Openrpc = "1.1.3"
	OpenrpcEnum15 Openrpc = "1.1.2"
	OpenrpcEnum16 Openrpc = "1.1.1"
	OpenrpcEnum17 Openrpc = "1.1.0"
	OpenrpcEnum18 Openrpc = "1.0.0"
	OpenrpcEnum19 Openrpc = "1.0.0-rc1"
	OpenrpcEnum20 Openrpc = "1.0.0-rc0"
)
type InfoObjectProperties string
type InfoObjectDescription string
type InfoObjectTermsOfService string
type InfoObjectVersion string
type ContactObjectName string
type ContactObjectEmail string
type ContactObjectUrl string
type ContactObject struct {
	Name  *ContactObjectName  `json:"name,omitempty"`
	Email *ContactObjectEmail `json:"email,omitempty"`
	Url   *ContactObjectUrl   `json:"url,omitempty"`
}
type LicenseObjectName string
type LicenseObjectUrl string
type LicenseObject struct {
	Name *LicenseObjectName `json:"name,omitempty"`
	Url  *LicenseObjectUrl  `json:"url,omitempty"`
}
type InfoObject struct {
	Title          *InfoObjectProperties     `json:"title"`
	Description    *InfoObjectDescription    `json:"description,omitempty"`
	TermsOfService *InfoObjectTermsOfService `json:"termsOfService,omitempty"`
	Version        *InfoObjectVersion        `json:"version"`
	Contact        *ContactObject            `json:"contact,omitempty"`
	License        *LicenseObject            `json:"license,omitempty"`
}
type ExternalDocumentationObjectDescription string
type ExternalDocumentationObjectUrl string
// information about external documentation
type ExternalDocumentationObject struct {
	Description *ExternalDocumentationObjectDescription `json:"description,omitempty"`
	Url         *ExternalDocumentationObjectUrl         `json:"url"`
}
type ServerObjectUrl string
type ServerObjectName string
type ServerObjectDescription string
type ServerObjectSummary string
type ServerObjectVariables map[string]interface{}
type ServerObject struct {
	Url         *ServerObjectUrl         `json:"url"`
	Name        *ServerObjectName        `json:"name,omitempty"`
	Description *ServerObjectDescription `json:"description,omitempty"`
	Summary     *ServerObjectSummary     `json:"summary,omitempty"`
	Variables   *ServerObjectVariables   `json:"variables,omitempty"`
}
type Servers []ServerObject
// The cannonical name for the method. The name MUST be unique within the methods array.
type MethodObjectName string
// A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.
type MethodObjectDescription string
// A short summary of what the method does.
type MethodObjectSummary string
type TagObjectName string
type TagObjectDescription string
type TagObject struct {
	Name         *TagObjectName               `json:"name"`
	Description  *TagObjectDescription        `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocs,omitempty"`
}
type Ref string
type ReferenceObject struct {
	$Ref *Ref `json:"$ref"`
}
type OneOfReferenceObjectTagObjectMTCfXRqB struct {
	TagObject       *TagObject
	ReferenceObject *ReferenceObject
}
type MethodObjectTags []OneOfReferenceObjectTagObjectMTCfXRqB
// Format the server expects the params. Defaults to 'by-position'.
//
// --- Default ---
//
// by-position
type MethodObjectParamStructure string
const (
	MethodObjectParamStructureEnum0 MethodObjectParamStructure = "by-position"
	MethodObjectParamStructureEnum1 MethodObjectParamStructure = "by-name"
	MethodObjectParamStructureEnum2 MethodObjectParamStructure = "either"
)
type ContentDescriptorObjectName string
type ContentDescriptorObjectDescription string
type ContentDescriptorObjectSummary string
type StringVPPt56NS string
type StringNQRYvFt5 string
type StringDoaGddGA string
type AnyL9Fw4VUO interface{}
type BooleanG3T6Tn0M bool
type UnorderedSetOfAnyL9Fw4VUO55Bn0VNb []AnyL9Fw4VUO
type Number0ErlT0It float64
type NumberHo1ClIqD float64
type Integer7Bd9WOt2 int64
type AnyXumYU1GW interface{}
type AllOfAnyXumYU1GWInteger7Bd9WOt20HjcoOji map[string]interface{}
type String3JBlmrip string
type UnorderedSetOfJSONSchemawrpyYBUS []JSONSchema
//
// --- Default ---
//
// true
type AnyOfJSONSchemaUnorderedSetOfJSONSchemawrpyYBUSSidr3R5Q struct {
	JSONSchema                       *JSONSchema
	UnorderedSetOfJSONSchemawrpyYBUS *UnorderedSetOfJSONSchemawrpyYBUS
}
//
// --- Default ---
//
// []
type UnorderedSetOfStringDoaGddGAIEp1G0PF []StringDoaGddGA
type ObjectWrpyYBUS map[string]interface{}
type UnorderedSetOfAnyL9Fw4VUOyeAFYsFq []AnyL9Fw4VUO
type Any17L18NF5 interface{}
type UnorderedSetOfAny17L18NF5VWcS9ROi []Any17L18NF5
type AnyOfAny17L18NF5UnorderedSetOfAny17L18NF5VWcS9ROiRlIv9QVc struct {
	Any17L18NF5                       *Any17L18NF5
	UnorderedSetOfAny17L18NF5VWcS9ROi *UnorderedSetOfAny17L18NF5VWcS9ROi
}
//
// --- Default ---
//
// true
type JSONSchema struct {
	$Id                  *StringVPPt56NS                                            `json:"$id,omitempty"`
	$Schema              *StringNQRYvFt5                                            `json:"$schema,omitempty"`
	$Ref                 *StringVPPt56NS                                            `json:"$ref,omitempty"`
	$Comment             *StringDoaGddGA                                            `json:"$comment,omitempty"`
	Title                *StringDoaGddGA                                            `json:"title,omitempty"`
	Description          *StringDoaGddGA                                            `json:"description,omitempty"`
	Default              *AnyL9Fw4VUO                                               `json:"default,omitempty"`
	ReadOnly             *BooleanG3T6Tn0M                                           `json:"readOnly,omitempty"`
	WriteOnly            *BooleanG3T6Tn0M                                           `json:"writeOnly,omitempty"`
	Examples             *UnorderedSetOfAnyL9Fw4VUO55Bn0VNb                         `json:"examples,omitempty"`
	MultipleOf           *Number0ErlT0It                                            `json:"multipleOf,omitempty"`
	Maximum              *NumberHo1ClIqD                                            `json:"maximum,omitempty"`
	ExclusiveMaximum     *NumberHo1ClIqD                                            `json:"exclusiveMaximum,omitempty"`
	Minimum              *NumberHo1ClIqD                                            `json:"minimum,omitempty"`
	ExclusiveMinimum     *NumberHo1ClIqD                                            `json:"exclusiveMinimum,omitempty"`
	MaxLength            *Integer7Bd9WOt2                                           `json:"maxLength,omitempty"`
	MinLength            *AllOfAnyXumYU1GWInteger7Bd9WOt20HjcoOji                   `json:"minLength,omitempty"`
	Pattern              *String3JBlmrip                                            `json:"pattern,omitempty"`
	AdditionalItems      *JSONSchema                                                `json:"additionalItems,omitempty"`
	Items                *AnyOfJSONSchemaUnorderedSetOfJSONSchemawrpyYBUSSidr3R5Q   `json:"items,omitempty"`
	MaxItems             *Integer7Bd9WOt2                                           `json:"maxItems,omitempty"`
	MinItems             *AllOfAnyXumYU1GWInteger7Bd9WOt20HjcoOji                   `json:"minItems,omitempty"`
	UniqueItems          *BooleanG3T6Tn0M                                           `json:"uniqueItems,omitempty"`
	Contains             *JSONSchema                                                `json:"contains,omitempty"`
	MaxProperties        *Integer7Bd9WOt2                                           `json:"maxProperties,omitempty"`
	MinProperties        *AllOfAnyXumYU1GWInteger7Bd9WOt20HjcoOji                   `json:"minProperties,omitempty"`
	Required             *UnorderedSetOfStringDoaGddGAIEp1G0PF                      `json:"required,omitempty"`
	AdditionalProperties *JSONSchema                                                `json:"additionalProperties,omitempty"`
	Definitions          *ObjectWrpyYBUS                                            `json:"definitions,omitempty"`
	Properties           *ObjectWrpyYBUS                                            `json:"properties,omitempty"`
	PatternProperties    *ObjectWrpyYBUS                                            `json:"patternProperties,omitempty"`
	Dependencies         *ObjectWrpyYBUS                                            `json:"dependencies,omitempty"`
	PropertyNames        *JSONSchema                                                `json:"propertyNames,omitempty"`
	Const                *AnyL9Fw4VUO                                               `json:"const,omitempty"`
	Enum                 *UnorderedSetOfAnyL9Fw4VUOyeAFYsFq                         `json:"enum,omitempty"`
	Type                 *AnyOfAny17L18NF5UnorderedSetOfAny17L18NF5VWcS9ROiRlIv9QVc `json:"type,omitempty"`
	Format               *StringDoaGddGA                                            `json:"format,omitempty"`
	ContentMediaType     *StringDoaGddGA                                            `json:"contentMediaType,omitempty"`
	ContentEncoding      *StringDoaGddGA                                            `json:"contentEncoding,omitempty"`
	If                   *JSONSchema                                                `json:"if,omitempty"`
	Then                 *JSONSchema                                                `json:"then,omitempty"`
	Else                 *JSONSchema                                                `json:"else,omitempty"`
	AllOf                *UnorderedSetOfJSONSchemawrpyYBUS                          `json:"allOf,omitempty"`
	AnyOf                *UnorderedSetOfJSONSchemawrpyYBUS                          `json:"anyOf,omitempty"`
	OneOf                *UnorderedSetOfJSONSchemawrpyYBUS                          `json:"oneOf,omitempty"`
	Not                  *JSONSchema                                                `json:"not,omitempty"`
}
type ContentDescriptorObjectRequired bool
type ContentDescriptorObjectDeprecated bool
type ContentDescriptorObject struct {
	Name        *ContentDescriptorObjectName        `json:"name"`
	Description *ContentDescriptorObjectDescription `json:"description,omitempty"`
	Summary     *ContentDescriptorObjectSummary     `json:"summary,omitempty"`
	Schema      *JSONSchema                         `json:"schema"`
	Required    *ContentDescriptorObjectRequired    `json:"required,omitempty"`
	Deprecated  *ContentDescriptorObjectDeprecated  `json:"deprecated,omitempty"`
}
type OneOfContentDescriptorObjectReferenceObjectI0Ye8PrQ struct {
	ContentDescriptorObject *ContentDescriptorObject
	ReferenceObject         *ReferenceObject
}
type MethodObjectParams []OneOfContentDescriptorObjectReferenceObjectI0Ye8PrQ
type MethodObjectResult struct {
	ContentDescriptorObject *ContentDescriptorObject
	ReferenceObject         *ReferenceObject
}
// A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.
type ErrorObjectCode int64
// A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
type ErrorObjectMessage string
// A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
type ErrorObjectData interface{}
// Defines an application level error.
type ErrorObject struct {
	Code    *ErrorObjectCode    `json:"code"`
	Message *ErrorObjectMessage `json:"message"`
	Data    *ErrorObjectData    `json:"data,omitempty"`
}
type OneOfErrorObjectReferenceObject1KnseVEO struct {
	ErrorObject     *ErrorObject
	ReferenceObject *ReferenceObject
}
// Defines an application level error.
type MethodObjectErrors []OneOfErrorObjectReferenceObject1KnseVEO
type LinkObjectName string
type LinkObjectSummary string
type LinkObjectMethod string
type LinkObjectDescription string
type LinkObjectParams interface{}
type LinkObjectServer struct {
	Url         *ServerObjectUrl         `json:"url"`
	Name        *ServerObjectName        `json:"name,omitempty"`
	Description *ServerObjectDescription `json:"description,omitempty"`
	Summary     *ServerObjectSummary     `json:"summary,omitempty"`
	Variables   *ServerObjectVariables   `json:"variables,omitempty"`
}
type LinkObject struct {
	Name        *LinkObjectName        `json:"name,omitempty"`
	Summary     *LinkObjectSummary     `json:"summary,omitempty"`
	Method      *LinkObjectMethod      `json:"method,omitempty"`
	Description *LinkObjectDescription `json:"description,omitempty"`
	Params      *LinkObjectParams      `json:"params,omitempty"`
	Server      *LinkObjectServer      `json:"server,omitempty"`
}
type OneOfLinkObjectReferenceObjectXyKfUxb0 struct {
	LinkObject      *LinkObject
	ReferenceObject *ReferenceObject
}
type MethodObjectLinks []OneOfLinkObjectReferenceObjectXyKfUxb0
type ExamplePairingObjectName string
type ExamplePairingObjectDescription string
type ExampleObjectSummary string
type ExampleObjectValue interface{}
type ExampleObjectDescription string
type ExampleObjectName string
type ExampleObject struct {
	Summary     *ExampleObjectSummary     `json:"summary,omitempty"`
	Value       *ExampleObjectValue       `json:"value"`
	Description *ExampleObjectDescription `json:"description,omitempty"`
	Name        *ExampleObjectName        `json:"name"`
}
type OneOfExampleObjectReferenceObject5DJ6EmZt struct {
	ExampleObject   *ExampleObject
	ReferenceObject *ReferenceObject
}
type ExamplePairingObjectParams []OneOfExampleObjectReferenceObject5DJ6EmZt
type ExamplePairingObjectresult struct {
	ExampleObject   *ExampleObject
	ReferenceObject *ReferenceObject
}
type ExamplePairingObject struct {
	Name        *ExamplePairingObjectName        `json:"name"`
	Description *ExamplePairingObjectDescription `json:"description,omitempty"`
	Params      *ExamplePairingObjectParams      `json:"params"`
	Result      *ExamplePairingObjectresult      `json:"result"`
}
type OneOfExamplePairingObjectReferenceObjectWEBfRSyK struct {
	ExamplePairingObject *ExamplePairingObject
	ReferenceObject      *ReferenceObject
}
type MethodObjectExamples []OneOfExamplePairingObjectReferenceObjectWEBfRSyK
type MethodObjectDeprecated bool
type MethodObject struct {
	Name           *MethodObjectName            `json:"name"`
	Description    *MethodObjectDescription     `json:"description,omitempty"`
	Summary        *MethodObjectSummary         `json:"summary,omitempty"`
	Servers        *Servers                     `json:"servers,omitempty"`
	Tags           *MethodObjectTags            `json:"tags,omitempty"`
	ParamStructure *MethodObjectParamStructure  `json:"paramStructure,omitempty"`
	Params         *MethodObjectParams          `json:"params"`
	Result         *MethodObjectResult          `json:"result"`
	Errors         *MethodObjectErrors          `json:"errors,omitempty"`
	Links          *MethodObjectLinks           `json:"links,omitempty"`
	Examples       *MethodObjectExamples        `json:"examples,omitempty"`
	Deprecated     *MethodObjectDeprecated      `json:"deprecated,omitempty"`
	ExternalDocs   *ExternalDocumentationObject `json:"externalDocs,omitempty"`
}
type Methods []MethodObject
type SchemaComponents map[string]interface{}
type LinkComponents map[string]interface{}
type ObjectTfFA84LI map[string]interface{}
type ExampleComponents map[string]interface{}
type ExamplePairingComponents map[string]interface{}
type ContentDescriptorComponents map[string]interface{}
type TagComponents map[string]interface{}
type Components struct {
	Schemas            *SchemaComponents            `json:"schemas,omitempty"`
	Links              *LinkComponents              `json:"links,omitempty"`
	Errors             *ObjectTfFA84LI              `json:"errors,omitempty"`
	Examples           *ExampleComponents           `json:"examples,omitempty"`
	ExamplePairings    *ExamplePairingComponents    `json:"examplePairings,omitempty"`
	ContentDescriptors *ContentDescriptorComponents `json:"contentDescriptors,omitempty"`
	Tags               *TagComponents               `json:"tags,omitempty"`
}
type OpenrpcDocument struct {
	Openrpc      *Openrpc                     `json:"openrpc"`
	Info         *InfoObject                  `json:"info"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocs,omitempty"`
	Servers      *Servers                     `json:"servers,omitempty"`
	Methods      *Methods                     `json:"methods"`
	Components   *Components                  `json:"components,omitempty"`
}