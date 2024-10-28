import {
  ApiService,
  ApiVersion,
  EndpointData,
  EndpointList,
  EndpointParam,
} from "@/src/layers/network/types.ts";
import { env } from "@/src/env.ts";

class EndpointsController {
  private _endpointList: EndpointList = new Map<string, EndpointData>();

  get endpointList() {
    return this._endpointList;
  }

  public add(endpoint: Endpoint) {
    if (endpoint.endpoint?.code) {
      if (this._endpointList.has(endpoint.endpoint.code)) {
        throw new Error("endpoint_controller_endpoint_already_exists");
      }

      this._endpointList.set(endpoint.endpoint?.code, endpoint.endpoint);
    }

    return this;
  }

  public getUrl(code: string, params?: EndpointParam[]) {
    const endpoint = this._endpointList.has(code)
      ? this._endpointList.get(code)
      : undefined;

    if (endpoint) {
      const base = env.baseUrl ?? "<base>";
      const api = endpoint.api ?? "<api>";
      const service = endpoint.service ?? "<service>";
      let path = endpoint.path ?? "<empty_path>";

      if (params && params.length > 0) {
        params.forEach((param) => {
          path = path.replace(
            `%${param.key}%`,
            param.value ? param.value : "null",
          );
        });
      }

      return [base, "api", api, service, path].filter((e) => e).join("/");
    } else {
      throw new Error("endpoint_not_found");
    }
  }
}

class Endpoint {
  private _endpoint: EndpointData | null = null;

  get endpoint() {
    return this._endpoint;
  }

  constructor(code: string) {
    this.update({ code });
  }

  public name(name: string) {
    this.update({ name });
    return this;
  }

  public api(api: ApiVersion) {
    this.update({ api });
    return this;
  }

  public service(service: ApiService) {
    this.update({ service });
    return this;
  }

  public path(path: string): Endpoint {
    this.update({ path });
    return this;
  }

  private update(data: Partial<EndpointData>) {
    this._endpoint = Object.assign({}, this._endpoint, data);
  }
}

export { Endpoint, EndpointsController };
