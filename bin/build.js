#!/usr/bin/env node

const schema = require('../schema.json');
const fs = require('fs');
const path = require('path');
const { promisify } = require('util');
const writeFile = promisify(fs.writeFile);
const { ensureDir } = require('fs-extra');
const {listReleases} = require("@etclabscore/dl-github-releases");
const { quicktype } = require("quicktype");

// errors if you try to run with $ref to draft 7 json schema
schema.definitions.schema.$ref = undefined;

const getQuickTypeSources = (s) => {
  return [{
    kind: "schema",
    name: "OpenRPC",
    schema: JSON.stringify(s),
  }];
};

const generateTypes = async (s) => {
  const sources = getQuickTypeSources(s);
  const result = await quicktype({
    lang: "typescript",
    leadingComments: undefined,
    rendererOptions: { "just-types": true },
    sources,
  });
  const ts = result.lines.join("\n");
  const dir = path.resolve(__dirname, "../build/src/");
  await ensureDir(dir);
  await writeFile(`${dir}/index.d.ts`, ts, "utf8");

  console.log("Generating types complete!");
};

const setOpenRPCVersionEnum = async (s) => {
  s.properties.openrpc.enum = await listReleases("open-rpc", "spec");
  return s;
};

const build = async () => {
  const withVersionEnum = await setOpenRPCVersionEnum(schema);

  await generateTypes(withVersionEnum);

  const dir = path.resolve(__dirname, "../build/");
  await ensureDir(dir);

  await writeFile(`${dir}/schema.json`, JSON.stringify(withVersionEnum, undefined, "  "));

  console.log("Finished building");
};

module.exports = {setOpenRPCVersionEnum};
if (require.main === module) {
  build();
}
