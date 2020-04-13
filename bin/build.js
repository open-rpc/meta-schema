#!/usr/bin/env node

const schema = require('../schema.json');
const fs = require('fs');
const path = require('path');
const { promisify } = require('util');
const writeFile = promisify(fs.writeFile);
const { ensureDir } = require('fs-extra');
const {listReleases} = require("@etclabscore/dl-github-releases");
const {JsonSchemaToTypes} = require("@etclabscore/json-schema-to-types");
const refParser = require("json-schema-ref-parser");

const generateTypes = async (s) => {
  const escapedS = JSON.stringify(s).replace(/"/g, "\\\"");
  const parsed = await refParser.dereference(s);
  // the title set is particularly ugly, so we set a new one
  parsed.definitions.schema.title = "JSONSchema";

  // we must fix a bug with the deref util...
  parsed.definitions.contentDescriptorObject.properties.schema = parsed.definitions.schema;

  const transpiler = new JsonSchemaToTypes(parsed);
  const ts = transpiler.toTs();

  // Check if our generated macros use json.Un/Marshal.
  // If so, add the standard json package as an import.
  const transpiled = transpiler.toGo();
  var imports = "";
  if (/json\.(Un|)Marshal/g.test(transpiled)) {
    imports += `import "encoding/json"`
  }

  const go = [
    "package meta_schema",
    "",
    `${imports}`,
    "",
    `const MetaSchemaJSON = "${escapedS}"`,
    "",
    transpiled,
  ].join("\n");

  const dir = path.resolve(__dirname, "../build/src/");
  const goDir = path.resolve(__dirname, "../");
  await ensureDir(dir);
  await ensureDir(goDir);
  await writeFile(`${dir}/index.d.ts`, ts, "utf8");
  await writeFile(`${goDir}/meta_schema.go`, go, "utf8");

  console.log("Generating types complete!");
};

const setOpenRPCVersionEnum = async (s) => {
  s.properties.openrpc.enum = await listReleases("open-rpc", "spec");
  return s;
};

const build = async () => {
  const withVersionEnum = await setOpenRPCVersionEnum(schema);

  const dir = path.resolve(__dirname, "../build/");
  await ensureDir(dir);
  await writeFile(`${dir}/schema.json`, JSON.stringify(withVersionEnum, undefined, "  "));
  console.log("wrote schema.json");

  try {
    await generateTypes(withVersionEnum);
  } catch (e) {
    console.error(e);
    process.exit(1);
  }


  console.log("Finished building");
};

module.exports = {setOpenRPCVersionEnum};
if (require.main === module) {
  build();
}
