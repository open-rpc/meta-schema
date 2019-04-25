const metaSchema = require('../schema.json');
const openRpcExamples = require('@open-rpc/examples').default;
const fetch = require('node-fetch');
const Ajv = require('ajv');
const ajv = new Ajv();

const getJsonSchemaDraft7 = async () => ajv.addMetaSchema(
  await fetch("https://json-schema.org/draft-07/schema").then((res:any) => res.json()),
  "https://json-schema.org/draft-07/schema#"
);

describe('validates all examples without error', () => {
  beforeAll(async () => await getJsonSchemaDraft7());

  const exampleNames:string[] = Object.keys(openRpcExamples);
  exampleNames.forEach((exampleName:string) => {
    it(`validates the example: ${exampleName}`, () => {
      const result:boolean = ajv.validate(metaSchema, openRpcExamples[exampleName]);
      if (ajv.errors && ajv.errors.length > 0) {
        console.error(ajv.errors);
      }
      expect(ajv.errors).toEqual(null);
      expect(result).toBe(true);
    });
  });
});
