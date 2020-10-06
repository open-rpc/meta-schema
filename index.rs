use serde::{Serialize, Deserialize};
use std::collections::HashMap;
extern crate serde_json;

#[derive(Serialize, Deserialize)]
pub enum Openrpc {
    #[serde(rename = "1.2.6")]
    26,
    #[serde(rename = "1.2.5")]
    25,
    #[serde(rename = "1.2.4")]
    24,
    #[serde(rename = "1.2.3")]
    23,
    #[serde(rename = "1.2.2")]
    22,
    #[serde(rename = "1.2.1")]
    21,
    #[serde(rename = "1.2.0")]
    20,
    #[serde(rename = "1.1.12")]
    112,
    #[serde(rename = "1.1.11")]
    111,
    #[serde(rename = "1.1.10")]
    110,
    #[serde(rename = "1.1.9")]
    19,
    #[serde(rename = "1.1.8")]
    18,
    #[serde(rename = "1.1.7")]
    17,
    #[serde(rename = "1.1.6")]
    16,
    #[serde(rename = "1.1.5")]
    15,
    #[serde(rename = "1.1.4")]
    14,
    #[serde(rename = "1.1.3")]
    13,
    #[serde(rename = "1.1.2")]
    12,
    #[serde(rename = "1.1.1")]
    11,
    #[serde(rename = "1.1.0")]
    10,
    #[serde(rename = "1.0.0")]
    00,
    #[serde(rename = "1.0.0-rc1")]
    00Rc1,
    #[serde(rename = "1.0.0-rc0")]
    00Rc0,
}
pub type InfoObjectProperties = String;
pub type InfoObjectDescription = String;
pub type InfoObjectTermsOfService = String;
pub type InfoObjectVersion = String;
pub type ContactObjectName = String;
pub type ContactObjectEmail = String;
pub type ContactObjectUrl = String;
pub type AnyL9Fw4VUO = serde_json::Value;
#[derive(Serialize, Deserialize)]
pub struct ContactObject {
    pub(crate) name: Option<ContactObjectName>,
    pub(crate) email: Option<ContactObjectEmail>,
    pub(crate) url: Option<ContactObjectUrl>,
}
pub type LicenseObjectName = String;
pub type LicenseObjectUrl = String;
#[derive(Serialize, Deserialize)]
pub struct LicenseObject {
    pub(crate) name: Option<LicenseObjectName>,
    pub(crate) url: Option<LicenseObjectUrl>,
}
#[derive(Serialize, Deserialize)]
pub struct InfoObject {
    pub(crate) title: InfoObjectProperties,
    pub(crate) description: Option<InfoObjectDescription>,
    pub(crate) termsOfService: Option<InfoObjectTermsOfService>,
    pub(crate) version: InfoObjectVersion,
    pub(crate) contact: Option<ContactObject>,
    pub(crate) license: Option<LicenseObject>,
}
pub type ExternalDocumentationObjectDescription = String;
pub type ExternalDocumentationObjectUrl = String;
/// ExternalDocumentationObject
///
/// information about external documentation
///
#[derive(Serialize, Deserialize)]
pub struct ExternalDocumentationObject {
    pub(crate) description: Option<ExternalDocumentationObjectDescription>,
    pub(crate) url: ExternalDocumentationObjectUrl,
}
pub type ServerObjectUrl = String;
pub type ServerObjectName = String;
pub type ServerObjectDescription = String;
pub type ServerObjectSummary = String;
pub type ServerObjectVariableDefault = String;
pub type ServerObjectVariableDescription = String;
pub type ServerObjectVariableEnumItem = String;
pub type ServerObjectVariableEnum = Vec<ServerObjectVariableEnumItem>;
#[derive(Serialize, Deserialize)]
pub struct ServerObjectVariable {
    pub(crate) default: ServerObjectVariableDefault,
    pub(crate) description: Option<ServerObjectVariableDescription>,
    pub(crate) enum: Option<ServerObjectVariableEnum>,
}
pub type ServerObjectVariables = HashMap<String, Option<serde_json::Value>>;
#[derive(Serialize, Deserialize)]
pub struct ServerObject {
    pub(crate) url: ServerObjectUrl,
    pub(crate) name: Option<ServerObjectName>,
    pub(crate) description: Option<ServerObjectDescription>,
    pub(crate) summary: Option<ServerObjectSummary>,
    pub(crate) variables: Option<ServerObjectVariables>,
}
pub type Servers = Vec<ServerObject>;
/// MethodObjectName
///
/// The cannonical name for the method. The name MUST be unique within the methods array.
///
pub type MethodObjectName = String;
/// MethodObjectDescription
///
/// A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.
///
pub type MethodObjectDescription = String;
/// MethodObjectSummary
///
/// A short summary of what the method does.
///
pub type MethodObjectSummary = String;
pub type TagObjectName = String;
pub type TagObjectDescription = String;
#[derive(Serialize, Deserialize)]
pub struct TagObject {
    pub(crate) name: TagObjectName,
    pub(crate) description: Option<TagObjectDescription>,
    pub(crate) externalDocs: Option<ExternalDocumentationObject>,
}
pub type Ref = String;
#[derive(Serialize, Deserialize)]
pub struct ReferenceObject {
    pub(crate) $ref: Ref,
}
#[derive(Serialize, Deserialize)]
pub enum OneOfReferenceObjectTagObjectMTCfXRqB {
    TagObject,
    ReferenceObject
}
pub type MethodObjectTags = Vec<OneOfReferenceObjectTagObjectMTCfXRqB>;
/// MethodObjectParamStructure
///
/// Format the server expects the params. Defaults to 'either'.
///
/// # Default
///
/// either
///
#[derive(Serialize, Deserialize)]
pub enum MethodObjectParamStructure {
    #[serde(rename = "by-position")]
    ByPosition,
    #[serde(rename = "by-name")]
    ByName,
    #[serde(rename = "either")]
    Either,
}
pub type ContentDescriptorObjectName = String;
pub type ContentDescriptorObjectDescription = String;
pub type ContentDescriptorObjectSummary = String;
pub type $Id = String;
pub type $Schema = String;
pub type $Ref = String;
pub type $Comment = String;
pub type Title = String;
pub type Description = String;
type AlwaysTrue = serde_json::Value;
pub type ReadOnly = bool;
pub type Examples = Vec<AlwaysTrue>;
pub type MultipleOf = f64;
pub type Maximum = f64;
pub type ExclusiveMaximum = f64;
pub type Minimum = f64;
pub type ExclusiveMinimum = f64;
pub type NonNegativeInteger = i64;
pub type DefaultZero = serde_json::Value;
pub type NonNegativeIntegerDefaultZero = HashMap<String, Option<serde_json::Value>>;
pub type Pattern = String;
pub type SchemaArray = Vec<JSONSchema>;
#[derive(Serialize, Deserialize)]
pub enum Items {
    JSONSchema,
    SchemaArray
}
pub type UniqueItems = bool;
pub type StringDoaGddGA = String;
/// StringArray
///
/// # Default
///
/// []
///
pub type StringArray = Vec<StringDoaGddGA>;
/// Definitions
///
/// # Default
///
/// {}
///
pub type Definitions = HashMap<String, Option<serde_json::Value>>;
/// Properties
///
/// # Default
///
/// {}
///
pub type Properties = HashMap<String, Option<serde_json::Value>>;
/// PatternProperties
///
/// # Default
///
/// {}
///
pub type PatternProperties = HashMap<String, Option<serde_json::Value>>;
#[derive(Serialize, Deserialize)]
pub enum DependenciesSet {
    JSONSchema,
    StringArray
}
pub type Dependencies = HashMap<String, Option<serde_json::Value>>;
pub type Enum = Vec<AlwaysTrue>;
pub type SimpleTypes = serde_json::Value;
pub type ArrayOfSimpleTypes = Vec<SimpleTypes>;
#[derive(Serialize, Deserialize)]
pub enum Type {
    SimpleTypes,
    ArrayOfSimpleTypes
}
pub type Format = String;
pub type ContentMediaType = String;
pub type ContentEncoding = String;
#[derive(Serialize, Deserialize)]
pub struct JSONSchemaObject {
    pub(crate) $id: Option<$Id>,
    pub(crate) $schema: Option<$Schema>,
    pub(crate) $ref: Option<$Ref>,
    pub(crate) $comment: Option<$Comment>,
    pub(crate) title: Option<Title>,
    pub(crate) description: Option<Description>,
    pub(crate) default: Option<AlwaysTrue>,
    pub(crate) readOnly: Option<ReadOnly>,
    pub(crate) examples: Option<Examples>,
    pub(crate) multipleOf: Option<MultipleOf>,
    pub(crate) maximum: Option<Maximum>,
    pub(crate) exclusiveMaximum: Option<ExclusiveMaximum>,
    pub(crate) minimum: Option<Minimum>,
    pub(crate) exclusiveMinimum: Option<ExclusiveMinimum>,
    pub(crate) maxLength: Option<NonNegativeInteger>,
    pub(crate) minLength: Option<NonNegativeIntegerDefaultZero>,
    pub(crate) pattern: Option<Pattern>,
    pub(crate) additionalItems: Option<JSONSchema>,
    pub(crate) items: Option<Items>,
    pub(crate) maxItems: Option<NonNegativeInteger>,
    pub(crate) minItems: Option<NonNegativeIntegerDefaultZero>,
    pub(crate) uniqueItems: Option<UniqueItems>,
    pub(crate) contains: Option<JSONSchema>,
    pub(crate) maxProperties: Option<NonNegativeInteger>,
    pub(crate) minProperties: Option<NonNegativeIntegerDefaultZero>,
    pub(crate) required: Option<StringArray>,
    pub(crate) additionalProperties: Option<JSONSchema>,
    pub(crate) definitions: Option<Definitions>,
    pub(crate) properties: Option<Properties>,
    pub(crate) patternProperties: Option<PatternProperties>,
    pub(crate) dependencies: Option<Dependencies>,
    pub(crate) propertyNames: Option<JSONSchema>,
    pub(crate) const: Option<AlwaysTrue>,
    pub(crate) enum: Option<Enum>,
    pub(crate) type: Option<Type>,
    pub(crate) format: Option<Format>,
    pub(crate) contentMediaType: Option<ContentMediaType>,
    pub(crate) contentEncoding: Option<ContentEncoding>,
    pub(crate) if: Option<JSONSchema>,
    pub(crate) then: Option<JSONSchema>,
    pub(crate) else: Option<JSONSchema>,
    pub(crate) allOf: Option<SchemaArray>,
    pub(crate) anyOf: Option<SchemaArray>,
    pub(crate) oneOf: Option<SchemaArray>,
    pub(crate) not: Option<JSONSchema>,
}
/// JSONSchemaBoolean
///
/// Always valid if true. Never valid if false. Is constant.
///
pub type JSONSchemaBoolean = bool;
#[derive(Serialize, Deserialize)]
pub enum JSONSchema {
    JSONSchemaObject,
    JSONSchemaBoolean
}
pub type ContentDescriptorObjectRequired = bool;
pub type ContentDescriptorObjectDeprecated = bool;
#[derive(Serialize, Deserialize)]
pub struct ContentDescriptorObject {
    pub(crate) name: ContentDescriptorObjectName,
    pub(crate) description: Option<ContentDescriptorObjectDescription>,
    pub(crate) summary: Option<ContentDescriptorObjectSummary>,
    pub(crate) schema: JSONSchema,
    pub(crate) required: Option<ContentDescriptorObjectRequired>,
    pub(crate) deprecated: Option<ContentDescriptorObjectDeprecated>,
}
#[derive(Serialize, Deserialize)]
pub enum OneOfContentDescriptorObjectReferenceObjectI0Ye8PrQ {
    ContentDescriptorObject,
    ReferenceObject
}
pub type MethodObjectParams = Vec<OneOfContentDescriptorObjectReferenceObjectI0Ye8PrQ>;
#[derive(Serialize, Deserialize)]
pub enum MethodObjectResult {
    ContentDescriptorObject,
    ReferenceObject
}
/// ErrorObjectCode
///
/// A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.
///
pub type ErrorObjectCode = i64;
/// ErrorObjectMessage
///
/// A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
///
pub type ErrorObjectMessage = String;
/// ErrorObjectData
///
/// A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
///
pub type ErrorObjectData = serde_json::Value;
/// ErrorObject
///
/// Defines an application level error.
///
#[derive(Serialize, Deserialize)]
pub struct ErrorObject {
    pub(crate) code: ErrorObjectCode,
    pub(crate) message: ErrorObjectMessage,
    pub(crate) data: Option<ErrorObjectData>,
}
#[derive(Serialize, Deserialize)]
pub enum OneOfErrorObjectReferenceObject1KnseVEO {
    ErrorObject,
    ReferenceObject
}
/// MethodObjectErrors
///
/// Defines an application level error.
///
pub type MethodObjectErrors = Vec<OneOfErrorObjectReferenceObject1KnseVEO>;
pub type LinkObjectName = String;
pub type LinkObjectSummary = String;
pub type LinkObjectMethod = String;
pub type LinkObjectDescription = String;
pub type LinkObjectParams = serde_json::Value;
#[derive(Serialize, Deserialize)]
pub struct LinkObject {
    pub(crate) name: Option<LinkObjectName>,
    pub(crate) summary: Option<LinkObjectSummary>,
    pub(crate) method: Option<LinkObjectMethod>,
    pub(crate) description: Option<LinkObjectDescription>,
    pub(crate) params: Option<LinkObjectParams>,
    pub(crate) server: Option<ServerObject>,
}
#[derive(Serialize, Deserialize)]
pub enum OneOfLinkObjectReferenceObjectXyKfUxb0 {
    LinkObject,
    ReferenceObject
}
pub type MethodObjectLinks = Vec<OneOfLinkObjectReferenceObjectXyKfUxb0>;
pub type ExamplePairingObjectName = String;
pub type ExamplePairingObjectDescription = String;
pub type ExampleObjectSummary = String;
pub type ExampleObjectValue = serde_json::Value;
pub type ExampleObjectDescription = String;
pub type ExampleObjectName = String;
#[derive(Serialize, Deserialize)]
pub struct ExampleObject {
    pub(crate) summary: Option<ExampleObjectSummary>,
    pub(crate) value: ExampleObjectValue,
    pub(crate) description: Option<ExampleObjectDescription>,
    pub(crate) name: ExampleObjectName,
}
#[derive(Serialize, Deserialize)]
pub enum OneOfExampleObjectReferenceObject5DJ6EmZt {
    ExampleObject,
    ReferenceObject
}
pub type ExamplePairingObjectParams = Vec<OneOfExampleObjectReferenceObject5DJ6EmZt>;
#[derive(Serialize, Deserialize)]
pub enum ExamplePairingObjectresult {
    ExampleObject,
    ReferenceObject
}
#[derive(Serialize, Deserialize)]
pub struct ExamplePairingObject {
    pub(crate) name: ExamplePairingObjectName,
    pub(crate) description: Option<ExamplePairingObjectDescription>,
    pub(crate) params: ExamplePairingObjectParams,
    pub(crate) result: ExamplePairingObjectresult,
}
#[derive(Serialize, Deserialize)]
pub enum OneOfExamplePairingObjectReferenceObjectWEBfRSyK {
    ExamplePairingObject,
    ReferenceObject
}
pub type MethodObjectExamples = Vec<OneOfExamplePairingObjectReferenceObjectWEBfRSyK>;
pub type MethodObjectDeprecated = bool;
#[derive(Serialize, Deserialize)]
pub struct MethodObject {
    pub(crate) name: MethodObjectName,
    pub(crate) description: Option<MethodObjectDescription>,
    pub(crate) summary: Option<MethodObjectSummary>,
    pub(crate) servers: Option<Servers>,
    pub(crate) tags: Option<MethodObjectTags>,
    pub(crate) paramStructure: Option<MethodObjectParamStructure>,
    pub(crate) params: MethodObjectParams,
    pub(crate) result: MethodObjectResult,
    pub(crate) errors: Option<MethodObjectErrors>,
    pub(crate) links: Option<MethodObjectLinks>,
    pub(crate) examples: Option<MethodObjectExamples>,
    pub(crate) deprecated: Option<MethodObjectDeprecated>,
    pub(crate) externalDocs: Option<ExternalDocumentationObject>,
}
pub type Methods = Vec<MethodObject>;
pub type SchemaComponents = HashMap<String, Option<serde_json::Value>>;
pub type LinkComponents = HashMap<String, Option<serde_json::Value>>;
pub type ObjectTfFA84LI = HashMap<String, Option<serde_json::Value>>;
pub type ExampleComponents = HashMap<String, Option<serde_json::Value>>;
pub type ExamplePairingComponents = HashMap<String, Option<serde_json::Value>>;
pub type ContentDescriptorComponents = HashMap<String, Option<serde_json::Value>>;
pub type TagComponents = HashMap<String, Option<serde_json::Value>>;
#[derive(Serialize, Deserialize)]
pub struct Components {
    pub(crate) schemas: Option<SchemaComponents>,
    pub(crate) links: Option<LinkComponents>,
    pub(crate) errors: Option<ObjectTfFA84LI>,
    pub(crate) examples: Option<ExampleComponents>,
    pub(crate) examplePairings: Option<ExamplePairingComponents>,
    pub(crate) contentDescriptors: Option<ContentDescriptorComponents>,
    pub(crate) tags: Option<TagComponents>,
}
#[derive(Serialize, Deserialize)]
pub struct OpenrpcDocument {
    pub(crate) openrpc: Openrpc,
    pub(crate) info: InfoObject,
    pub(crate) externalDocs: Option<ExternalDocumentationObject>,
    pub(crate) servers: Option<Servers>,
    pub(crate) methods: Methods,
    pub(crate) components: Option<Components>,
}