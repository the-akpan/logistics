import axios from "axios";

const BASE_URL = "https://brighttimelogistics.xyz/api/";

export const URLS = {
  auth: {
    login: "auth/signin/",
    logout: "auth/logout/",
    reset_password: "auth/reset-password/",
  },
};

export const Methods = {
  POST: "POST",
  GET: "GET",
};

const config = {
  baseURL: BASE_URL,
  withCredentials: true,
};

const api = axios.create(config);

export async function makeRequest(endpoint, method, params) {
  let result = null;
  let request;

  switch (method) {
    case Methods.GET:
      request = api.get;
      break;
    case Methods.POST:
      request = api.post;
      break;
    default:
      throw new Error("Unknown method " + method);
  }

  try {
    const response = await request(endpoint, params);
    result = response.data;
  } catch (error) {
    result = error?.response?.data;
    result.error = true;
  }

  return result;
}
