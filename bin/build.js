#!/usr/bin/env node

const schema = require('../schema.json');
const fs = require('fs');
const path = require('path');
const { promisify } = require('util');
const writeFile = promisify(fs.writeFile);
const { ensureDir } = require('fs-extra');
const {listReleases} = require("@etclabscore/dl-github-releases");
const { quicktype } = require("quicktype");
const _ = require("lodash");
const { toSafeString } = require("json-schema-to-typescript/dist/src/utils");
const refParser = require("json-schema-ref-parser");
// errors if you try to run with $ref to draft 7 json schema
schema.definitions.schema.$ref = undefined;
const schemaToRef = (s) => ({ $ref: `#/definitions/${s.title}` });
const collectAndRefSchemas = (schema) => {
  const newS = { ...schema };
  const subS = [];

  if (schema.anyOf) {
    subS.push(schema.anyOf);
    newS.anyOf = schema.anyOf.map(schemaToRef);
  }

  if (schema.allOf) {
    subS.push(schema.allOf);
    newS.allOf = schema.allOf.map(schemaToRef);
  }

  if (schema.oneOf) {
    subS.push(schema.oneOf);
    newS.oneOf = schema.oneOf.map(schemaToRef);
  }


  if (schema.items) {
    subS.push(schema.items);

    if (schema.items instanceof Array) {
      newS.items = schema.items.map(schemaToRef);
    } else {
      newS.items = schemaToRef(schema.items);
    }
  }

  if (schema.properties) {
    subS.push(Object.values(schema.properties));
    newS.properties = _.mapValues(schema.properties, schemaToRef);
  }

  const subSchemas = _.chain(subS)
    .flatten()
    .compact()
    .value();

  const collectedSubSchemas = _.map(subSchemas, collectAndRefSchemas);

  return _.chain(collectedSubSchemas)
    .push([newS])
    .flatten()
    .value();
};
const isComment = (line) => {
  const trimmed = line.trim();
  return _.startsWith(trimmed, "/**") || _.startsWith(trimmed, "*") || _.startsWith(trimmed, "*/");
};

const getDefs = (lines) => {
  let commentBuffer = [];
  return _.chain(lines.split("\n"))
    .reduce((memoLines, line) => {
      const lastItem = _.last(memoLines);
      const singleLine = line.match(/export (.*);/);

      if (isComment(line)) {
        commentBuffer.push(line);
      } else if (singleLine) {
        memoLines.push([...commentBuffer, line]);
        commentBuffer = [];
      } else {
        const isStruct = line.match(/export (.*)/);
        if (isStruct) {
          memoLines.push([...commentBuffer, line]);
          commentBuffer = [];
        } else if (_.isArray(lastItem)) {
          lastItem.push(line);
          if (line === "}") {
            memoLines.push("");
          }
        }
      }

      return memoLines;
    }, [])
    .compact()
    .uniqBy((exportLine) => {
      const toTest = exportLine instanceof Array ?
        _.reject(exportLine, isComment)[0] : exportLine;
      const [all, exportType, name, rest] = toTest.match(/export\s(type|interface|enum)\s(\S*)/);
      return name;
    })
    .flattenDeep()
    .join("\n")
    .value();
};

const getQuickTypeSources = async (s) => {
  const parsed = await refParser.dereference(s);
  delete parsed.definitions
  delete parsed.$id
  delete parsed.$schema

  const allSchemas = _.chain(collectAndRefSchemas(parsed))
        .uniqBy("title")
        .map((schema, i, collection) => ({ ...schema, definitions: _.keyBy(collection, (ss) => ss.title) }))
        .value();

  const rawTypes = await Promise.all(allSchemas.map(async (allS) => {
    const typingsForSchema = await quicktype({
      lang: "typescript",
      leadingComments: undefined,
      rendererOptions: { "just-types": "true", "explicit-unions": "true"},
      sources: [{
        kind: "schema",
        name: allS.title,
        schema: JSON.stringify(allS),
      }],
    });
    console.log(_.omit(allS,["definitions"]))

    return typingsForSchema.lines;
  }));

  console.log("finished generating all the types");

  const types = _.chain(rawTypes)
        .flatten()
        .compact()
        .map(_.trimEnd)
        .join("\n")
        .trim()
        .value();

  return getDefs(types);
};
  // const openrpcDef = _.omit(s, ["definitions", "$schema", "$id"]);
  // const defs = {
  //   OpenRPC: openrpcDef,
  //   ...s.definitions
  // };

  // const _s = {
  //   title: "AllTypes",
  //   ..._.pick(s, ["$schema", "$id"]),
  //   type: "object",
  //   properties: _.mapValues(defs, (s) => { $ref: `#/definitions/${s.title}` }),
  //   definitions: defs
  // };
  // console.log(JSON.stringify(_s, undefined, "  "));
  // return [{
  //   kind: "schema",
  //   name: "OpenRPC",
  //   schema: JSON.stringify(_s),
  // }];

const generateTypes = async (s) => {
  const ts = await getQuickTypeSources(s);
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
