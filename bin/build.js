#!/usr/bin/env node

const schema = require('../schema.json');
const { compile } = require('json-schema-to-typescript');
const fs = require('fs');
const path = require('path');
const { promisify } = require('util');
const writeFile = promisify(fs.writeFile);
const { ensureDir } = require('fs-extra');
const {listReleases} = require("@etclabscore/dl-github-releases");

// errors if you try to run with $ref to draft 7 json schema
schema.definitions.schema.$ref = undefined;

const generateTypes = async () => {
  const ts = await compile(schema, 'OpenRPC');
  const dir = path.resolve(__dirname, '../build/src/');
  await ensureDir(dir);
  await writeFile(`${dir}/index.d.ts`, ts, 'utf8');

  console.log('Generating types complete!');
};

const setOpenRPCVersionEnum = async (s) => {
  s.properties.openrpc.enum = await listReleases("open-rpc", "spec");
  return s;
};

const build = async () => {
  const withVersionEnum = await setOpenRPCVersionEnum(schema);
  await generateTypes(withVersionEnum);
};

build();
