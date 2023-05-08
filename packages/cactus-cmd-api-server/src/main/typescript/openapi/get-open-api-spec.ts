import { Checks } from "@hyperledger/cactus-common";
import { PluginRegistry } from "@hyperledger/cactus-core";
import {
  ICactusPlugin,
  IPluginWebService,
  isIPluginWebService,
} from "@hyperledger/cactus-core-api";
import type { OpenAPIV3 } from "express-openapi-validator/dist/framework/types";

export async function getOpenApiSpec(
  pluginRegistry: PluginRegistry,
): Promise<OpenAPIV3.Document[]> {
  Checks.truthy(
    pluginRegistry,
    `openApiSpec()
  pluginRegistry (PluginRegistry)`,
  );

  const openApiJsonSpecsPromises = pluginRegistry
    .getPlugins()
    .filter((pluginInstance) => isIPluginWebService(pluginInstance))
    .map(async (plugin: ICactusPlugin) => {
      const webPlugin = plugin as IPluginWebService;
      const openApiSpec = (await webPlugin.getOpenApiSpec()) as OpenAPIV3.Document;
      return openApiSpec;
    });

  const openApiJsonSpecs = await Promise.all(openApiJsonSpecsPromises);
  return openApiJsonSpecs;
}
