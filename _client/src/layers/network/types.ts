// endpoints
type ApiVersion = "v1" | "v2";
type ApiService = "users" | "todos";
type EndpointData = {
  name: string;
  code: string;
  api: ApiVersion;
  service: ApiService;
  path: string;
};
type EndpointList = Map<string, EndpointData>;
type EndpointParam = {
  key: string;
  value: string | null;
};

// utils
type KeyValueDict = { key: string; value: string | number };

// network

// entity status can only be idle or loading
// this is not request status
// and request status (fail or success) will be stored separately
type NetworkEntityStatus = "idle" | "loading";
type NetworkRequestStatus = "fail" | "success";

type XHeaders = {
  "Content-Type": string;
  Authorization?: string;
};

type XHTTPMethod = "get" | "post" | "put" | "delete";

type XEndpoint<Params extends EndpointParam[]> = {
  code: string;
  params?: Params;
  query?: KeyValueDict[]; // key value dict
};

type XRequest<Body> = {
  method: XHTTPMethod;
  useToken?: true;
  body?: Body;
};

type XResponse<DataType = never> = {
  status: number;
  error: null | boolean;
  message: null | string;
  data: DataType;
};

export type { ApiVersion, ApiService, EndpointData, EndpointList };

export type {
  NetworkEntityStatus,
  NetworkRequestStatus,
  KeyValueDict,
  EndpointParam,
  XHeaders,
  XHTTPMethod,
  XEndpoint,
  XRequest,
  XResponse,
};
