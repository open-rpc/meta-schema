#!/usr/bin/env node

const { promisify } = require('util');
const writeFile = promisify(require('fs').writeFile);
const rm = promisify(require('fs').unlink);
const exec = promisify(require('child_process').exec);

const GenerateSchema = require('generate-schema');
const fetch = require('node-fetch');

const getExampleSchemas = () => {
  const exampleBaseUrl = 'https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/';
  const examples = [
    'petstore',
    'petstore-expanded',
    'link-example',
    'api-with-examples'
  ];

  return Promise.all(
    examples
      .map((example) => `${exampleBaseUrl}${example}.json`)
      .map((exampleUrl) => fetch(exampleUrl).then((res) => res.json()))
  );
};


const generate = async () => {
  const exampleSchema = await getExampleSchemas();

  const examplesJsonLines = exampleSchema.map((example) => JSON.stringify(example)).join('\n');
  await writeFile('samples.jsonl', examplesJsonLines);
  const metaSchema = await exec('genson -s schema4.json samples.jsonl');
  console.log(metaSchema.stdout);
  await writeFile('schema4.json', metaSchema.stdout);
  await rm('samples.jsonl');
};

generate().then(() => console.log('done'));
