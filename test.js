const Ajv = require('ajv');
const ajv = new Ajv();
const metaSchema = require('./schema.json');
const fetch = require('node-fetch');

describe('meta-schema', () => {
  let metaSchemaValidator, examples;
  it('can be compiled by ajv', () => {
    metaSchemaValidator = ajv.compile(metaSchema);
  });

  describe('validates all examples without error', () => {
    it(`validates all the examples properly`, async () => {
      const examples = [
        await fetch('https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/api-with-examples.json').then((res) => res.json()),
        await fetch('https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/link-example.json').then((res) => res.json()),
        await fetch('https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/petstore-expanded.json').then((res) => res.json()),
        await fetch('https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/petstore.json').then((res) => res.json())
      ];
      examples.forEach((example) => {
        const valid = metaSchemaValidator(example.schema);
        expect(valid).toBeTruthy();
      });
    });
  });
});
