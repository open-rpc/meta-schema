const metaSchema = require('../schema.json');
const openRpcExamples = require('@open-rpc/examples').default;
const fetch = require('node-fetch');
const Ajv = require('ajv');
const ajv = new Ajv();
const {cloneDeep} = require("lodash");

const testOpenRPCDocument = {
  info: {
    title: "jipperjobber",
    version: "3.2.1",
  },
  methods: [
    {
      name: "jibber",
      params: [
        {
          name: "jibberNiptip",
          schema: { title: "niptip", type: "number" },
        },
      ],
      result: {
        name: "jibberRipslip",
        schema: {
          properties: {
            reepadoop: { type: "number" },
            skeepadeep: { title: "skeepadeep", type: "integer" },
          },
          title: "ripslip",
          type: "object"
        },
      },
    },
  ],
  openrpc: "1.0.0",
};

const getJsonSchemaDraft7 = async () => ajv.addMetaSchema(
  await fetch("https://json-schema.org/draft-07/schema").then((res) => res.json()),
  "https://json-schema.org/draft-07/schema#"
);

describe('validates all examples without error', () => {
  beforeAll(async () => await getJsonSchemaDraft7());

  it("can validate a simple document", () => {
    const result = ajv.validate(metaSchema, testOpenRPCDocument);
    expect(result).toBe(true);
  });

  it("can invalidate a simple document", () => {
    const copy = cloneDeep(testOpenRPCDocument);
    copy.methods[0].params[0].name = undefined;
    const result = ajv.validate(metaSchema, copy);
    expect(result).toBe(false);
  });

  it("can validate a document with ref to the metaschema itself", () => {
    const copy = cloneDeep(testOpenRPCDocument);
    copy.methods[0].params[0].schema = metaSchema;
    const result = ajv.validate(metaSchema, true);
    expect(result).toBe(false);
  });

  const exampleNames = Object.keys(openRpcExamples);
  exampleNames.forEach((exampleName) => {
    it(`validates the example: ${exampleName}`, () => {
      const result = ajv.validate(metaSchema, openRpcExamples[exampleName]);
      if (ajv.errors && ajv.errors.length > 0) {
        console.error(ajv.errors);
      }
      expect(ajv.errors).toEqual(null);
      expect(result).toBe(true);
    });
  });
});
