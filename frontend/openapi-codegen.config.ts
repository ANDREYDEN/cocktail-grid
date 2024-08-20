import {
  generateSchemaTypes,
  generateFetchers,
} from "@openapi-codegen/typescript";
import { defineConfig } from "@openapi-codegen/cli";
export default defineConfig({
  cocktailGrid: {
    from: {
      relativePath: "./src/openapi/swagger.json",
      source: "file",
    },
    outputDir: "./src/openapi",
    to: async (context) => {
      const filenamePrefix = "cocktailGrid";
      const { schemasFiles } = await generateSchemaTypes(context, {
        filenamePrefix,
      });
      await generateFetchers(context, {
        filenamePrefix,
        schemasFiles,
      });
    },
  },
});
