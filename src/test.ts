import * as fs from "fs";
import Dereferencer from "@json-schema-tools/dereferencer";
import Transpiler from "@json-schema-tools/transpiler";
import { JSONSchema, JSONSchemaObject } from "..";

describe("meta-schema", () => {
  const s = fs.readFileSync("./schema.json", "utf8");
  let schema: JSONSchema = false;
  beforeEach(async () => {
    schema = JSON.parse(s);
  });


  it("is valid json", () => {
    expect(schema).toBeTruthy();
  });

  it("has references that are valid", async () => {
    expect.assertions(5);
    const deref = new Dereferencer(schema);
    const dereffed = await deref.resolve();
    expect(dereffed).toBeTruthy();
    expect(dereffed.definitions).not.toBeDefined();
    expect(dereffed.properties.methods.items.oneOf[0].properties.result.oneOf[0].properties.schema.title).toBe("JSONSchema");
    expect(dereffed.properties.methods.items.oneOf[0].properties.result.oneOf[1].title).toBe("referenceObject");
    expect(dereffed.properties.methods.items.oneOf[0].properties.result.oneOf[1].properties.$ref.title).toBe("$ref");
  });


  describe('transpiled output', () => {
    let dereffed: JSONSchema;

    beforeEach(async () => {
      const deref = new Dereferencer(schema);
      dereffed = await deref.resolve();
    });

    it('Doesnt give any errors', () => {
      const transpiler = new Transpiler(dereffed);
      const ts = transpiler.toTs();
      expect(typeof ts).toBe('string');
    });

    it('doesnt contain the name "Undefined"', () => {
      const transpiler = new Transpiler({...dereffed as JSONSchemaObject});
      const ts = transpiler.toTs();
      expect(ts.indexOf('Undefined')).toBe(-1);
    });
  });
});
