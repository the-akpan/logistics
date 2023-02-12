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

  switch (method) {
    case Methods.GET:
      api.get;
      break;
    case Methods.POST:
      api.post;
      break;
    default:
      throw new Error("Unknown method " + method);
  }

  try {
    const response = await api(endpoint, params);
    result = response.data;
  } catch (error) {
    result = error?.response?.data;
    result.error = true;
  }

  return result;
}
