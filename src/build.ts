import * as fs from "fs";
import Dereferencer from "@json-schema-tools/dereferencer";
import {
  FieldDef,
  TypeInfo,
  Section,
  renderSections,
  toTitleCase,
} from "./util";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const getDereffedSchema = async (): Promise<any> => {
  const s = fs.readFileSync("./schema.json", "utf8");
  let schema = false;
  try {
    schema = JSON.parse(s);
    const deref = new Dereferencer(schema);
    const dereffed = await deref.resolve();
    return dereffed;
  } catch (e) {
    console.error(
      `Cannot parse meta-schema or dereference it. Recieved: ${s}`
    );
  }
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const extractFields = (schema: any): FieldDef[] => {
  const properties = schema.properties || {};
  const requiredFields: string[] = schema.required || [];

  return Object.entries(properties).map(([name, propSchema]) => {
    const ps = propSchema as TypeInfo & { description?: string };
    return {
      name,
      schema: {
        title: ps.title,
        type: ps.type,
        items: ps.items,
        oneOf: ps.oneOf,
        properties: ps.properties,
        patternProperties: ps.patternProperties,
      },
      description: ps.description,
      required: requiredFields.includes(name),
    };
  });
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const collectSections = (
  schema: any,
  depth: number,
  visited: Set<string>
): Section[] => {
  const title = schema.title || "object";
  if (visited.has(title)) return [];
  visited.add(title);

  const fields = extractFields(schema);
  if (fields.length === 0) return [];

  // Cap at h6 for markdown
  const headingDepth = Math.min(depth + 1, 6) as 1 | 2 | 3 | 4 | 5 | 6;

  // Start with current section
  const sections: Section[] = [
    {
      title: toTitleCase(title),
      parentTitle: title,
      fields,
      headingDepth,
      description: schema.description,
      hasSpecExtensions: schema.patternProperties?.["^x-"] !== undefined,
    },
  ];

  // Depth-first: immediately recurse into each child before moving to next sibling
  for (const field of fields) {
    const propSchema = schema.properties?.[field.name];

    // Direct object with properties
    if (propSchema?.properties) {
      sections.push(...collectSections(propSchema, depth + 1, visited));
    }
    // Array items with properties
    else if (propSchema?.items?.properties) {
      sections.push(...collectSections(propSchema.items, depth + 1, visited));
    }
    // Array items with oneOf
    else if (propSchema?.items?.oneOf) {
      for (const variant of propSchema.items.oneOf) {
        if (variant?.properties) {
          sections.push(...collectSections(variant, depth + 1, visited));
        }
      }
    }
    // Direct oneOf - traverse each variant
    else if (propSchema?.oneOf) {
      for (const variant of propSchema.oneOf) {
        if (variant?.properties) {
          sections.push(...collectSections(variant, depth + 1, visited));
        }
      }
    }

    // patternProperties - traverse values that are objects (skip ^x- extensions)
    if (propSchema?.patternProperties) {
      for (const [pattern, patternValue] of Object.entries(
        propSchema.patternProperties
      )) {
        if (pattern === "^x-") continue;
        const pv = patternValue as any;
        if (pv?.properties) {
          sections.push(...collectSections(pv, depth + 1, visited));
        }
      }
    }
  }
  return sections;
};

const build = async () => {
  const schema = await getDereffedSchema();
  if (!schema) return;

  const sections = collectSections(schema, 0, new Set());
  const markdown = renderSections(sections);

  fs.writeFileSync("./table.md", markdown);
  console.log(markdown);
};


build();
