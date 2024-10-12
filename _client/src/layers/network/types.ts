// entity status can only be idle or loading
// this is not request status
// and request status (fail or success) will be stored separately
type NetworkEntityStatus = "idle" | "loading";

export type { NetworkEntityStatus };
