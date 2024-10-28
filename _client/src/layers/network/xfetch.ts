/* eslint-disable @typescript-eslint/no-explicit-any */
import {
  EndpointParam,
  XEndpoint,
  XHeaders,
  XRequest,
} from "@/src/layers/network/types.ts";
import { endpoints } from "@/src/layers/network/endpoints.ts";

const xfetch = async <Params extends EndpointParam[], Body = any>(
  endpoint: XEndpoint<Params>,
  request: XRequest<Body>,
) => {
  const headers: XHeaders = {
    "Content-Type": "application/json",
  };

  const endpointUrl = endpoints.getUrl(endpoint.code);

  if (request.useToken) {
    // todo: get token
    const token: string | null = null;
    if (token) {
      headers["Authorization"] = token;
    }
  }

  return fetch(endpointUrl, {
    method: request.method,
    headers: headers,
    ...(request.body && {
      body: JSON.stringify(request.body),
    }),
  });
};

export { xfetch };
