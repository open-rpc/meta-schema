import { readJson } from "fs-extra";

describe("meta-schema", () => {

  it("is valid json", async () => {
    const schema = await readJson("./schema.json");
    expect(schema.type).toBe("object");
  });
});
