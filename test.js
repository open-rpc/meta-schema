const Djv = require('djv');
const metaSchema = require('./schema.json');
const fetch = require('node-fetch');
const djv = new Djv();
const openRpcExamples = require('@open-rpc/examples');

describe('meta-schema', () => {
  let metaSchemaValidator, examples;
  it('can be compiled by ajv', () => {
    djv.addSchema('test', metaSchema);
  });

  describe('validates all examples without error', () => {
    const exampleNames = Object.keys(openRpcExamples);
    exampleNames.forEach((exampleName) => {
      it(`validates the example: ${exampleName}`, () => {
        openRpcExamples[exampleName];
        const result = djv.validate(openRpcExamples[exampleName]);
        expect(result).toBeUndefined();
      });
    });
  });
});
