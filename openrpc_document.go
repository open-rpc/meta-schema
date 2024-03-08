package openrpc_document


import "encoding/json"
import "errors"
type Openrpc string
const (
	OpenrpcEnum0 Openrpc = "1.3.2"
	OpenrpcEnum1 Openrpc = "1.3.1"
	OpenrpcEnum2 Openrpc = "1.3.0"
	OpenrpcEnum3 Openrpc = "1.2.6"
	OpenrpcEnum4 Openrpc = "1.2.5"
	OpenrpcEnum5 Openrpc = "1.2.4"
	OpenrpcEnum6 Openrpc = "1.2.3"
	OpenrpcEnum7 Openrpc = "1.2.2"
	OpenrpcEnum8 Openrpc = "1.2.1"
	OpenrpcEnum9 Openrpc = "1.2.0"
	OpenrpcEnum10 Openrpc = "1.1.12"
	OpenrpcEnum11 Openrpc = "1.1.11"
	OpenrpcEnum12 Openrpc = "1.1.10"
	OpenrpcEnum13 Openrpc = "1.1.9"
	OpenrpcEnum14 Openrpc = "1.1.8"
	OpenrpcEnum15 Openrpc = "1.1.7"
	OpenrpcEnum16 Openrpc = "1.1.6"
	OpenrpcEnum17 Openrpc = "1.1.5"
	OpenrpcEnum18 Openrpc = "1.1.4"
	OpenrpcEnum19 Openrpc = "1.1.3"
	OpenrpcEnum20 Openrpc = "1.1.2"
	OpenrpcEnum21 Openrpc = "1.1.1"
	OpenrpcEnum22 Openrpc = "1.1.0"
	OpenrpcEnum23 Openrpc = "1.0.0"
	OpenrpcEnum24 Openrpc = "1.0.0-rc1"
	OpenrpcEnum25 Openrpc = "1.0.0-rc0"
)
type InfoObjectProperties string
type InfoObjectDescription string
type InfoObjectTermsOfService string
type InfoObjectVersion string
type ContactObjectName string
type ContactObjectEmail string
type ContactObjectUrl string
type SpecificationExtension interface{}
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
type ServerObjectVariableDefault string
type ServerObjectVariableDescription string
type ServerObjectVariableEnumItem string
type ServerObjectVariableEnum []ServerObjectVariableEnumItem
type ServerObjectVariable struct {
	Default     *ServerObjectVariableDefault     `json:"default"`
	Description *ServerObjectVariableDescription `json:"description,omitempty"`
	Enum        *ServerObjectVariableEnum        `json:"enum,omitempty"`
}
type ServerObjectVariables map[string]interface{}
type ServerObject struct {
	Url         *ServerObjectUrl         `json:"url"`
	Name        *ServerObjectName        `json:"name,omitempty"`
	Description *ServerObjectDescription `json:"description,omitempty"`
	Summary     *ServerObjectSummary     `json:"summary,omitempty"`
	Variables   *ServerObjectVariables   `json:"variables,omitempty"`
}
type AlwaysFalse interface{}
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
	Ref *Ref `json:"$ref"`
}
type TagOrReference struct {
	TagObject       *TagObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *TagOrReference) UnmarshalJSON(bytes []byte) error {
	var myTagObject TagObject
	if err := json.Unmarshal(bytes, &myTagObject); err == nil {
		o.TagObject = &myTagObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o TagOrReference) MarshalJSON() ([]byte, error) {
	if o.TagObject != nil {
		return json.Marshal(o.TagObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type MethodObjectTags []TagOrReference
// Format the server expects the params. Defaults to 'either'.
//
// --- Default ---
//
// either
type MethodObjectParamStructure string
const (
	MethodObjectParamStructureEnum0 MethodObjectParamStructure = "by-position"
	MethodObjectParamStructureEnum1 MethodObjectParamStructure = "by-name"
	MethodObjectParamStructureEnum2 MethodObjectParamStructure = "either"
)
type ContentDescriptorObjectName string
type ContentDescriptorObjectDescription string
type ContentDescriptorObjectSummary string
type Id string
type Schema string
type Comment string
type Title string
type Description string
type AlwaysTrue interface{}
type ReadOnly bool
type Examples []AlwaysTrue
type MultipleOf float64
type Maximum float64
type ExclusiveMaximum float64
type Minimum float64
type ExclusiveMinimum float64
type NonNegativeInteger int64
type NonNegativeIntegerDefaultZero int64
type Pattern string
type SchemaArray []JSONSchema
//
// --- Default ---
//
// true
type Items struct {
	JSONSchema  *JSONSchema
	SchemaArray *SchemaArray
}
func (a *Items) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var myJSONSchema JSONSchema
	if err := json.Unmarshal(bytes, &myJSONSchema); err == nil {
		ok = true
		a.JSONSchema= &myJSONSchema
	}
	var mySchemaArray SchemaArray
	if err := json.Unmarshal(bytes, &mySchemaArray); err == nil {
		ok = true
		a.SchemaArray = &mySchemaArray
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o Items) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.JSONSchema != nil {
		out = append(out, o.JSONSchema)
	}
	if o.SchemaArray != nil {
		out = append(out, o.SchemaArray)
	}
	return json.Marshal(out)
}
type UniqueItems bool
type StringDoaGddGA string
//
// --- Default ---
//
// []
type StringArray []StringDoaGddGA
//
// --- Default ---
//
// {}
type Definitions map[string]interface{}
//
// --- Default ---
//
// {}
type Properties map[string]interface{}
//
// --- Default ---
//
// {}
type PatternProperties map[string]interface{}
type DependenciesSet struct {
	JSONSchema  *JSONSchema
	StringArray *StringArray
}
func (a *DependenciesSet) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var myJSONSchema JSONSchema
	if err := json.Unmarshal(bytes, &myJSONSchema); err == nil {
		ok = true
		a.JSONSchema= &myJSONSchema
	}
	var myStringArray StringArray
	if err := json.Unmarshal(bytes, &myStringArray); err == nil {
		ok = true
		a.StringArray = &myStringArray
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o DependenciesSet) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.JSONSchema!= nil {
		out = append(out, o.JSONSchema)
	}
	if o.StringArray != nil {
		out = append(out, o.StringArray)
	}
	return json.Marshal(out)
}
type Dependencies map[string]interface{}
type Enum []AlwaysTrue
type SimpleTypes interface{}
type ArrayOfSimpleTypes []SimpleTypes
type Type struct {
	SimpleTypes        *SimpleTypes
	ArrayOfSimpleTypes *ArrayOfSimpleTypes
}
func (a *Type) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var mySimpleTypes SimpleTypes
	if err := json.Unmarshal(bytes, &mySimpleTypes); err == nil {
		ok = true
		a.SimpleTypes = &mySimpleTypes
	}
	var myArrayOfSimpleTypes ArrayOfSimpleTypes
	if err := json.Unmarshal(bytes, &myArrayOfSimpleTypes); err == nil {
		ok = true
		a.ArrayOfSimpleTypes = &myArrayOfSimpleTypes
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o Type) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.SimpleTypes != nil {
		out = append(out, o.SimpleTypes)
	}
	if o.ArrayOfSimpleTypes != nil {
		out = append(out, o.ArrayOfSimpleTypes)
	}
	return json.Marshal(out)
}
type Format string
type ContentMediaType string
type ContentEncoding string
type JSONSchemaObject struct {
	Id                   *Id                            `json:"$id,omitempty"`
	Schema               *Schema                        `json:"$schema,omitempty"`
	Ref                  *Ref                           `json:"$ref,omitempty"`
	Comment              *Comment                       `json:"$comment,omitempty"`
	Title                *Title                         `json:"title,omitempty"`
	Description          *Description                   `json:"description,omitempty"`
	Default              *AlwaysTrue                    `json:"default,omitempty"`
	ReadOnly             *ReadOnly                      `json:"readOnly,omitempty"`
	Examples             *Examples                      `json:"examples,omitempty"`
	MultipleOf           *MultipleOf                    `json:"multipleOf,omitempty"`
	Maximum              *Maximum                       `json:"maximum,omitempty"`
	ExclusiveMaximum     *ExclusiveMaximum              `json:"exclusiveMaximum,omitempty"`
	Minimum              *Minimum                       `json:"minimum,omitempty"`
	ExclusiveMinimum     *ExclusiveMinimum              `json:"exclusiveMinimum,omitempty"`
	MaxLength            *NonNegativeInteger            `json:"maxLength,omitempty"`
	MinLength            *NonNegativeIntegerDefaultZero `json:"minLength,omitempty"`
	Pattern              *Pattern                       `json:"pattern,omitempty"`
	AdditionalItems      *JSONSchema                    `json:"additionalItems,omitempty"`
	Items                *Items                         `json:"items,omitempty"`
	MaxItems             *NonNegativeInteger            `json:"maxItems,omitempty"`
	MinItems             *NonNegativeIntegerDefaultZero `json:"minItems,omitempty"`
	UniqueItems          *UniqueItems                   `json:"uniqueItems,omitempty"`
	Contains             *JSONSchema                    `json:"contains,omitempty"`
	MaxProperties        *NonNegativeInteger            `json:"maxProperties,omitempty"`
	MinProperties        *NonNegativeIntegerDefaultZero `json:"minProperties,omitempty"`
	Required             *StringArray                   `json:"required,omitempty"`
	AdditionalProperties *JSONSchema                    `json:"additionalProperties,omitempty"`
	Definitions          *Definitions                   `json:"definitions,omitempty"`
	Properties           *Properties                    `json:"properties,omitempty"`
	PatternProperties    *PatternProperties             `json:"patternProperties,omitempty"`
	Dependencies         *Dependencies                  `json:"dependencies,omitempty"`
	PropertyNames        *JSONSchema                    `json:"propertyNames,omitempty"`
	Const                *AlwaysTrue                    `json:"const,omitempty"`
	Enum                 *Enum                          `json:"enum,omitempty"`
	Type                 *Type                          `json:"type,omitempty"`
	Format               *Format                        `json:"format,omitempty"`
	ContentMediaType     *ContentMediaType              `json:"contentMediaType,omitempty"`
	ContentEncoding      *ContentEncoding               `json:"contentEncoding,omitempty"`
	If                   *JSONSchema                    `json:"if,omitempty"`
	Then                 *JSONSchema                    `json:"then,omitempty"`
	Else                 *JSONSchema                    `json:"else,omitempty"`
	AllOf                *SchemaArray                   `json:"allOf,omitempty"`
	AnyOf                *SchemaArray                   `json:"anyOf,omitempty"`
	OneOf                *SchemaArray                   `json:"oneOf,omitempty"`
	Not                  *JSONSchema                    `json:"not,omitempty"`
}
// Always valid if true. Never valid if false. Is constant.
type JSONSchemaBoolean bool
//
// --- Default ---
//
// {}
type JSONSchema struct {
	JSONSchemaObject  *JSONSchemaObject
	JSONSchemaBoolean *JSONSchemaBoolean
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *JSONSchema) UnmarshalJSON(bytes []byte) error {
	var myJSONSchemaObject JSONSchemaObject
	if err := json.Unmarshal(bytes, &myJSONSchemaObject); err == nil {
		o.JSONSchemaObject = &myJSONSchemaObject
		return nil
	}
	var myJSONSchemaBoolean JSONSchemaBoolean
	if err := json.Unmarshal(bytes, &myJSONSchemaBoolean); err == nil {
		o.JSONSchemaBoolean = &myJSONSchemaBoolean
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o JSONSchema) MarshalJSON() ([]byte, error) {
	if o.JSONSchemaObject != nil {
		return json.Marshal(o.JSONSchemaObject)
	}
	if o.JSONSchemaBoolean != nil {
		return json.Marshal(o.JSONSchemaBoolean)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
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
type ContentDescriptorOrReference struct {
	ContentDescriptorObject *ContentDescriptorObject
	ReferenceObject         *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ContentDescriptorOrReference) UnmarshalJSON(bytes []byte) error {
	var myContentDescriptorObject ContentDescriptorObject
	if err := json.Unmarshal(bytes, &myContentDescriptorObject); err == nil {
		o.ContentDescriptorObject = &myContentDescriptorObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ContentDescriptorOrReference) MarshalJSON() ([]byte, error) {
	if o.ContentDescriptorObject != nil {
		return json.Marshal(o.ContentDescriptorObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type MethodObjectParams []ContentDescriptorOrReference
type MethodObjectResult struct {
	ContentDescriptorObject *ContentDescriptorObject
	ReferenceObject         *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *MethodObjectResult) UnmarshalJSON(bytes []byte) error {
	var myContentDescriptorObject ContentDescriptorObject
	if err := json.Unmarshal(bytes, &myContentDescriptorObject); err == nil {
		o.ContentDescriptorObject = &myContentDescriptorObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o MethodObjectResult) MarshalJSON() ([]byte, error) {
	if o.ContentDescriptorObject != nil {
		return json.Marshal(o.ContentDescriptorObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
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
type ErrorOrReference struct {
	ErrorObject     *ErrorObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ErrorOrReference) UnmarshalJSON(bytes []byte) error {
	var myErrorObject ErrorObject
	if err := json.Unmarshal(bytes, &myErrorObject); err == nil {
		o.ErrorObject = &myErrorObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ErrorOrReference) MarshalJSON() ([]byte, error) {
	if o.ErrorObject != nil {
		return json.Marshal(o.ErrorObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Defines an application level error.
type MethodObjectErrors []ErrorOrReference
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
type LinkOrReference struct {
	LinkObject      *LinkObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *LinkOrReference) UnmarshalJSON(bytes []byte) error {
	var myLinkObject LinkObject
	if err := json.Unmarshal(bytes, &myLinkObject); err == nil {
		o.LinkObject = &myLinkObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o LinkOrReference) MarshalJSON() ([]byte, error) {
	if o.LinkObject != nil {
		return json.Marshal(o.LinkObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type MethodObjectLinks []LinkOrReference
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
type ExampleOrReference struct {
	ExampleObject   *ExampleObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ExampleOrReference) UnmarshalJSON(bytes []byte) error {
	var myExampleObject ExampleObject
	if err := json.Unmarshal(bytes, &myExampleObject); err == nil {
		o.ExampleObject = &myExampleObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ExampleOrReference) MarshalJSON() ([]byte, error) {
	if o.ExampleObject != nil {
		return json.Marshal(o.ExampleObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type ExamplePairingObjectParams []ExampleOrReference
type ExamplePairingObjectResult struct {
	ExampleObject   *ExampleObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ExamplePairingObjectResult) UnmarshalJSON(bytes []byte) error {
	var myExampleObject ExampleObject
	if err := json.Unmarshal(bytes, &myExampleObject); err == nil {
		o.ExampleObject = &myExampleObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ExamplePairingObjectResult) MarshalJSON() ([]byte, error) {
	if o.ExampleObject != nil {
		return json.Marshal(o.ExampleObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type ExamplePairingObject struct {
	Name        *ExamplePairingObjectName        `json:"name"`
	Description *ExamplePairingObjectDescription `json:"description,omitempty"`
	Params      *ExamplePairingObjectParams      `json:"params"`
	Result      *ExamplePairingObjectResult      `json:"result,omitempty"`
}
type ExamplePairingOrReference struct {
	ExamplePairingObject *ExamplePairingObject
	ReferenceObject      *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ExamplePairingOrReference) UnmarshalJSON(bytes []byte) error {
	var myExamplePairingObject ExamplePairingObject
	if err := json.Unmarshal(bytes, &myExamplePairingObject); err == nil {
		o.ExamplePairingObject = &myExamplePairingObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ExamplePairingOrReference) MarshalJSON() ([]byte, error) {
	if o.ExamplePairingObject != nil {
		return json.Marshal(o.ExamplePairingObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type MethodObjectExamples []ExamplePairingOrReference
type MethodObjectDeprecated bool
type MethodObject struct {
	Name           *MethodObjectName            `json:"name"`
	Description    *MethodObjectDescription     `json:"description,omitempty"`
	Summary        *MethodObjectSummary         `json:"summary,omitempty"`
	Servers        *Servers                     `json:"servers,omitempty"`
	Tags           *MethodObjectTags            `json:"tags,omitempty"`
	ParamStructure *MethodObjectParamStructure  `json:"paramStructure,omitempty"`
	Params         *MethodObjectParams          `json:"params"`
	Result         *MethodObjectResult          `json:"result,omitempty"`
	Errors         *MethodObjectErrors          `json:"errors,omitempty"`
	Links          *MethodObjectLinks           `json:"links,omitempty"`
	Examples       *MethodObjectExamples        `json:"examples,omitempty"`
	Deprecated     *MethodObjectDeprecated      `json:"deprecated,omitempty"`
	ExternalDocs   *ExternalDocumentationObject `json:"externalDocs,omitempty"`
}
type MethodOrReference struct {
	MethodObject    *MethodObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *MethodOrReference) UnmarshalJSON(bytes []byte) error {
	var myMethodObject MethodObject
	if err := json.Unmarshal(bytes, &myMethodObject); err == nil {
		o.MethodObject = &myMethodObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o MethodOrReference) MarshalJSON() ([]byte, error) {
	if o.MethodObject != nil {
		return json.Marshal(o.MethodObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type Methods []MethodOrReference
type SchemaComponents map[string]interface{}
type LinkComponents map[string]interface{}
type ErrorComponents map[string]interface{}
type ExampleComponents map[string]interface{}
type ExamplePairingComponents map[string]interface{}
type ContentDescriptorComponents map[string]interface{}
type TagComponents map[string]interface{}
type Components struct {
	Schemas            *SchemaComponents            `json:"schemas,omitempty"`
	Links              *LinkComponents              `json:"links,omitempty"`
	Errors             *ErrorComponents             `json:"errors,omitempty"`
	Examples           *ExampleComponents           `json:"examples,omitempty"`
	ExamplePairings    *ExamplePairingComponents    `json:"examplePairings,omitempty"`
	ContentDescriptors *ContentDescriptorComponents `json:"contentDescriptors,omitempty"`
	Tags               *TagComponents               `json:"tags,omitempty"`
}
// JSON Schema URI (used by some editors)
//
// --- Default ---
//
// https://meta.open-rpc.org/
type MetaSchema string
type OpenrpcDocument struct {
	Openrpc      *Openrpc                     `json:"openrpc"`
	Info         *InfoObject                  `json:"info"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocs,omitempty"`
	Servers      *Servers                     `json:"servers,omitempty"`
	Methods      *Methods                     `json:"methods"`
	Components   *Components                  `json:"components,omitempty"`
	Schema       *MetaSchema                  `json:"$schema,omitempty"`
}

const RawOpenrpc_document = "{\"$schema\":\"https://meta.json-schema.tools/\",\"$id\":\"https://meta.open-rpc.org/\",\"title\":\"openrpcDocument\",\"type\":\"object\",\"required\":[\"info\",\"methods\",\"openrpc\"],\"additionalProperties\":false,\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}},\"properties\":{\"openrpc\":{\"$ref\":\"#/definitions/openrpc\"},\"info\":{\"$ref\":\"#/definitions/infoObject\"},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"},\"servers\":{\"$ref\":\"#/definitions/servers\"},\"methods\":{\"$ref\":\"#/definitions/methods\"},\"components\":{\"$ref\":\"#/definitions/components\"},\"$schema\":{\"$ref\":\"#/definitions/metaSchema\"}},\"definitions\":{\"openrpc\":{\"title\":\"openrpc\",\"type\":\"string\",\"enum\":[\"1.3.2\",\"1.3.1\",\"1.3.0\",\"1.2.6\",\"1.2.5\",\"1.2.4\",\"1.2.3\",\"1.2.2\",\"1.2.1\",\"1.2.0\",\"1.1.12\",\"1.1.11\",\"1.1.10\",\"1.1.9\",\"1.1.8\",\"1.1.7\",\"1.1.6\",\"1.1.5\",\"1.1.4\",\"1.1.3\",\"1.1.2\",\"1.1.1\",\"1.1.0\",\"1.0.0\",\"1.0.0-rc1\",\"1.0.0-rc0\"]},\"infoObjectProperties\":{\"title\":\"infoObjectProperties\",\"type\":\"string\"},\"infoObjectDescription\":{\"title\":\"infoObjectDescription\",\"type\":\"string\"},\"infoObjectTermsOfService\":{\"title\":\"infoObjectTermsOfService\",\"type\":\"string\",\"format\":\"uri\"},\"infoObjectVersion\":{\"title\":\"infoObjectVersion\",\"type\":\"string\"},\"contactObjectName\":{\"title\":\"contactObjectName\",\"type\":\"string\"},\"contactObjectEmail\":{\"title\":\"contactObjectEmail\",\"type\":\"string\"},\"contactObjectUrl\":{\"title\":\"contactObjectUrl\",\"type\":\"string\"},\"specificationExtension\":{\"title\":\"specificationExtension\"},\"contactObject\":{\"title\":\"contactObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"$ref\":\"#/definitions/contactObjectName\"},\"email\":{\"$ref\":\"#/definitions/contactObjectEmail\"},\"url\":{\"$ref\":\"#/definitions/contactObjectUrl\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"licenseObjectName\":{\"title\":\"licenseObjectName\",\"type\":\"string\"},\"licenseObjectUrl\":{\"title\":\"licenseObjectUrl\",\"type\":\"string\"},\"licenseObject\":{\"title\":\"licenseObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"$ref\":\"#/definitions/licenseObjectName\"},\"url\":{\"$ref\":\"#/definitions/licenseObjectUrl\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"infoObject\":{\"title\":\"infoObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"title\",\"version\"],\"properties\":{\"title\":{\"$ref\":\"#/definitions/infoObjectProperties\"},\"description\":{\"$ref\":\"#/definitions/infoObjectDescription\"},\"termsOfService\":{\"$ref\":\"#/definitions/infoObjectTermsOfService\"},\"version\":{\"$ref\":\"#/definitions/infoObjectVersion\"},\"contact\":{\"$ref\":\"#/definitions/contactObject\"},\"license\":{\"$ref\":\"#/definitions/licenseObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"externalDocumentationObjectDescription\":{\"title\":\"externalDocumentationObjectDescription\",\"type\":\"string\"},\"externalDocumentationObjectUrl\":{\"title\":\"externalDocumentationObjectUrl\",\"type\":\"string\",\"format\":\"uri\"},\"externalDocumentationObject\":{\"title\":\"externalDocumentationObject\",\"type\":\"object\",\"additionalProperties\":false,\"description\":\"information about external documentation\",\"required\":[\"url\"],\"properties\":{\"description\":{\"$ref\":\"#/definitions/externalDocumentationObjectDescription\"},\"url\":{\"$ref\":\"#/definitions/externalDocumentationObjectUrl\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"serverObjectUrl\":{\"title\":\"serverObjectUrl\",\"type\":\"string\",\"format\":\"uri\"},\"serverObjectName\":{\"title\":\"serverObjectName\",\"type\":\"string\"},\"serverObjectDescription\":{\"title\":\"serverObjectDescription\",\"type\":\"string\"},\"serverObjectSummary\":{\"title\":\"serverObjectSummary\",\"type\":\"string\"},\"serverObjectVariableDefault\":{\"title\":\"serverObjectVariableDefault\",\"type\":\"string\"},\"serverObjectVariableDescription\":{\"title\":\"serverObjectVariableDescription\",\"type\":\"string\"},\"serverObjectVariableEnumItem\":{\"title\":\"serverObjectVariableEnumItem\",\"type\":\"string\"},\"serverObjectVariableEnum\":{\"title\":\"serverObjectVariableEnum\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/serverObjectVariableEnumItem\"}},\"serverObjectVariable\":{\"title\":\"serverObjectVariable\",\"type\":\"object\",\"required\":[\"default\"],\"properties\":{\"default\":{\"$ref\":\"#/definitions/serverObjectVariableDefault\"},\"description\":{\"$ref\":\"#/definitions/serverObjectVariableDescription\"},\"enum\":{\"$ref\":\"#/definitions/serverObjectVariableEnum\"}}},\"serverObjectVariables\":{\"title\":\"serverObjectVariables\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/serverObjectVariable\"}}},\"serverObject\":{\"title\":\"serverObject\",\"type\":\"object\",\"required\":[\"url\"],\"additionalProperties\":false,\"properties\":{\"url\":{\"$ref\":\"#/definitions/serverObjectUrl\"},\"name\":{\"$ref\":\"#/definitions/serverObjectName\"},\"description\":{\"$ref\":\"#/definitions/serverObjectDescription\"},\"summary\":{\"$ref\":\"#/definitions/serverObjectSummary\"},\"variables\":{\"$ref\":\"#/definitions/serverObjectVariables\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"AlwaysFalse\":false,\"servers\":{\"title\":\"servers\",\"type\":\"array\",\"additionalItems\":false,\"items\":{\"$ref\":\"#/definitions/serverObject\"}},\"methodObjectName\":{\"title\":\"methodObjectName\",\"description\":\"The cannonical name for the method. The name MUST be unique within the methods array.\",\"type\":\"string\",\"minLength\":1},\"methodObjectDescription\":{\"title\":\"methodObjectDescription\",\"description\":\"A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.\",\"type\":\"string\"},\"methodObjectSummary\":{\"title\":\"methodObjectSummary\",\"description\":\"A short summary of what the method does.\",\"type\":\"string\"},\"tagObjectName\":{\"title\":\"tagObjectName\",\"type\":\"string\",\"minLength\":1},\"tagObjectDescription\":{\"title\":\"tagObjectDescription\",\"type\":\"string\"},\"tagObject\":{\"title\":\"tagObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"name\"],\"properties\":{\"name\":{\"$ref\":\"#/definitions/tagObjectName\"},\"description\":{\"$ref\":\"#/definitions/tagObjectDescription\"},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"$ref\":{\"title\":\"$ref\",\"type\":\"string\",\"format\":\"uri-reference\"},\"referenceObject\":{\"title\":\"referenceObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"$ref\"],\"properties\":{\"$ref\":{\"$ref\":\"#/definitions/$ref\"}}},\"tagOrReference\":{\"title\":\"tagOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/tagObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"methodObjectTags\":{\"title\":\"methodObjectTags\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/tagOrReference\"}},\"methodObjectParamStructure\":{\"title\":\"methodObjectParamStructure\",\"type\":\"string\",\"description\":\"Format the server expects the params. Defaults to 'either'.\",\"enum\":[\"by-position\",\"by-name\",\"either\"],\"default\":\"either\"},\"contentDescriptorObjectName\":{\"title\":\"contentDescriptorObjectName\",\"type\":\"string\",\"minLength\":1},\"contentDescriptorObjectDescription\":{\"title\":\"contentDescriptorObjectDescription\",\"type\":\"string\"},\"contentDescriptorObjectSummary\":{\"title\":\"contentDescriptorObjectSummary\",\"type\":\"string\"},\"$id\":{\"title\":\"$id\",\"type\":\"string\",\"format\":\"uri-reference\"},\"$schema\":{\"title\":\"$schema\",\"type\":\"string\",\"format\":\"uri\"},\"$comment\":{\"title\":\"$comment\",\"type\":\"string\"},\"title\":{\"title\":\"title\",\"type\":\"string\"},\"description\":{\"title\":\"description\",\"type\":\"string\"},\"AlwaysTrue\":true,\"readOnly\":{\"title\":\"readOnly\",\"type\":\"boolean\",\"default\":false},\"examples\":{\"title\":\"examples\",\"type\":\"array\",\"items\":true},\"multipleOf\":{\"title\":\"multipleOf\",\"type\":\"number\",\"exclusiveMinimum\":0},\"maximum\":{\"title\":\"maximum\",\"type\":\"number\"},\"exclusiveMaximum\":{\"title\":\"exclusiveMaximum\",\"type\":\"number\"},\"minimum\":{\"title\":\"minimum\",\"type\":\"number\"},\"exclusiveMinimum\":{\"title\":\"exclusiveMinimum\",\"type\":\"number\"},\"nonNegativeInteger\":{\"title\":\"nonNegativeInteger\",\"type\":\"integer\",\"minimum\":0},\"nonNegativeIntegerDefaultZero\":{\"title\":\"nonNegativeIntegerDefaultZero\",\"type\":\"integer\",\"minimum\":0,\"default\":0},\"pattern\":{\"title\":\"pattern\",\"type\":\"string\",\"format\":\"regex\"},\"schemaArray\":{\"title\":\"schemaArray\",\"type\":\"array\",\"minItems\":1,\"items\":{\"$ref\":\"#/definitions/undefined\"}},\"items\":{\"title\":\"items\",\"anyOf\":[{\"$ref\":\"#/definitions/undefined\"},{\"$ref\":\"#/definitions/schemaArray\"}],\"default\":true},\"uniqueItems\":{\"title\":\"uniqueItems\",\"type\":\"boolean\",\"default\":false},\"string_doaGddGA\":{\"type\":\"string\",\"title\":\"string_doaGddGA\"},\"stringArray\":{\"title\":\"stringArray\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/string_doaGddGA\"},\"uniqueItems\":true,\"default\":[]},\"definitions\":{\"title\":\"definitions\",\"type\":\"object\",\"additionalProperties\":{\"$ref\":\"#/definitions/undefined\"},\"default\":{}},\"properties\":{\"title\":\"properties\",\"type\":\"object\",\"additionalProperties\":{\"$ref\":\"#/definitions/undefined\"},\"default\":{}},\"patternProperties\":{\"title\":\"patternProperties\",\"type\":\"object\",\"additionalProperties\":{\"$ref\":\"#/definitions/undefined\"},\"propertyNames\":{\"title\":\"propertyNames\",\"format\":\"regex\"},\"default\":{}},\"dependenciesSet\":{\"title\":\"dependenciesSet\",\"anyOf\":[{\"$ref\":\"#/definitions/undefined\"},{\"$ref\":\"#/definitions/stringArray\"}]},\"dependencies\":{\"title\":\"dependencies\",\"type\":\"object\",\"additionalProperties\":{\"$ref\":\"#/definitions/dependenciesSet\"}},\"enum\":{\"title\":\"enum\",\"type\":\"array\",\"items\":true,\"minItems\":1,\"uniqueItems\":true},\"simpleTypes\":{\"title\":\"simpleTypes\",\"enum\":[\"array\",\"boolean\",\"integer\",\"null\",\"number\",\"object\",\"string\"]},\"arrayOfSimpleTypes\":{\"title\":\"arrayOfSimpleTypes\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/simpleTypes\"},\"minItems\":1,\"uniqueItems\":true},\"type\":{\"title\":\"type\",\"anyOf\":[{\"$ref\":\"#/definitions/simpleTypes\"},{\"$ref\":\"#/definitions/arrayOfSimpleTypes\"}]},\"format\":{\"title\":\"format\",\"type\":\"string\"},\"contentMediaType\":{\"title\":\"contentMediaType\",\"type\":\"string\"},\"contentEncoding\":{\"title\":\"contentEncoding\",\"type\":\"string\"},\"JSONSchemaObject\":{\"title\":\"JSONSchemaObject\",\"type\":\"object\",\"properties\":{\"$id\":{\"$ref\":\"#/definitions/$id\"},\"$schema\":{\"$ref\":\"#/definitions/$schema\"},\"$ref\":{\"$ref\":\"#/definitions/$ref\"},\"$comment\":{\"$ref\":\"#/definitions/$comment\"},\"title\":{\"$ref\":\"#/definitions/title\"},\"description\":{\"$ref\":\"#/definitions/description\"},\"default\":true,\"readOnly\":{\"$ref\":\"#/definitions/readOnly\"},\"examples\":{\"$ref\":\"#/definitions/examples\"},\"multipleOf\":{\"$ref\":\"#/definitions/multipleOf\"},\"maximum\":{\"$ref\":\"#/definitions/maximum\"},\"exclusiveMaximum\":{\"$ref\":\"#/definitions/exclusiveMaximum\"},\"minimum\":{\"$ref\":\"#/definitions/minimum\"},\"exclusiveMinimum\":{\"$ref\":\"#/definitions/exclusiveMinimum\"},\"maxLength\":{\"$ref\":\"#/definitions/nonNegativeInteger\"},\"minLength\":{\"$ref\":\"#/definitions/nonNegativeIntegerDefaultZero\"},\"pattern\":{\"$ref\":\"#/definitions/pattern\"},\"additionalItems\":{\"$ref\":\"#/definitions/undefined\"},\"items\":{\"$ref\":\"#/definitions/items\"},\"maxItems\":{\"$ref\":\"#/definitions/nonNegativeInteger\"},\"minItems\":{\"$ref\":\"#/definitions/nonNegativeIntegerDefaultZero\"},\"uniqueItems\":{\"$ref\":\"#/definitions/uniqueItems\"},\"contains\":{\"$ref\":\"#/definitions/undefined\"},\"maxProperties\":{\"$ref\":\"#/definitions/nonNegativeInteger\"},\"minProperties\":{\"$ref\":\"#/definitions/nonNegativeIntegerDefaultZero\"},\"required\":{\"$ref\":\"#/definitions/stringArray\"},\"additionalProperties\":{\"$ref\":\"#/definitions/undefined\"},\"definitions\":{\"$ref\":\"#/definitions/definitions\"},\"properties\":{\"$ref\":\"#/definitions/properties\"},\"patternProperties\":{\"$ref\":\"#/definitions/patternProperties\"},\"dependencies\":{\"$ref\":\"#/definitions/dependencies\"},\"propertyNames\":{\"$ref\":\"#/definitions/undefined\"},\"const\":true,\"enum\":{\"$ref\":\"#/definitions/enum\"},\"type\":{\"$ref\":\"#/definitions/type\"},\"format\":{\"$ref\":\"#/definitions/format\"},\"contentMediaType\":{\"$ref\":\"#/definitions/contentMediaType\"},\"contentEncoding\":{\"$ref\":\"#/definitions/contentEncoding\"},\"if\":{\"$ref\":\"#/definitions/undefined\"},\"then\":{\"$ref\":\"#/definitions/undefined\"},\"else\":{\"$ref\":\"#/definitions/undefined\"},\"allOf\":{\"$ref\":\"#/definitions/schemaArray\"},\"anyOf\":{\"$ref\":\"#/definitions/schemaArray\"},\"oneOf\":{\"$ref\":\"#/definitions/schemaArray\"},\"not\":{\"$ref\":\"#/definitions/undefined\"}}},\"JSONSchemaBoolean\":{\"title\":\"JSONSchemaBoolean\",\"description\":\"Always valid if true. Never valid if false. Is constant.\",\"type\":\"boolean\"},\"JSONSchema\":{\"$schema\":\"https://meta.json-schema.tools/\",\"$id\":\"https://meta.json-schema.tools/\",\"title\":\"JSONSchema\",\"default\":{},\"oneOf\":[{\"$ref\":\"#/definitions/JSONSchemaObject\"},{\"$ref\":\"#/definitions/JSONSchemaBoolean\"}],\"definitions\":{\"JSONSchemaBoolean\":{\"$ref\":\"#/definitions/JSONSchemaBoolean\"},\"JSONSchemaObject\":{\"$ref\":\"#/definitions/JSONSchemaObject\"},\"schemaArray\":{\"$ref\":\"#/definitions/schemaArray\"},\"nonNegativeInteger\":{\"$ref\":\"#/definitions/nonNegativeInteger\"},\"nonNegativeIntegerDefault0\":{\"$ref\":\"#/definitions/nonNegativeIntegerDefaultZero\"},\"simpleTypes\":{\"$ref\":\"#/definitions/simpleTypes\"},\"stringArray\":{\"$ref\":\"#/definitions/stringArray\"}},\"isCycle\":true},\"contentDescriptorObjectRequired\":{\"title\":\"contentDescriptorObjectRequired\",\"type\":\"boolean\",\"default\":false},\"contentDescriptorObjectDeprecated\":{\"title\":\"contentDescriptorObjectDeprecated\",\"type\":\"boolean\",\"default\":false},\"contentDescriptorObject\":{\"title\":\"contentDescriptorObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"name\",\"schema\"],\"properties\":{\"name\":{\"$ref\":\"#/definitions/contentDescriptorObjectName\"},\"description\":{\"$ref\":\"#/definitions/contentDescriptorObjectDescription\"},\"summary\":{\"$ref\":\"#/definitions/contentDescriptorObjectSummary\"},\"schema\":{\"$ref\":\"#/definitions/undefined\"},\"required\":{\"$ref\":\"#/definitions/contentDescriptorObjectRequired\"},\"deprecated\":{\"$ref\":\"#/definitions/contentDescriptorObjectDeprecated\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"contentDescriptorOrReference\":{\"title\":\"contentDescriptorOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/contentDescriptorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"methodObjectParams\":{\"title\":\"methodObjectParams\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/contentDescriptorOrReference\"}},\"methodObjectResult\":{\"title\":\"methodObjectResult\",\"oneOf\":[{\"$ref\":\"#/definitions/contentDescriptorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"errorObjectCode\":{\"title\":\"errorObjectCode\",\"description\":\"A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.\",\"type\":\"integer\"},\"errorObjectMessage\":{\"title\":\"errorObjectMessage\",\"description\":\"A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.\",\"type\":\"string\"},\"errorObjectData\":{\"title\":\"errorObjectData\",\"description\":\"A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).\"},\"errorObject\":{\"title\":\"errorObject\",\"type\":\"object\",\"description\":\"Defines an application level error.\",\"additionalProperties\":false,\"required\":[\"code\",\"message\"],\"properties\":{\"code\":{\"$ref\":\"#/definitions/errorObjectCode\"},\"message\":{\"$ref\":\"#/definitions/errorObjectMessage\"},\"data\":{\"$ref\":\"#/definitions/errorObjectData\"}}},\"errorOrReference\":{\"title\":\"errorOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/errorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"methodObjectErrors\":{\"title\":\"methodObjectErrors\",\"description\":\"Defines an application level error.\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/errorOrReference\"}},\"linkObjectName\":{\"title\":\"linkObjectName\",\"type\":\"string\",\"minLength\":1},\"linkObjectSummary\":{\"title\":\"linkObjectSummary\",\"type\":\"string\"},\"linkObjectMethod\":{\"title\":\"linkObjectMethod\",\"type\":\"string\"},\"linkObjectDescription\":{\"title\":\"linkObjectDescription\",\"type\":\"string\"},\"linkObjectParams\":{\"title\":\"linkObjectParams\"},\"linkObjectServer\":{\"title\":\"linkObjectServer\",\"type\":\"object\",\"required\":[\"url\"],\"additionalProperties\":false,\"properties\":{\"url\":{\"$ref\":\"#/definitions/serverObjectUrl\"},\"name\":{\"$ref\":\"#/definitions/serverObjectName\"},\"description\":{\"$ref\":\"#/definitions/serverObjectDescription\"},\"summary\":{\"$ref\":\"#/definitions/serverObjectSummary\"},\"variables\":{\"$ref\":\"#/definitions/serverObjectVariables\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"linkObject\":{\"title\":\"linkObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"$ref\":\"#/definitions/linkObjectName\"},\"summary\":{\"$ref\":\"#/definitions/linkObjectSummary\"},\"method\":{\"$ref\":\"#/definitions/linkObjectMethod\"},\"description\":{\"$ref\":\"#/definitions/linkObjectDescription\"},\"params\":{\"$ref\":\"#/definitions/linkObjectParams\"},\"server\":{\"$ref\":\"#/definitions/linkObjectServer\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"linkOrReference\":{\"title\":\"linkOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/linkObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"methodObjectLinks\":{\"title\":\"methodObjectLinks\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/linkOrReference\"}},\"examplePairingObjectName\":{\"title\":\"examplePairingObjectName\",\"type\":\"string\",\"minLength\":1},\"examplePairingObjectDescription\":{\"title\":\"examplePairingObjectDescription\",\"type\":\"string\"},\"exampleObjectSummary\":{\"title\":\"exampleObjectSummary\",\"type\":\"string\"},\"exampleObjectValue\":{\"title\":\"exampleObjectValue\"},\"exampleObjectDescription\":{\"title\":\"exampleObjectDescription\",\"type\":\"string\"},\"exampleObjectName\":{\"title\":\"exampleObjectName\",\"type\":\"string\",\"minLength\":1},\"exampleObject\":{\"title\":\"exampleObject\",\"type\":\"object\",\"required\":[\"name\",\"value\"],\"properties\":{\"summary\":{\"$ref\":\"#/definitions/exampleObjectSummary\"},\"value\":{\"$ref\":\"#/definitions/exampleObjectValue\"},\"description\":{\"$ref\":\"#/definitions/exampleObjectDescription\"},\"name\":{\"$ref\":\"#/definitions/exampleObjectName\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"exampleOrReference\":{\"title\":\"exampleOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/exampleObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"examplePairingObjectParams\":{\"title\":\"examplePairingObjectParams\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/exampleOrReference\"}},\"examplePairingObjectResult\":{\"title\":\"examplePairingObjectResult\",\"oneOf\":[{\"$ref\":\"#/definitions/exampleObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"examplePairingObject\":{\"title\":\"examplePairingObject\",\"type\":\"object\",\"required\":[\"name\",\"params\"],\"properties\":{\"name\":{\"$ref\":\"#/definitions/examplePairingObjectName\"},\"description\":{\"$ref\":\"#/definitions/examplePairingObjectDescription\"},\"params\":{\"$ref\":\"#/definitions/examplePairingObjectParams\"},\"result\":{\"$ref\":\"#/definitions/examplePairingObjectResult\"}}},\"examplePairingOrReference\":{\"title\":\"examplePairingOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/examplePairingObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"methodObjectExamples\":{\"title\":\"methodObjectExamples\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/examplePairingOrReference\"}},\"methodObjectDeprecated\":{\"title\":\"methodObjectDeprecated\",\"type\":\"boolean\",\"default\":false},\"methodObject\":{\"title\":\"methodObject\",\"type\":\"object\",\"required\":[\"name\",\"params\"],\"additionalProperties\":false,\"properties\":{\"name\":{\"$ref\":\"#/definitions/methodObjectName\"},\"description\":{\"$ref\":\"#/definitions/methodObjectDescription\"},\"summary\":{\"$ref\":\"#/definitions/methodObjectSummary\"},\"servers\":{\"$ref\":\"#/definitions/servers\"},\"tags\":{\"$ref\":\"#/definitions/methodObjectTags\"},\"paramStructure\":{\"$ref\":\"#/definitions/methodObjectParamStructure\"},\"params\":{\"$ref\":\"#/definitions/methodObjectParams\"},\"result\":{\"$ref\":\"#/definitions/methodObjectResult\"},\"errors\":{\"$ref\":\"#/definitions/methodObjectErrors\"},\"links\":{\"$ref\":\"#/definitions/methodObjectLinks\"},\"examples\":{\"$ref\":\"#/definitions/methodObjectExamples\"},\"deprecated\":{\"$ref\":\"#/definitions/methodObjectDeprecated\"},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"methodOrReference\":{\"title\":\"methodOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/methodObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"methods\":{\"title\":\"methods\",\"type\":\"array\",\"additionalItems\":false,\"items\":{\"$ref\":\"#/definitions/methodOrReference\"}},\"schemaComponents\":{\"title\":\"schemaComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/undefined\"}}},\"linkComponents\":{\"title\":\"linkComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/linkObject\"}}},\"errorComponents\":{\"title\":\"errorComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/errorObject\"}}},\"exampleComponents\":{\"title\":\"exampleComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/exampleObject\"}}},\"examplePairingComponents\":{\"title\":\"examplePairingComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/examplePairingObject\"}}},\"contentDescriptorComponents\":{\"title\":\"contentDescriptorComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/contentDescriptorObject\"}}},\"tagComponents\":{\"title\":\"tagComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/tagObject\"}}},\"components\":{\"title\":\"components\",\"type\":\"object\",\"properties\":{\"schemas\":{\"$ref\":\"#/definitions/schemaComponents\"},\"links\":{\"$ref\":\"#/definitions/linkComponents\"},\"errors\":{\"$ref\":\"#/definitions/errorComponents\"},\"examples\":{\"$ref\":\"#/definitions/exampleComponents\"},\"examplePairings\":{\"$ref\":\"#/definitions/examplePairingComponents\"},\"contentDescriptors\":{\"$ref\":\"#/definitions/contentDescriptorComponents\"},\"tags\":{\"$ref\":\"#/definitions/tagComponents\"}}},\"metaSchema\":{\"title\":\"metaSchema\",\"description\":\"JSON Schema URI (used by some editors)\",\"type\":\"string\",\"default\":\"https://meta.open-rpc.org/\"}}}"