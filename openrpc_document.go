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
		a.JSONSchema = &myJSONSchema
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
		a.JSONSchema = &myJSONSchema
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
	if o.JSONSchema != nil {
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

const RawOpenrpc_document = "{\"$schema\":\"https://meta.json-schema.tools/\",\"$id\":\"https://meta.open-rpc.org/\",\"title\":\"openrpcDocument\",\"type\":\"object\",\"required\":[\"info\",\"methods\",\"openrpc\"],\"additionalProperties\":false,\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}},\"properties\":{\"openrpc\":{\"title\":\"openrpc\",\"type\":\"string\",\"enum\":[\"1.3.2\",\"1.3.1\",\"1.3.0\",\"1.2.6\",\"1.2.5\",\"1.2.4\",\"1.2.3\",\"1.2.2\",\"1.2.1\",\"1.2.0\",\"1.1.12\",\"1.1.11\",\"1.1.10\",\"1.1.9\",\"1.1.8\",\"1.1.7\",\"1.1.6\",\"1.1.5\",\"1.1.4\",\"1.1.3\",\"1.1.2\",\"1.1.1\",\"1.1.0\",\"1.0.0\",\"1.0.0-rc1\",\"1.0.0-rc0\"]},\"info\":{\"$ref\":\"#/definitions/infoObject\"},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"},\"servers\":{\"title\":\"servers\",\"type\":\"array\",\"additionalItems\":false,\"items\":{\"$ref\":\"#/definitions/serverObject\"}},\"methods\":{\"title\":\"methods\",\"type\":\"array\",\"additionalItems\":false,\"items\":{\"title\":\"methodOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/methodObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"components\":{\"title\":\"components\",\"type\":\"object\",\"properties\":{\"schemas\":{\"title\":\"schemaComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/JSONSchema\"}}},\"links\":{\"title\":\"linkComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/linkObject\"}}},\"errors\":{\"title\":\"errorComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/errorObject\"}}},\"examples\":{\"title\":\"exampleComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/exampleObject\"}}},\"examplePairings\":{\"title\":\"examplePairingComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/examplePairingObject\"}}},\"contentDescriptors\":{\"title\":\"contentDescriptorComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/contentDescriptorObject\"}}},\"tags\":{\"title\":\"tagComponents\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/tagObject\"}}}}},\"$schema\":{\"title\":\"metaSchema\",\"description\":\"JSON Schema URI (used by some editors)\",\"type\":\"string\",\"default\":\"https://meta.open-rpc.org/\"}},\"definitions\":{\"specificationExtension\":{\"title\":\"specificationExtension\"},\"JSONSchema\":{\"$ref\":\"https://raw.githubusercontent.com/json-schema-tools/meta-schema/1.5.9/src/schema.json\"},\"referenceObject\":{\"title\":\"referenceObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"$ref\"],\"properties\":{\"$ref\":{\"$ref\":\"https://raw.githubusercontent.com/json-schema-tools/meta-schema/1.5.9/src/schema.json#/definitions/JSONSchemaObject/properties/$ref\"}}},\"errorObject\":{\"title\":\"errorObject\",\"type\":\"object\",\"description\":\"Defines an application level error.\",\"additionalProperties\":false,\"required\":[\"code\",\"message\"],\"properties\":{\"code\":{\"title\":\"errorObjectCode\",\"description\":\"A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.\",\"type\":\"integer\"},\"message\":{\"title\":\"errorObjectMessage\",\"description\":\"A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.\",\"type\":\"string\"},\"data\":{\"title\":\"errorObjectData\",\"description\":\"A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).\"}}},\"licenseObject\":{\"title\":\"licenseObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"licenseObjectName\",\"type\":\"string\"},\"url\":{\"title\":\"licenseObjectUrl\",\"type\":\"string\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"contactObject\":{\"title\":\"contactObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"contactObjectName\",\"type\":\"string\"},\"email\":{\"title\":\"contactObjectEmail\",\"type\":\"string\"},\"url\":{\"title\":\"contactObjectUrl\",\"type\":\"string\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"infoObject\":{\"title\":\"infoObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"title\",\"version\"],\"properties\":{\"title\":{\"title\":\"infoObjectProperties\",\"type\":\"string\"},\"description\":{\"title\":\"infoObjectDescription\",\"type\":\"string\"},\"termsOfService\":{\"title\":\"infoObjectTermsOfService\",\"type\":\"string\",\"format\":\"uri\"},\"version\":{\"title\":\"infoObjectVersion\",\"type\":\"string\"},\"contact\":{\"$ref\":\"#/definitions/contactObject\"},\"license\":{\"$ref\":\"#/definitions/licenseObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"serverObject\":{\"title\":\"serverObject\",\"type\":\"object\",\"required\":[\"url\"],\"additionalProperties\":false,\"properties\":{\"url\":{\"title\":\"serverObjectUrl\",\"type\":\"string\",\"format\":\"uri\"},\"name\":{\"title\":\"serverObjectName\",\"type\":\"string\"},\"description\":{\"title\":\"serverObjectDescription\",\"type\":\"string\"},\"summary\":{\"title\":\"serverObjectSummary\",\"type\":\"string\"},\"variables\":{\"title\":\"serverObjectVariables\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"title\":\"serverObjectVariable\",\"type\":\"object\",\"required\":[\"default\"],\"properties\":{\"default\":{\"title\":\"serverObjectVariableDefault\",\"type\":\"string\"},\"description\":{\"title\":\"serverObjectVariableDescription\",\"type\":\"string\"},\"enum\":{\"title\":\"serverObjectVariableEnum\",\"type\":\"array\",\"items\":{\"title\":\"serverObjectVariableEnumItem\",\"type\":\"string\"}}}}}}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"linkObject\":{\"title\":\"linkObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"linkObjectName\",\"type\":\"string\",\"minLength\":1},\"summary\":{\"title\":\"linkObjectSummary\",\"type\":\"string\"},\"method\":{\"title\":\"linkObjectMethod\",\"type\":\"string\"},\"description\":{\"title\":\"linkObjectDescription\",\"type\":\"string\"},\"params\":{\"title\":\"linkObjectParams\"},\"server\":{\"title\":\"linkObjectServer\",\"$ref\":\"#/definitions/serverObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"externalDocumentationObject\":{\"title\":\"externalDocumentationObject\",\"type\":\"object\",\"additionalProperties\":false,\"description\":\"information about external documentation\",\"required\":[\"url\"],\"properties\":{\"description\":{\"title\":\"externalDocumentationObjectDescription\",\"type\":\"string\"},\"url\":{\"title\":\"externalDocumentationObjectUrl\",\"type\":\"string\",\"format\":\"uri\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"methodObject\":{\"title\":\"methodObject\",\"type\":\"object\",\"required\":[\"name\",\"params\"],\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"methodObjectName\",\"description\":\"The cannonical name for the method. The name MUST be unique within the methods array.\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"methodObjectDescription\",\"description\":\"A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.\",\"type\":\"string\"},\"summary\":{\"title\":\"methodObjectSummary\",\"description\":\"A short summary of what the method does.\",\"type\":\"string\"},\"servers\":{\"title\":\"servers\",\"type\":\"array\",\"additionalItems\":false,\"items\":{\"$ref\":\"#/definitions/serverObject\"}},\"tags\":{\"title\":\"methodObjectTags\",\"type\":\"array\",\"items\":{\"title\":\"tagOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/tagObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"paramStructure\":{\"title\":\"methodObjectParamStructure\",\"type\":\"string\",\"description\":\"Format the server expects the params. Defaults to 'either'.\",\"enum\":[\"by-position\",\"by-name\",\"either\"],\"default\":\"either\"},\"params\":{\"title\":\"methodObjectParams\",\"type\":\"array\",\"items\":{\"title\":\"contentDescriptorOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/contentDescriptorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"result\":{\"title\":\"methodObjectResult\",\"oneOf\":[{\"$ref\":\"#/definitions/contentDescriptorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"errors\":{\"title\":\"methodObjectErrors\",\"description\":\"Defines an application level error.\",\"type\":\"array\",\"items\":{\"title\":\"errorOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/errorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"links\":{\"title\":\"methodObjectLinks\",\"type\":\"array\",\"items\":{\"title\":\"linkOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/linkObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"examples\":{\"title\":\"methodObjectExamples\",\"type\":\"array\",\"items\":{\"title\":\"examplePairingOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/examplePairingObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"deprecated\":{\"title\":\"methodObjectDeprecated\",\"type\":\"boolean\",\"default\":false},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"tagObject\":{\"title\":\"tagObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"name\"],\"properties\":{\"name\":{\"title\":\"tagObjectName\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"tagObjectDescription\",\"type\":\"string\"},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"exampleObject\":{\"title\":\"exampleObject\",\"type\":\"object\",\"required\":[\"name\",\"value\"],\"properties\":{\"summary\":{\"title\":\"exampleObjectSummary\",\"type\":\"string\"},\"value\":{\"title\":\"exampleObjectValue\"},\"description\":{\"title\":\"exampleObjectDescription\",\"type\":\"string\"},\"name\":{\"title\":\"exampleObjectName\",\"type\":\"string\",\"minLength\":1}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"examplePairingObject\":{\"title\":\"examplePairingObject\",\"type\":\"object\",\"required\":[\"name\",\"params\"],\"properties\":{\"name\":{\"title\":\"examplePairingObjectName\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"examplePairingObjectDescription\",\"type\":\"string\"},\"params\":{\"title\":\"examplePairingObjectParams\",\"type\":\"array\",\"items\":{\"title\":\"exampleOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/exampleObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"result\":{\"title\":\"examplePairingObjectResult\",\"oneOf\":[{\"$ref\":\"#/definitions/exampleObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}}},\"contentDescriptorObject\":{\"title\":\"contentDescriptorObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"name\",\"schema\"],\"properties\":{\"name\":{\"title\":\"contentDescriptorObjectName\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"contentDescriptorObjectDescription\",\"type\":\"string\"},\"summary\":{\"title\":\"contentDescriptorObjectSummary\",\"type\":\"string\"},\"schema\":{\"$ref\":\"#/definitions/JSONSchema\"},\"required\":{\"title\":\"contentDescriptorObjectRequired\",\"type\":\"boolean\",\"default\":false},\"deprecated\":{\"title\":\"contentDescriptorObjectDeprecated\",\"type\":\"boolean\",\"default\":false}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}}}}"