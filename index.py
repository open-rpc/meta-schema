from typing import TypedDict
from typing import Optional
from enum import Enum
from typing import NewType
from typing import Any
from typing import List
from typing import Mapping
from typing import Union

class Openrpc(Enum):
    OneThreeTwo = 0
    OneThreeOne = 1
    OneThreeZero = 2
    OneTwoSix = 3
    OneTwoFive = 4
    OneTwoFour = 5
    OneTwoThree = 6
    OneTwoTwo = 7
    OneTwoOne = 8
    OneTwoZero = 9
    OneOneOneTwo = 10
    OneOneOneOne = 11
    OneOneOneZero = 12
    OneOneNine = 13
    OneOneEight = 14
    OneOneSeven = 15
    OneOneSix = 16
    OneOneFive = 17
    OneOneFour = 18
    OneOneThree = 19
    OneOneTwo = 20
    OneOneOne = 21
    OneOneZero = 22
    OneZeroZero = 23
    OneZeroZeroRcOne = 24
    OneZeroZeroRcZero = 25

InfoObjectProperties = NewType("InfoObjectProperties", str)

InfoObjectDescription = NewType("InfoObjectDescription", str)

InfoObjectTermsOfService = NewType("InfoObjectTermsOfService", str)

InfoObjectVersion = NewType("InfoObjectVersion", str)

ContactObjectName = NewType("ContactObjectName", str)

ContactObjectEmail = NewType("ContactObjectEmail", str)

ContactObjectUrl = NewType("ContactObjectUrl", str)

SpecificationExtension = NewType("SpecificationExtension", Any)

class ContactObject(TypedDict):
    name: Optional[ContactObjectName]
    email: Optional[ContactObjectEmail]
    url: Optional[ContactObjectUrl]

LicenseObjectName = NewType("LicenseObjectName", str)

LicenseObjectUrl = NewType("LicenseObjectUrl", str)

class LicenseObject(TypedDict):
    name: Optional[LicenseObjectName]
    url: Optional[LicenseObjectUrl]

class InfoObject(TypedDict):
    title: Optional[InfoObjectProperties]
    description: Optional[InfoObjectDescription]
    termsOfService: Optional[InfoObjectTermsOfService]
    version: Optional[InfoObjectVersion]
    contact: Optional[ContactObject]
    license: Optional[LicenseObject]

ExternalDocumentationObjectDescription = NewType("ExternalDocumentationObjectDescription", str)

ExternalDocumentationObjectUrl = NewType("ExternalDocumentationObjectUrl", str)
"""information about external documentation
"""
class ExternalDocumentationObject(TypedDict):
    description: Optional[ExternalDocumentationObjectDescription]
    url: Optional[ExternalDocumentationObjectUrl]

ServerObjectUrl = NewType("ServerObjectUrl", str)

ServerObjectName = NewType("ServerObjectName", str)

ServerObjectDescription = NewType("ServerObjectDescription", str)

ServerObjectSummary = NewType("ServerObjectSummary", str)

ServerObjectVariableDefault = NewType("ServerObjectVariableDefault", str)

ServerObjectVariableDescription = NewType("ServerObjectVariableDescription", str)

ServerObjectVariableEnumItem = NewType("ServerObjectVariableEnumItem", str)

ServerObjectVariableEnum = NewType("ServerObjectVariableEnum", List[ServerObjectVariableEnumItem])

class ServerObjectVariable(TypedDict):
    default: Optional[ServerObjectVariableDefault]
    description: Optional[ServerObjectVariableDescription]
    enum: Optional[ServerObjectVariableEnum]

ServerObjectVariables = NewType("ServerObjectVariables", Mapping[Any, Any])

class ServerObject(TypedDict):
    url: Optional[ServerObjectUrl]
    name: Optional[ServerObjectName]
    description: Optional[ServerObjectDescription]
    summary: Optional[ServerObjectSummary]
    variables: Optional[ServerObjectVariables]
AlwaysFalse = NewType("AlwaysFalse", Any)

Servers = NewType("Servers", List[ServerObject])
"""The cannonical name for the method. The name MUST be unique within the methods array.
"""
MethodObjectName = NewType("MethodObjectName", str)
"""A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.
"""
MethodObjectDescription = NewType("MethodObjectDescription", str)
"""A short summary of what the method does.
"""
MethodObjectSummary = NewType("MethodObjectSummary", str)

TagObjectName = NewType("TagObjectName", str)

TagObjectDescription = NewType("TagObjectDescription", str)

class TagObject(TypedDict):
    name: Optional[TagObjectName]
    description: Optional[TagObjectDescription]
    externalDocs: Optional[ExternalDocumentationObject]

Ref = NewType("Ref", str)

class ReferenceObject(TypedDict):
    $ref: undefined

TagOrReference = NewType("TagOrReference", Union[TagObject, ReferenceObject])

MethodObjectTags = NewType("MethodObjectTags", List[TagOrReference])
"""Format the server expects the params. Defaults to 'either'.
"""
class MethodObjectParamStructure(Enum):
    ByPosition = 0
    ByName = 1
    Either = 2

ContentDescriptorObjectName = NewType("ContentDescriptorObjectName", str)

ContentDescriptorObjectDescription = NewType("ContentDescriptorObjectDescription", str)

ContentDescriptorObjectSummary = NewType("ContentDescriptorObjectSummary", str)

Id = NewType("Id", str)

Schema = NewType("Schema", str)

Comment = NewType("Comment", str)

Title = NewType("Title", str)

Description = NewType("Description", str)
AlwaysTrue = NewType("AlwaysTrue", Any)

ReadOnly = NewType("ReadOnly", bool)

Examples = NewType("Examples", List[AlwaysTrue])

MultipleOf = NewType("MultipleOf", float)

Maximum = NewType("Maximum", float)

ExclusiveMaximum = NewType("ExclusiveMaximum", float)

Minimum = NewType("Minimum", float)

ExclusiveMinimum = NewType("ExclusiveMinimum", float)

NonNegativeInteger = NewType("NonNegativeInteger", int)

NonNegativeIntegerDefaultZero = NewType("NonNegativeIntegerDefaultZero", int)

Pattern = NewType("Pattern", str)

SchemaArray = NewType("SchemaArray", List[JSONSchema])

Items = NewType("Items", Union[JSONSchema, SchemaArray])

UniqueItems = NewType("UniqueItems", bool)

StringDoaGddGA = NewType("StringDoaGddGA", str)

StringArray = NewType("StringArray", List[StringDoaGddGA])

Definitions = NewType("Definitions", Mapping[Any, Any])

Properties = NewType("Properties", Mapping[Any, Any])

PatternProperties = NewType("PatternProperties", Mapping[Any, Any])

DependenciesSet = NewType("DependenciesSet", Union[JSONSchema, StringArray])

Dependencies = NewType("Dependencies", Mapping[Any, Any])

Enum = NewType("Enum", List[AlwaysTrue])

SimpleTypes = NewType("SimpleTypes", Any)

ArrayOfSimpleTypes = NewType("ArrayOfSimpleTypes", List[SimpleTypes])

Type = NewType("Type", Union[SimpleTypes, ArrayOfSimpleTypes])

Format = NewType("Format", str)

ContentMediaType = NewType("ContentMediaType", str)

ContentEncoding = NewType("ContentEncoding", str)

class JSONSchemaObject(TypedDict):
    $id: Optional[Id]
    $schema: Optional[Schema]
    $ref: Optional[Ref]
    $comment: Optional[Comment]
    title: Optional[Title]
    description: Optional[Description]
    default: Optional[AlwaysTrue]
    readOnly: Optional[ReadOnly]
    examples: Optional[Examples]
    multipleOf: Optional[MultipleOf]
    maximum: Optional[Maximum]
    exclusiveMaximum: Optional[ExclusiveMaximum]
    minimum: Optional[Minimum]
    exclusiveMinimum: Optional[ExclusiveMinimum]
    maxLength: Optional[NonNegativeInteger]
    minLength: Optional[NonNegativeIntegerDefaultZero]
    pattern: Optional[Pattern]
    additionalItems: Optional[JSONSchema]
    items: Optional[Items]
    maxItems: Optional[NonNegativeInteger]
    minItems: Optional[NonNegativeIntegerDefaultZero]
    uniqueItems: Optional[UniqueItems]
    contains: Optional[JSONSchema]
    maxProperties: Optional[NonNegativeInteger]
    minProperties: Optional[NonNegativeIntegerDefaultZero]
    required: Optional[StringArray]
    additionalProperties: Optional[JSONSchema]
    definitions: Optional[Definitions]
    properties: Optional[Properties]
    patternProperties: Optional[PatternProperties]
    dependencies: Optional[Dependencies]
    propertyNames: Optional[JSONSchema]
    const: Optional[AlwaysTrue]
    enum: Optional[Enum]
    type: Optional[Type]
    format: Optional[Format]
    contentMediaType: Optional[ContentMediaType]
    contentEncoding: Optional[ContentEncoding]
    if: Optional[JSONSchema]
    then: Optional[JSONSchema]
    else: Optional[JSONSchema]
    allOf: Optional[SchemaArray]
    anyOf: Optional[SchemaArray]
    oneOf: Optional[SchemaArray]
    not: Optional[JSONSchema]
"""Always valid if true. Never valid if false. Is constant.
"""
JSONSchemaBoolean = NewType("JSONSchemaBoolean", bool)

JSONSchema = NewType("JSONSchema", Union[JSONSchemaObject, JSONSchemaBoolean])

ContentDescriptorObjectRequired = NewType("ContentDescriptorObjectRequired", bool)

ContentDescriptorObjectDeprecated = NewType("ContentDescriptorObjectDeprecated", bool)

class ContentDescriptorObject(TypedDict):
    name: Optional[ContentDescriptorObjectName]
    description: Optional[ContentDescriptorObjectDescription]
    summary: Optional[ContentDescriptorObjectSummary]
    schema: Optional[JSONSchema]
    required: Optional[ContentDescriptorObjectRequired]
    deprecated: Optional[ContentDescriptorObjectDeprecated]

ContentDescriptorOrReference = NewType("ContentDescriptorOrReference", Union[ContentDescriptorObject, ReferenceObject])

MethodObjectParams = NewType("MethodObjectParams", List[ContentDescriptorOrReference])

MethodObjectResult = NewType("MethodObjectResult", Union[ContentDescriptorObject, ReferenceObject])
"""A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.
"""
ErrorObjectCode = NewType("ErrorObjectCode", int)
"""A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
"""
ErrorObjectMessage = NewType("ErrorObjectMessage", str)
"""A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
"""
ErrorObjectData = NewType("ErrorObjectData", Any)
"""Defines an application level error.
"""
class ErrorObject(TypedDict):
    code: Optional[ErrorObjectCode]
    message: Optional[ErrorObjectMessage]
    data: Optional[ErrorObjectData]

ErrorOrReference = NewType("ErrorOrReference", Union[ErrorObject, ReferenceObject])
"""Defines an application level error.
"""
MethodObjectErrors = NewType("MethodObjectErrors", List[ErrorOrReference])

LinkObjectName = NewType("LinkObjectName", str)

LinkObjectSummary = NewType("LinkObjectSummary", str)

LinkObjectMethod = NewType("LinkObjectMethod", str)

LinkObjectDescription = NewType("LinkObjectDescription", str)

LinkObjectParams = NewType("LinkObjectParams", Any)

class LinkObjectServer(TypedDict):
    url: Optional[ServerObjectUrl]
    name: Optional[ServerObjectName]
    description: Optional[ServerObjectDescription]
    summary: Optional[ServerObjectSummary]
    variables: Optional[ServerObjectVariables]

class LinkObject(TypedDict):
    name: Optional[LinkObjectName]
    summary: Optional[LinkObjectSummary]
    method: Optional[LinkObjectMethod]
    description: Optional[LinkObjectDescription]
    params: Optional[LinkObjectParams]
    server: Optional[LinkObjectServer]

LinkOrReference = NewType("LinkOrReference", Union[LinkObject, ReferenceObject])

MethodObjectLinks = NewType("MethodObjectLinks", List[LinkOrReference])

ExamplePairingObjectName = NewType("ExamplePairingObjectName", str)

ExamplePairingObjectDescription = NewType("ExamplePairingObjectDescription", str)

ExampleObjectSummary = NewType("ExampleObjectSummary", str)

ExampleObjectValue = NewType("ExampleObjectValue", Any)

ExampleObjectDescription = NewType("ExampleObjectDescription", str)

ExampleObjectName = NewType("ExampleObjectName", str)

class ExampleObject(TypedDict):
    summary: Optional[ExampleObjectSummary]
    value: Optional[ExampleObjectValue]
    description: Optional[ExampleObjectDescription]
    name: Optional[ExampleObjectName]

ExampleOrReference = NewType("ExampleOrReference", Union[ExampleObject, ReferenceObject])

ExamplePairingObjectParams = NewType("ExamplePairingObjectParams", List[ExampleOrReference])

ExamplePairingObjectResult = NewType("ExamplePairingObjectResult", Union[ExampleObject, ReferenceObject])

class ExamplePairingObject(TypedDict):
    name: Optional[ExamplePairingObjectName]
    description: Optional[ExamplePairingObjectDescription]
    params: Optional[ExamplePairingObjectParams]
    result: Optional[ExamplePairingObjectResult]

ExamplePairingOrReference = NewType("ExamplePairingOrReference", Union[ExamplePairingObject, ReferenceObject])

MethodObjectExamples = NewType("MethodObjectExamples", List[ExamplePairingOrReference])

MethodObjectDeprecated = NewType("MethodObjectDeprecated", bool)

class MethodObject(TypedDict):
    name: Optional[MethodObjectName]
    description: Optional[MethodObjectDescription]
    summary: Optional[MethodObjectSummary]
    servers: Optional[Servers]
    tags: Optional[MethodObjectTags]
    paramStructure: Optional[MethodObjectParamStructure]
    params: Optional[MethodObjectParams]
    result: Optional[MethodObjectResult]
    errors: Optional[MethodObjectErrors]
    links: Optional[MethodObjectLinks]
    examples: Optional[MethodObjectExamples]
    deprecated: Optional[MethodObjectDeprecated]
    externalDocs: Optional[ExternalDocumentationObject]

MethodOrReference = NewType("MethodOrReference", Union[MethodObject, ReferenceObject])

Methods = NewType("Methods", List[MethodOrReference])

SchemaComponents = NewType("SchemaComponents", Mapping[Any, Any])

LinkComponents = NewType("LinkComponents", Mapping[Any, Any])

ErrorComponents = NewType("ErrorComponents", Mapping[Any, Any])

ExampleComponents = NewType("ExampleComponents", Mapping[Any, Any])

ExamplePairingComponents = NewType("ExamplePairingComponents", Mapping[Any, Any])

ContentDescriptorComponents = NewType("ContentDescriptorComponents", Mapping[Any, Any])

TagComponents = NewType("TagComponents", Mapping[Any, Any])

class Components(TypedDict):
    schemas: Optional[SchemaComponents]
    links: Optional[LinkComponents]
    errors: Optional[ErrorComponents]
    examples: Optional[ExampleComponents]
    examplePairings: Optional[ExamplePairingComponents]
    contentDescriptors: Optional[ContentDescriptorComponents]
    tags: Optional[TagComponents]
"""JSON Schema URI (used by some editors)
"""
MetaSchema = NewType("MetaSchema", str)

class OpenrpcDocument(TypedDict):
    openrpc: undefined
    info: Optional[InfoObject]
    externalDocs: Optional[ExternalDocumentationObject]
    servers: Optional[Servers]
    methods: undefined
    components: Optional[Components]
    $schema: Optional[MetaSchema]