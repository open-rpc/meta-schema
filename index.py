from enum import Enum

class Openrpc(Enum):
    1.2.6 = 0
    1.2.5 = 1
    1.2.4 = 2
    1.2.3 = 3
    1.2.2 = 4
    1.2.1 = 5
    1.2.0 = 6
    1.1.12 = 7
    1.1.11 = 8
    1.1.10 = 9
    1.1.9 = 10
    1.1.8 = 11
    1.1.7 = 12
    1.1.6 = 13
    1.1.5 = 14
    1.1.4 = 15
    1.1.3 = 16
    1.1.2 = 17
    1.1.1 = 18
    1.1.0 = 19
    1.0.0 = 20
    1.0.0-RC1 = 21
    1.0.0-RC0 = 22
from typing import NewType

InfoObjectProperties = NewType("InfoObjectProperties", str)
from typing import NewType

InfoObjectDescription = NewType("InfoObjectDescription", str)
from typing import NewType

InfoObjectTermsOfService = NewType("InfoObjectTermsOfService", str)
from typing import NewType

InfoObjectVersion = NewType("InfoObjectVersion", str)
from typing import NewType

ContactObjectName = NewType("ContactObjectName", str)
from typing import NewType

ContactObjectEmail = NewType("ContactObjectEmail", str)
from typing import NewType

ContactObjectUrl = NewType("ContactObjectUrl", str)
from typing import Any, NewType

AnyL9Fw4VUO = NewType("AnyL9Fw4VUO", Any)
from typing import TypedDict, Optional

class ContactObject(TypedDict):
    name: Optional[ContactObjectName]
    email: Optional[ContactObjectEmail]
    url: Optional[ContactObjectUrl]
from typing import NewType

LicenseObjectName = NewType("LicenseObjectName", str)
from typing import NewType

LicenseObjectUrl = NewType("LicenseObjectUrl", str)
from typing import TypedDict, Optional

class LicenseObject(TypedDict):
    name: Optional[LicenseObjectName]
    url: Optional[LicenseObjectUrl]
from typing import TypedDict, Optional

class InfoObject(TypedDict):
    title: Optional[InfoObjectProperties]
    description: Optional[InfoObjectDescription]
    termsOfService: Optional[InfoObjectTermsOfService]
    version: Optional[InfoObjectVersion]
    contact: Optional[ContactObject]
    license: Optional[LicenseObject]
from typing import NewType

ExternalDocumentationObjectDescription = NewType("ExternalDocumentationObjectDescription", str)
from typing import NewType

ExternalDocumentationObjectUrl = NewType("ExternalDocumentationObjectUrl", str)
from typing import TypedDict, Optional
"""information about external documentation
"""
class ExternalDocumentationObject(TypedDict):
    description: Optional[ExternalDocumentationObjectDescription]
    url: Optional[ExternalDocumentationObjectUrl]
from typing import NewType

ServerObjectUrl = NewType("ServerObjectUrl", str)
from typing import NewType

ServerObjectName = NewType("ServerObjectName", str)
from typing import NewType

ServerObjectDescription = NewType("ServerObjectDescription", str)
from typing import NewType

ServerObjectSummary = NewType("ServerObjectSummary", str)
from typing import NewType

ServerObjectVariableDefault = NewType("ServerObjectVariableDefault", str)
from typing import NewType

ServerObjectVariableDescription = NewType("ServerObjectVariableDescription", str)
from typing import NewType

ServerObjectVariableEnumItem = NewType("ServerObjectVariableEnumItem", str)
from typing import List, NewType

ServerObjectVariableEnum = NewType("ServerObjectVariableEnum", List[ServerObjectVariableEnumItem])
from typing import TypedDict, Optional

class ServerObjectVariable(TypedDict):
    default: Optional[ServerObjectVariableDefault]
    description: Optional[ServerObjectVariableDescription]
    enum: Optional[ServerObjectVariableEnum]
from typing import NewType, Any, Mapping

ServerObjectVariables = NewType("ServerObjectVariables", Mapping[Any, Any])
from typing import TypedDict, Optional

class ServerObject(TypedDict):
    url: Optional[ServerObjectUrl]
    name: Optional[ServerObjectName]
    description: Optional[ServerObjectDescription]
    summary: Optional[ServerObjectSummary]
    variables: Optional[ServerObjectVariables]
from typing import List, NewType

Servers = NewType("Servers", List[ServerObject])
from typing import NewType
"""The cannonical name for the method. The name MUST be unique within the methods array.
"""
MethodObjectName = NewType("MethodObjectName", str)
from typing import NewType
"""A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.
"""
MethodObjectDescription = NewType("MethodObjectDescription", str)
from typing import NewType
"""A short summary of what the method does.
"""
MethodObjectSummary = NewType("MethodObjectSummary", str)
from typing import NewType

TagObjectName = NewType("TagObjectName", str)
from typing import NewType

TagObjectDescription = NewType("TagObjectDescription", str)
from typing import TypedDict, Optional

class TagObject(TypedDict):
    name: Optional[TagObjectName]
    description: Optional[TagObjectDescription]
    externalDocs: Optional[ExternalDocumentationObject]
from typing import NewType

Ref = NewType("Ref", str)
from typing import TypedDict, Optional

class ReferenceObject(TypedDict):
    $ref: Optional[Ref]
from typing import NewType, Union

OneOfReferenceObjectTagObjectMTCfXRqB = NewType("OneOfReferenceObjectTagObjectMTCfXRqB", Union[TagObject, ReferenceObject])
from typing import List, NewType

MethodObjectTags = NewType("MethodObjectTags", List[OneOfReferenceObjectTagObjectMTCfXRqB])
from enum import Enum
"""Format the server expects the params. Defaults to 'either'.
"""
class MethodObjectParamStructure(Enum):
    BY-POSITION = 0
    BY-NAME = 1
    EITHER = 2
from typing import NewType

ContentDescriptorObjectName = NewType("ContentDescriptorObjectName", str)
from typing import NewType

ContentDescriptorObjectDescription = NewType("ContentDescriptorObjectDescription", str)
from typing import NewType

ContentDescriptorObjectSummary = NewType("ContentDescriptorObjectSummary", str)
from typing import NewType

$Id = NewType("$Id", str)
from typing import NewType

$Schema = NewType("$Schema", str)
from typing import NewType

$Ref = NewType("$Ref", str)
from typing import NewType

$Comment = NewType("$Comment", str)
from typing import NewType

Title = NewType("Title", str)
from typing import NewType

Description = NewType("Description", str)
from typing import Any, NewType

AlwaysTrue = NewType("AlwaysTrue", Any)
from typing import NewType

ReadOnly = NewType("ReadOnly", bool)
from typing import List, NewType

Examples = NewType("Examples", List[AlwaysTrue])
from typing import NewType

MultipleOf = NewType("MultipleOf", float)
from typing import NewType

Maximum = NewType("Maximum", float)
from typing import NewType

ExclusiveMaximum = NewType("ExclusiveMaximum", float)
from typing import NewType

Minimum = NewType("Minimum", float)
from typing import NewType

ExclusiveMinimum = NewType("ExclusiveMinimum", float)
from typing import NewType

NonNegativeInteger = NewType("NonNegativeInteger", int)
from typing import Any, NewType

DefaultZero = NewType("DefaultZero", Any)
from typing import NewType, Any, Mapping

NonNegativeIntegerDefaultZero = NewType("NonNegativeIntegerDefaultZero", Mapping[Any, Any])
from typing import NewType

Pattern = NewType("Pattern", str)
from typing import List, NewType

SchemaArray = NewType("SchemaArray", List[JSONSchema])
from typing import NewType, Union

Items = NewType("Items", Union[JSONSchema, SchemaArray])
from typing import NewType

UniqueItems = NewType("UniqueItems", bool)
from typing import NewType

StringDoaGddGA = NewType("StringDoaGddGA", str)
from typing import List, NewType

StringArray = NewType("StringArray", List[StringDoaGddGA])
from typing import NewType, Any, Mapping

Definitions = NewType("Definitions", Mapping[Any, Any])
from typing import NewType, Any, Mapping

Properties = NewType("Properties", Mapping[Any, Any])
from typing import NewType, Any, Mapping

PatternProperties = NewType("PatternProperties", Mapping[Any, Any])
from typing import NewType, Union

DependenciesSet = NewType("DependenciesSet", Union[JSONSchema, StringArray])
from typing import NewType, Any, Mapping

Dependencies = NewType("Dependencies", Mapping[Any, Any])
from typing import List, NewType

Enum = NewType("Enum", List[AlwaysTrue])
from typing import Any, NewType

SimpleTypes = NewType("SimpleTypes", Any)
from typing import List, NewType

ArrayOfSimpleTypes = NewType("ArrayOfSimpleTypes", List[SimpleTypes])
from typing import NewType, Union

Type = NewType("Type", Union[SimpleTypes, ArrayOfSimpleTypes])
from typing import NewType

Format = NewType("Format", str)
from typing import NewType

ContentMediaType = NewType("ContentMediaType", str)
from typing import NewType

ContentEncoding = NewType("ContentEncoding", str)
from typing import TypedDict, Optional

class JSONSchemaObject(TypedDict):
    $id: Optional[$Id]
    $schema: Optional[$Schema]
    $ref: Optional[$Ref]
    $comment: Optional[$Comment]
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
from typing import NewType
"""Always valid if true. Never valid if false. Is constant.
"""
JSONSchemaBoolean = NewType("JSONSchemaBoolean", bool)
from typing import NewType, Union

JSONSchema = NewType("JSONSchema", Union[JSONSchemaObject, JSONSchemaBoolean])
from typing import NewType

ContentDescriptorObjectRequired = NewType("ContentDescriptorObjectRequired", bool)
from typing import NewType

ContentDescriptorObjectDeprecated = NewType("ContentDescriptorObjectDeprecated", bool)
from typing import TypedDict, Optional

class ContentDescriptorObject(TypedDict):
    name: Optional[ContentDescriptorObjectName]
    description: Optional[ContentDescriptorObjectDescription]
    summary: Optional[ContentDescriptorObjectSummary]
    schema: Optional[JSONSchema]
    required: Optional[ContentDescriptorObjectRequired]
    deprecated: Optional[ContentDescriptorObjectDeprecated]
from typing import NewType, Union

OneOfContentDescriptorObjectReferenceObjectI0Ye8PrQ = NewType("OneOfContentDescriptorObjectReferenceObjectI0Ye8PrQ", Union[ContentDescriptorObject, ReferenceObject])
from typing import List, NewType

MethodObjectParams = NewType("MethodObjectParams", List[OneOfContentDescriptorObjectReferenceObjectI0Ye8PrQ])
from typing import NewType, Union

MethodObjectResult = NewType("MethodObjectResult", Union[ContentDescriptorObject, ReferenceObject])
from typing import NewType
"""A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.
"""
ErrorObjectCode = NewType("ErrorObjectCode", int)
from typing import NewType
"""A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
"""
ErrorObjectMessage = NewType("ErrorObjectMessage", str)
from typing import Any, NewType
"""A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
"""
ErrorObjectData = NewType("ErrorObjectData", Any)
from typing import TypedDict, Optional
"""Defines an application level error.
"""
class ErrorObject(TypedDict):
    code: Optional[ErrorObjectCode]
    message: Optional[ErrorObjectMessage]
    data: Optional[ErrorObjectData]
from typing import NewType, Union

OneOfErrorObjectReferenceObject1KnseVEO = NewType("OneOfErrorObjectReferenceObject1KnseVEO", Union[ErrorObject, ReferenceObject])
from typing import List, NewType
"""Defines an application level error.
"""
MethodObjectErrors = NewType("MethodObjectErrors", List[OneOfErrorObjectReferenceObject1KnseVEO])
from typing import NewType

LinkObjectName = NewType("LinkObjectName", str)
from typing import NewType

LinkObjectSummary = NewType("LinkObjectSummary", str)
from typing import NewType

LinkObjectMethod = NewType("LinkObjectMethod", str)
from typing import NewType

LinkObjectDescription = NewType("LinkObjectDescription", str)
from typing import Any, NewType

LinkObjectParams = NewType("LinkObjectParams", Any)
from typing import TypedDict, Optional

class LinkObject(TypedDict):
    name: Optional[LinkObjectName]
    summary: Optional[LinkObjectSummary]
    method: Optional[LinkObjectMethod]
    description: Optional[LinkObjectDescription]
    params: Optional[LinkObjectParams]
    server: Optional[ServerObject]
from typing import NewType, Union

OneOfLinkObjectReferenceObjectXyKfUxb0 = NewType("OneOfLinkObjectReferenceObjectXyKfUxb0", Union[LinkObject, ReferenceObject])
from typing import List, NewType

MethodObjectLinks = NewType("MethodObjectLinks", List[OneOfLinkObjectReferenceObjectXyKfUxb0])
from typing import NewType

ExamplePairingObjectName = NewType("ExamplePairingObjectName", str)
from typing import NewType

ExamplePairingObjectDescription = NewType("ExamplePairingObjectDescription", str)
from typing import NewType

ExampleObjectSummary = NewType("ExampleObjectSummary", str)
from typing import Any, NewType

ExampleObjectValue = NewType("ExampleObjectValue", Any)
from typing import NewType

ExampleObjectDescription = NewType("ExampleObjectDescription", str)
from typing import NewType

ExampleObjectName = NewType("ExampleObjectName", str)
from typing import TypedDict, Optional

class ExampleObject(TypedDict):
    summary: Optional[ExampleObjectSummary]
    value: Optional[ExampleObjectValue]
    description: Optional[ExampleObjectDescription]
    name: Optional[ExampleObjectName]
from typing import NewType, Union

OneOfExampleObjectReferenceObject5DJ6EmZt = NewType("OneOfExampleObjectReferenceObject5DJ6EmZt", Union[ExampleObject, ReferenceObject])
from typing import List, NewType

ExamplePairingObjectParams = NewType("ExamplePairingObjectParams", List[OneOfExampleObjectReferenceObject5DJ6EmZt])
from typing import NewType, Union

ExamplePairingObjectresult = NewType("ExamplePairingObjectresult", Union[ExampleObject, ReferenceObject])
from typing import TypedDict, Optional

class ExamplePairingObject(TypedDict):
    name: Optional[ExamplePairingObjectName]
    description: Optional[ExamplePairingObjectDescription]
    params: Optional[ExamplePairingObjectParams]
    result: Optional[ExamplePairingObjectresult]
from typing import NewType, Union

OneOfExamplePairingObjectReferenceObjectWEBfRSyK = NewType("OneOfExamplePairingObjectReferenceObjectWEBfRSyK", Union[ExamplePairingObject, ReferenceObject])
from typing import List, NewType

MethodObjectExamples = NewType("MethodObjectExamples", List[OneOfExamplePairingObjectReferenceObjectWEBfRSyK])
from typing import NewType

MethodObjectDeprecated = NewType("MethodObjectDeprecated", bool)
from typing import TypedDict, Optional

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
from typing import List, NewType

Methods = NewType("Methods", List[MethodObject])
from typing import NewType, Any, Mapping

SchemaComponents = NewType("SchemaComponents", Mapping[Any, Any])
from typing import NewType, Any, Mapping

LinkComponents = NewType("LinkComponents", Mapping[Any, Any])
from typing import NewType, Any, Mapping

ObjectTfFA84LI = NewType("ObjectTfFA84LI", Mapping[Any, Any])
from typing import NewType, Any, Mapping

ExampleComponents = NewType("ExampleComponents", Mapping[Any, Any])
from typing import NewType, Any, Mapping

ExamplePairingComponents = NewType("ExamplePairingComponents", Mapping[Any, Any])
from typing import NewType, Any, Mapping

ContentDescriptorComponents = NewType("ContentDescriptorComponents", Mapping[Any, Any])
from typing import NewType, Any, Mapping

TagComponents = NewType("TagComponents", Mapping[Any, Any])
from typing import TypedDict, Optional

class Components(TypedDict):
    schemas: Optional[SchemaComponents]
    links: Optional[LinkComponents]
    errors: Optional[ObjectTfFA84LI]
    examples: Optional[ExampleComponents]
    examplePairings: Optional[ExamplePairingComponents]
    contentDescriptors: Optional[ContentDescriptorComponents]
    tags: Optional[TagComponents]
from typing import TypedDict, Optional

class OpenrpcDocument(TypedDict):
    openrpc: undefined
    info: Optional[InfoObject]
    externalDocs: Optional[ExternalDocumentationObject]
    servers: Optional[Servers]
    methods: undefined
    components: Optional[Components]