# Openrpc Document

| Field Name                                                |                                      Type                                     | Description                                                                                                                                                                                                                                                                                                                                                                 |
| :-------------------------------------------------------- | :---------------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="openrpc-document-openrpc"></a>openrpc            |                                    `string`                                   | **REQUIRED**. This string MUST be the [semantic version number](https://semver.org/spec/v2.0.0.html) of the [OpenRPC Specification version](#versions) that the OpenRPC document uses. The `openrpc` field SHOULD be used by tooling specifications and clients to interpret the OpenRPC document. This is *not* related to the API [`info.version`](#info-version) string. |
| <a name="openrpc-document-info"></a>info                  |                          [Info Object](#info-object)                          | **REQUIRED**. The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.                                                                                                                                                                             |
| <a name="openrpc-document-external-docs"></a>externalDocs |        [External Documentation Object](#external-documentation-object)        | Additional external documentation.                                                                                                                                                                                                                                                                                                                                          |
| <a name="openrpc-document-servers"></a>servers            |                       \[[Server Object](#server-object)]                      | An array of Server Objects, which provide connectivity information to a target server. If the `servers` property is not provided, or is an empty array, the default value would be a [Server Object](#server-object) with a [url](#server-url) value of `localhost`.                                                                                                        |
| <a name="openrpc-document-methods"></a>methods            | \[[Method Object](#method-object) \\\| [Reference Object](#reference-object)] | **REQUIRED**. The available methods for the API. While it is required, the array may be empty (to handle security filtering, for example).                                                                                                                                                                                                                                  |
| <a name="openrpc-document-components"></a>components      |                           [Components](#components)                           | Holds a set of reusable objects for different aspects of the OpenRPC. All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.                                                                                                                                  |
| <a name="openrpc-document-$schema"></a>$schema            |                                    `string`                                   | JSON Schema URI (used by some editors)                                                                                                                                                                                                                                                                                                                                      |

This object MAY be extended with [Specification Extensions](#specification-extensions).

## Info Object

The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.

| Field Name                                         |                Type               | Description                                                                                                                                                         |
| :------------------------------------------------- | :-------------------------------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <a name="info-title"></a>title                     |              `string`             | **REQUIRED**. The title of the application.                                                                                                                         |
| <a name="info-description"></a>description         |              `string`             | A verbose description of the application. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.               |
| <a name="info-terms-of-service"></a>termsOfService |              `string`             | A URL to the Terms of Service for the API. MUST be in the format of a URL.                                                                                          |
| <a name="info-version"></a>version                 |              `string`             | **REQUIRED**. The version of the OpenRPC document (which is distinct from the [OpenRPC Specification version](#openrpc-version) or the API implementation version). |
| <a name="info-contact"></a>contact                 | [Contact Object](#contact-object) | Contact information for the exposed API.                                                                                                                            |
| <a name="info-license"></a>license                 | [License Object](#license-object) | License information for the exposed API.                                                                                                                            |

This object MAY be extended with [Specification Extensions](#specification-extensions).

### Contact Object

Contact information for the exposed API.

| Field Name                        |   Type   | Description                                                                                      |
| :-------------------------------- | :------: | :----------------------------------------------------------------------------------------------- |
| <a name="contact-name"></a>name   | `string` | The identifying name of the contact person/organization.                                         |
| <a name="contact-email"></a>email | `string` | The email address of the contact person/organization. MUST be in the format of an email address. |
| <a name="contact-url"></a>url     | `string` | The URL pointing to the contact information. MUST be in the format of a URL.                     |

This object MAY be extended with [Specification Extensions](#specification-extensions).

### License Object

License information for the exposed API.

| Field Name                      |   Type   | Description                                                            |
| :------------------------------ | :------: | :--------------------------------------------------------------------- |
| <a name="license-name"></a>name | `string` | The license name used for the API.                                     |
| <a name="license-url"></a>url   | `string` | A URL to the license used for the API. MUST be in the format of a URL. |

This object MAY be extended with [Specification Extensions](#specification-extensions).

## External Documentation Object

Additional external documentation.

| Field Name                                                   |   Type   | Description                                                                                                                                             |
| :----------------------------------------------------------- | :------: | :------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <a name="external-documentation-description"></a>description | `string` | A verbose explanation of the documentation. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation. |
| <a name="external-documentation-url"></a>url                 | `string` | **REQUIRED**. The URL for the target documentation. Value MUST be in the format of a URL.                                                               |

This object MAY be extended with [Specification Extensions](#specification-extensions).

## Server Object

A object representing a Server

| Field Name                                   |                                Type                               | Description                                                                                                                                                                                                                                                                                                                             |
| :------------------------------------------- | :---------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="server-url"></a>url                 |                              `string`                             | **REQUIRED**. A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenRPC document is being served. [Server Variables](#server-variables) are passed into the [Runtime Expression](#runtime-expression) to produce a server URL. |
| <a name="server-name"></a>name               |                              `string`                             | An optional string describing the name of the server. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.                                                                                                                                                                       |
| <a name="server-description"></a>description |                              `string`                             | An optional string describing the host designated by the URL. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.                                                                                                                                                               |
| <a name="server-summary"></a>summary         |                              `string`                             | A short summary of what the server is.                                                                                                                                                                                                                                                                                                  |
| <a name="server-variables"></a>variables     | Map\[`string`, [Server Object Variable](#server-object-variable)] | A map between a variable name and its value. The value is passed into the [Runtime Expression](#runtime-expression) to produce a server URL.                                                                                                                                                                                            |

This object MAY be extended with [Specification Extensions](#specification-extensions).

### Server Object Variable

An object representing a Server Variable for server URL template substitution.

| Field Name                                                   |     Type    | Description                                                                                                                                                                                                                                                                         |
| :----------------------------------------------------------- | :---------: | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="server-object-variable-default"></a>default         |   `string`  | **REQUIRED**. The default value to use for substitution, which SHALL be sent if an alternate value is *not* supplied. Note this behavior is different than the [Schema Object's](#schema-object) treatment of default values, because in those cases parameter values are optional. |
| <a name="server-object-variable-description"></a>description |   `string`  | An optional description for the server variable. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.                                                                                                                        |
| <a name="server-object-variable-enum"></a>enum               | \[`string`] | An enumeration of string values to be used if the substitution options are from a limited set.                                                                                                                                                                                      |

## Method Object

Describes the interface for the given method name. The method name is used as the `method` field of the JSON-RPC body. It therefore MUST be unique.

| Field Name                                          |                                                  Type                                                 | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| :-------------------------------------------------- | :---------------------------------------------------------------------------------------------------: | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="method-name"></a>name                      |                                                `string`                                               | **REQUIRED**. The cannonical name for the method. The name MUST be unique within the methods array.                                                                                                                                                                                                                                                                                                                                                                                              |
| <a name="method-description"></a>description        |                                                `string`                                               | A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.                                                                                                                                                                                                                                                                                                                                                                          |
| <a name="method-summary"></a>summary                |                                                `string`                                               | A short summary of what the method does.                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| <a name="method-servers"></a>servers                |                                   \[[Server Object](#server-object)]                                  | An array of Server Objects, which provide connectivity information to a target server. If the `servers` property is not provided, or is an empty array, the default value would be a [Server Object](#server-object) with a [url](#server-url) value of `localhost`.                                                                                                                                                                                                                             |
| <a name="method-tags"></a>tags                      |                \[[Tag Object](#tag-object) \\\| [Reference Object](#reference-object)]                | A list of tags for API documentation control. Tags can be used for logical grouping of methods by resources or any other qualifier.                                                                                                                                                                                                                                                                                                                                                              |
| <a name="method-param-structure"></a>paramStructure |                                                `string`                                               | Format the server expects the params. Defaults to 'either'.                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| <a name="method-params"></a>params                  | \[[Content Descriptor Object](#content-descriptor-object) \\\| [Reference Object](#reference-object)] | **REQUIRED**. A list of parameters that are applicable for this method. The list MUST NOT include duplicated parameters and therefore require [name](#content-descriptor-name) to be unique. The list can use the [Reference Object](#reference-object) to link to parameters that are defined by the [Content Descriptor Object](#content-descriptor-object). All optional params (content descriptor objects with "required": false) MUST be positioned after all required params in the list. |
| <a name="method-result"></a>result                  |   [Content Descriptor Object](#content-descriptor-object) \\\| [Reference Object](#reference-object)  | The description of the result returned by the method. If defined, it MUST be a Content Descriptor or Reference Object. If undefined, the method MUST only be used as a [notification](https://www.jsonrpc.org/specification#notification)                                                                                                                                                                                                                                                        |
| <a name="method-errors"></a>errors                  |              \[[Error Object](#error-object) \\\| [Reference Object](#reference-object)]              | A list of custom application defined errors that MAY be returned. The Errors MUST have unique error codes.                                                                                                                                                                                                                                                                                                                                                                                       |
| <a name="method-links"></a>links                    |               \[[Link Object](#link-object) \\\| [Reference Object](#reference-object)]               | A list of possible links from this method call.                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| <a name="method-examples"></a>examples              |    \[[Example Pairing Object](#example-pairing-object) \\\| [Reference Object](#reference-object)]    | Array of [Example Pairing Objects](#example-pairing-object) where each example includes a valid params-to-result [Content Descriptor](#content-descriptor-object) pairing.                                                                                                                                                                                                                                                                                                                       |
| <a name="method-deprecated"></a>deprecated          |                                               `boolean`                                               | Declares this method to be deprecated. Consumers SHOULD refrain from usage of the declared method. Default value is `false`.                                                                                                                                                                                                                                                                                                                                                                     |
| <a name="method-external-docs"></a>externalDocs     |                    [External Documentation Object](#external-documentation-object)                    | Additional external documentation.                                                                                                                                                                                                                                                                                                                                                                                                                                                               |

This object MAY be extended with [Specification Extensions](#specification-extensions).

### Tag Object

Adds metadata to a single tag that is used by the [Method Object](#method-object). It is not mandatory to have a Tag Object per tag defined in the Method Object instances.

| Field Name                                   |                               Type                              | Description                                                                                                                                    |
| :------------------------------------------- | :-------------------------------------------------------------: | :--------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="tag-name"></a>name                  |                             `string`                            | **REQUIRED**. The name of the tag.                                                                                                             |
| <a name="tag-description"></a>description    |                             `string`                            | A verbose explanation for the tag. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation. |
| <a name="tag-external-docs"></a>externalDocs | [External Documentation Object](#external-documentation-object) | Additional external documentation for this tag.                                                                                                |

This object MAY be extended with [Specification Extensions](#specification-extensions).

### Reference Object

| Field Name                        |   Type   | Description                         |
| :-------------------------------- | :------: | :---------------------------------- |
| <a name="reference-$ref"></a>$ref | `string` | **REQUIRED**. The reference string. |

### Content Descriptor Object

Content Descriptors are objects that do just as they suggest - describe content. They are reusable ways of describing either parameters or result. They MUST have a schema.

| Field Name                                               |                          Type                          | Description                                                                                                                                                                                                       |
| :------------------------------------------------------- | :----------------------------------------------------: | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="content-descriptor-name"></a>name               |                        `string`                        | **REQUIRED**. Name of the content that is being described. If the content described is a method parameter assignable [`by-name`](#method-param-structure), this field SHALL define the parameter's key (ie name). |
| <a name="content-descriptor-description"></a>description |                        `string`                        | A verbose explanation of the content descriptor behavior. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.                                             |
| <a name="content-descriptor-summary"></a>summary         |                        `string`                        | A short summary of the content that is being described.                                                                                                                                                           |
| <a name="content-descriptor-schema"></a>schema           | [JSONSchema Object](#jsonschema-object) \\\| `boolean` | **REQUIRED**. Schema that describes the content.                                                                                                                                                                  |
| <a name="content-descriptor-required"></a>required       |                        `boolean`                       | Determines if the content is a required field. Default value is `false`.                                                                                                                                          |
| <a name="content-descriptor-deprecated"></a>deprecated   |                        `boolean`                       | Specifies that the content is deprecated and SHOULD be transitioned out of usage. Default value is `false`.                                                                                                       |

This object MAY be extended with [Specification Extensions](#specification-extensions).

#### JSONSchema Object

| Field Name                                                          |                            Type                           | Description |
| :------------------------------------------------------------------ | :-------------------------------------------------------: | :---------- |
| <a name="jsonschema-$id"></a>$id                                    |                          `string`                         |             |
| <a name="jsonschema-$schema"></a>$schema                            |                          `string`                         |             |
| <a name="jsonschema-$ref"></a>$ref                                  |                          `string`                         |             |
| <a name="jsonschema-$comment"></a>$comment                          |                          `string`                         |             |
| <a name="jsonschema-title"></a>title                                |                          `string`                         |             |
| <a name="jsonschema-description"></a>description                    |                          `string`                         |             |
| <a name="jsonschema-default"></a>default                            |                          unknown                          |             |
| <a name="jsonschema-read-only"></a>readOnly                         |                         `boolean`                         |             |
| <a name="jsonschema-examples"></a>examples                          |                         \[unknown]                        |             |
| <a name="jsonschema-multiple-of"></a>multipleOf                     |                          `number`                         |             |
| <a name="jsonschema-maximum"></a>maximum                            |                          `number`                         |             |
| <a name="jsonschema-exclusive-maximum"></a>exclusiveMaximum         |                          `number`                         |             |
| <a name="jsonschema-minimum"></a>minimum                            |                          `number`                         |             |
| <a name="jsonschema-exclusive-minimum"></a>exclusiveMinimum         |                          `number`                         |             |
| <a name="jsonschema-max-length"></a>maxLength                       |                         `integer`                         |             |
| <a name="jsonschema-min-length"></a>minLength                       |                         `integer`                         |             |
| <a name="jsonschema-pattern"></a>pattern                            |                          `string`                         |             |
| <a name="jsonschema-additional-items"></a>additionalItems           |   [JSONSchema Object](#jsonschema-object) \\\| `boolean`  |             |
| <a name="jsonschema-items"></a>items                                |                      [Items](#items)                      |             |
| <a name="jsonschema-max-items"></a>maxItems                         |                         `integer`                         |             |
| <a name="jsonschema-min-items"></a>minItems                         |                         `integer`                         |             |
| <a name="jsonschema-unique-items"></a>uniqueItems                   |                         `boolean`                         |             |
| <a name="jsonschema-contains"></a>contains                          |   [JSONSchema Object](#jsonschema-object) \\\| `boolean`  |             |
| <a name="jsonschema-max-properties"></a>maxProperties               |                         `integer`                         |             |
| <a name="jsonschema-min-properties"></a>minProperties               |                         `integer`                         |             |
| <a name="jsonschema-required"></a>required                          |                        \[`string`]                        |             |
| <a name="jsonschema-additional-properties"></a>additionalProperties |   [JSONSchema Object](#jsonschema-object) \\\| `boolean`  |             |
| <a name="jsonschema-definitions"></a>definitions                    |                          `object`                         |             |
| <a name="jsonschema-properties"></a>properties                      |                          `object`                         |             |
| <a name="jsonschema-pattern-properties"></a>patternProperties       |                          `object`                         |             |
| <a name="jsonschema-dependencies"></a>dependencies                  |                          `object`                         |             |
| <a name="jsonschema-property-names"></a>propertyNames               |   [JSONSchema Object](#jsonschema-object) \\\| `boolean`  |             |
| <a name="jsonschema-const"></a>const                                |                          unknown                          |             |
| <a name="jsonschema-enum"></a>enum                                  |                         \[unknown]                        |             |
| <a name="jsonschema-type"></a>type                                  |                       [Type](#type)                       |             |
| <a name="jsonschema-format"></a>format                              |                          `string`                         |             |
| <a name="jsonschema-content-media-type"></a>contentMediaType        |                          `string`                         |             |
| <a name="jsonschema-content-encoding"></a>contentEncoding           |                          `string`                         |             |
| <a name="jsonschema-if"></a>if                                      |   [JSONSchema Object](#jsonschema-object) \\\| `boolean`  |             |
| <a name="jsonschema-then"></a>then                                  |   [JSONSchema Object](#jsonschema-object) \\\| `boolean`  |             |
| <a name="jsonschema-else"></a>else                                  |   [JSONSchema Object](#jsonschema-object) \\\| `boolean`  |             |
| <a name="jsonschema-all-of"></a>allOf                               | \[[JSONSchema Object](#jsonschema-object) \\\| `boolean`] |             |
| <a name="jsonschema-any-of"></a>anyOf                               | \[[JSONSchema Object](#jsonschema-object) \\\| `boolean`] |             |
| <a name="jsonschema-one-of"></a>oneOf                               | \[[JSONSchema Object](#jsonschema-object) \\\| `boolean`] |             |
| <a name="jsonschema-not"></a>not                                    |   [JSONSchema Object](#jsonschema-object) \\\| `boolean`  |             |

### Error Object

Defines an application level error.

| Field Name                          |                   Type                  | Description                                                                                                                                                                                                                                                            |
| :---------------------------------- | :-------------------------------------: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="error-code"></a>code       |                `integer`                | **REQUIRED**. A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api. |
| <a name="error-message"></a>message |                 `string`                | **REQUIRED**. A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.                                                                                                                                         |
| <a name="error-data"></a>data       | [Error Object Data](#error-object-data) | A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).                                                    |

### Link Object

A object representing a Link

| Field Name                                 |                    Type                   | Description                                                                                                                                                                                                                                                 |
| :----------------------------------------- | :---------------------------------------: | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="link-name"></a>name               |   [Link Object Name](#link-object-name)   | Cannonical name of the link.                                                                                                                                                                                                                                |
| <a name="link-summary"></a>summary         |                  `string`                 | Short description for the link.                                                                                                                                                                                                                             |
| <a name="link-method"></a>method           |                  `string`                 | The name of an existing, resolvable OpenRPC method, as defined with a unique `method`. This field MUST resolve to a unique [Method Object](#method-object). As opposed to Open Api, Relative `method` values ARE NOT permitted.                             |
| <a name="link-description"></a>description |                  `string`                 | A description of the link. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.                                                                                                                      |
| <a name="link-params"></a>params           | [Link Object Params](#link-object-params) | A map representing parameters to pass to a method as specified with `method`. The key is the parameter name to be used, whereas the value can be a constant or a [runtime expression](#runtime-expression) to be evaluated and passed to the linked method. |
| <a name="link-server"></a>server           | [Link Object Server](#link-object-server) | A server object to be used by the target method.                                                                                                                                                                                                            |

This object MAY be extended with [Specification Extensions](#specification-extensions).

#### Link Object Server

A server object to be used by the target method.

| Field Name                                               |                                Type                               | Description                                                                                                                                                                                                                                                                                                                             |
| :------------------------------------------------------- | :---------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="link-object-server-url"></a>url                 |                              `string`                             | **REQUIRED**. A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenRPC document is being served. [Server Variables](#server-variables) are passed into the [Runtime Expression](#runtime-expression) to produce a server URL. |
| <a name="link-object-server-name"></a>name               |                              `string`                             | An optional string describing the name of the server. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.                                                                                                                                                                       |
| <a name="link-object-server-description"></a>description |                              `string`                             | An optional string describing the host designated by the URL. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.                                                                                                                                                               |
| <a name="link-object-server-summary"></a>summary         |                              `string`                             | A short summary of what the server is.                                                                                                                                                                                                                                                                                                  |
| <a name="link-object-server-variables"></a>variables     | Map\[`string`, [Server Object Variable](#server-object-variable)] | A map between a variable name and its value. The value is passed into the [Runtime Expression](#runtime-expression) to produce a server URL.                                                                                                                                                                                            |

This object MAY be extended with [Specification Extensions](#specification-extensions).

### Example Pairing Object

The Example Pairing object consists of a set of example params and result. The result is what you can expect from the JSON-RPC service given the exact params.

| Field Name                                            |                                       Type                                      | Description                                                                                              |
| :---------------------------------------------------- | :-----------------------------------------------------------------------------: | :------------------------------------------------------------------------------------------------------- |
| <a name="example-pairing-name"></a>name               |                                     `string`                                    | **REQUIRED**. Name for the example pairing.                                                              |
| <a name="example-pairing-description"></a>description |                                     `string`                                    | A verbose explanation of the example pairing.                                                            |
| <a name="example-pairing-params"></a>params           | \[[Example Object](#example-object) \\\| [Reference Object](#reference-object)] | **REQUIRED**. Example parameters.                                                                        |
| <a name="example-pairing-result"></a>result           |   [Example Object](#example-object) \\\| [Reference Object](#reference-object)  | Example result. When not provided, the example pairing represents usage of the method as a notification. |

#### Example Object

The Example object is an object that defines an example that is intended to match the `schema` of a given [Content Descriptor](#content-descriptor-object).

| Field Name                                    |                      Type                     | Description                                                                                                                                                                                                                                                      |
| :-------------------------------------------- | :-------------------------------------------: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a name="example-summary"></a>summary         |                    `string`                   | Short description for the example.                                                                                                                                                                                                                               |
| <a name="example-value"></a>value             | [Example Object Value](#example-object-value) | **REQUIRED**. Embedded literal example. The `value` field and `externalValue` field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON, use a string value to contain the example, escaping where necessary. |
| <a name="example-description"></a>description |                    `string`                   | A verbose explanation of the example. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.                                                                                                                |
| <a name="example-name"></a>name               |                    `string`                   | **REQUIRED**. Cannonical name of the example.                                                                                                                                                                                                                    |

This object MAY be extended with [Specification Extensions](#specification-extensions).

## Components

Holds a set of reusable objects for different aspects of the OpenRPC. All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.

| Field Name                                                      |                                   Type                                  | Description                                                                          |
| :-------------------------------------------------------------- | :---------------------------------------------------------------------: | :----------------------------------------------------------------------------------- |
| <a name="components-schemas"></a>schemas                        |  Map\[`string`, [JSONSchema Object](#jsonschema-object) \\\| `boolean`] | An object to hold reusable [Schema Objects](#schema-object).                         |
| <a name="components-links"></a>links                            |               Map\[`string`, [Link Object](#link-object)]               | An object to hold reusable [Link Objects](#link-object).                             |
| <a name="components-errors"></a>errors                          |              Map\[`string`, [Error Object](#error-object)]              | An object to hold reusable [Error Objects](#error-object).                           |
| <a name="components-examples"></a>examples                      |            Map\[`string`, [Example Object](#example-object)]            | An object to hold reusable [Example Objects](#example-object).                       |
| <a name="components-example-pairings"></a>examplePairings       |    Map\[`string`, [Example Pairing Object](#example-pairing-object)]    | An object to hold reusable [Example Pairing Objects](#example-pairing-object).       |
| <a name="components-content-descriptors"></a>contentDescriptors | Map\[`string`, [Content Descriptor Object](#content-descriptor-object)] | An object to hold reusable [Content Descriptor Objects](#content-descriptor-object). |
| <a name="components-tags"></a>tags                              |                Map\[`string`, [Tag Object](#tag-object)]                | An object to hold reusable [Tag Objects](#tag-object).                               |
