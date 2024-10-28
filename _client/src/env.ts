type Env = {
  baseUrl?: string;
};

const env: Env = {
  baseUrl: import.meta.env.VITE_BASE_URL,
};

export { env };
