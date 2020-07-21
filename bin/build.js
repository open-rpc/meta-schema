#!/usr/bin/env node

const schema = require('../schema.json');
const fs = require('fs');
const path = require('path');
const { promisify } = require('util');
const writeFile = promisify(fs.writeFile);
const { ensureDir } = require('fs-extra');
const {listReleases} = require("@etclabscore/dl-github-releases");

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
};

module.exports = {setOpenRPCVersionEnum};
if (require.main === module) {
  build();
}
