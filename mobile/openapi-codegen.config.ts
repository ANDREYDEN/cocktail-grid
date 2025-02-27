import {
  generateSchemaTypes,
  generateReactQueryComponents,
} from "@openapi-codegen/typescript";
import { defineConfig } from "@openapi-codegen/cli";
export default defineConfig({
  cocktailGrid: {
    from: {
      relativePath: "./openapi/swagger.json",
      source: "file",
    },
    outputDir: "openapi",
    to: async (context) => {
      const filenamePrefix = "cocktailGrid";
      const { schemasFiles } = await generateSchemaTypes(context, {
        filenamePrefix,
      });
      await generateReactQueryComponents(context, {
        filenamePrefix,
        schemasFiles,
      });
    },
  },
});
