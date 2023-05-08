import { Express, Request, Response } from "express";

import {
  Logger,
  Checks,
  LogLevelDesc,
  LoggerProvider,
  IAsyncProvider,
} from "@hyperledger/cactus-common";

import {
  IWebServiceEndpoint,
  IExpressRequestHandler,
  IEndpointAuthzOptions,
} from "@hyperledger/cactus-core-api";

import OAS from "../../json/openapi.json";
import {
  PluginRegistry,
  registerWebServiceEndpoint,
} from "@hyperledger/cactus-core";
import { getOpenApiSpec } from "../openapi/get-open-api-spec";

export interface IGetOpenApiSpecV1Endpoint {
  logLevel?: LogLevelDesc;
  pluginRegistry: PluginRegistry;
}

export class GetOpenApiSpecV1Endpoint implements IWebServiceEndpoint {
  public static readonly CLASS_NAME = "GetOpenApiSpecV1Endpoint";

  private readonly log: Logger;

  public get className(): string {
    return GetOpenApiSpecV1Endpoint.CLASS_NAME;
  }

  constructor(public readonly opts: IGetOpenApiSpecV1Endpoint) {
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(opts, `${fnTag} arg options`);
    Checks.truthy(opts.pluginRegistry, `${fnTag} arg options.pluginRegistry`);

    const level = this.opts.logLevel || "INFO";
    const label = this.className;
    this.log = LoggerProvider.getOrCreate({ level, label });
  }

  public getExpressRequestHandler(): IExpressRequestHandler {
    return this.handleRequest.bind(this);
  }

  public get oasPath(): typeof OAS.paths["/api/v1/api-server/get-open-api-spec-v1-endpoint"] {
    return OAS.paths["/api/v1/api-server/get-open-api-spec-v1-endpoint"];
  }

  public getPath(): string {
    return this.oasPath.post["x-hyperledger-cactus"].http.path;
  }

  public getVerbLowerCase(): string {
    return this.oasPath.post["x-hyperledger-cactus"].http.verbLowerCase;
  }

  public getOperationId(): string {
    return this.oasPath.post.operationId;
  }

  public async registerExpress(
    expressApp: Express,
  ): Promise<IWebServiceEndpoint> {
    await registerWebServiceEndpoint(expressApp, this);
    return this;
  }

  getAuthorizationOptionsProvider(): IAsyncProvider<IEndpointAuthzOptions> {
    // TODO: make this an injectable dependency in the constructor
    return {
      get: async () => ({
        isProtected: true,
        requiredRoles: [],
      }),
    };
  }

  async handleRequest(req: Request, res: Response): Promise<void> {
    const fnTag = `${this.className}#handleRequest()`;
    const verbUpper = this.getVerbLowerCase().toUpperCase();
    this.log.debug(`${verbUpper} ${this.getPath()}`);

    try {
      const resBody = await getOpenApiSpec(this.opts.pluginRegistry);
      if (resBody === undefined) {
        res.status(400).json({
          message: "Bad request",
          error: resBody,
        });
      } else {
        res.status(200);
        res.send(resBody);
      }
    } catch (ex) {
      this.log.error(`${fnTag} failed to serve request`, ex);
      res.status(500).json({
        message: "Internal Server Error",
        error: ex?.stack || ex?.message,
      });
    }
  }

  public async getOpenApiSpec(): Promise<any> {
    return {};
  }
}
