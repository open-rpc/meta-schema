import * as fs from "fs";

describe("meta-schema", () => {

  it("is valid json", () => {
    const schema = fs.readFileSync("./schema.json", "utf8");
    const asJson = JSON.parse(schema);
    expect(asJson.type).toBe("object");
  });
});
