const schema = require('../schema.json');
const { compile } = require('json-schema-to-typescript');
const fs = require('fs');
const path = require('path');
const { promisify } = require('util');
const writeFile = promisify(fs.writeFile);
const { ensureDir } = require('fs-extra');
// errors if you try to run with $ref to draft 7 json schema
schema.definitions.schema.$ref = undefined;

const go = async () => {
  const ts = await compile(schema, 'OpenRPC');
  const dir = path.resolve(__dirname, '../build/src/');
  await ensureDir(dir);
  await writeFile(`${dir}/index.d.ts`, ts, 'utf8');

  console.log('Generating types complete!');
};
go();
