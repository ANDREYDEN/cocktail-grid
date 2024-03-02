import { generateFetchers, generateSchemaTypes } from "@openapi-codegen/typescript";
import { defineConfig } from "@openapi-codegen/cli";
export default defineConfig({
  cocktailGrid: {
    from: {
      relativePath: "../Backend/docs/swagger.json",
      source: "file",
    },
    outputDir: "./src/openapi",
    to: async (context) => {
      const { schemasFiles } = await generateSchemaTypes(context, {
        filenamePrefix: "cocktailGrid",
      });
      await generateFetchers(context, {
        filenamePrefix: "cocktailGrid",
        schemasFiles
      })
    },
  },
});
