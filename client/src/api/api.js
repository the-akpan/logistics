import axios from "axios";

const BASE_URL = "https://brighttimelogistics.xyz/";

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
};

function getAuthConfig() {
  const token = localStorage.getItem("token");

  return {
    ...config,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
}

const api = axios.create(config);
const authAPI = axios.create(getAuthConfig());

export async function makeRequest(endpoint, method, params, isAuth = false) {
  let result = null;
  let request = isAuth ? api : authAPI;

  switch (method) {
    case Methods.GET:
      request = request.get;
      break;
    case Methods.POST:
      request = request.post;
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
