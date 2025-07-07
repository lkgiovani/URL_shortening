import axios from "axios";
import {
  AuthResponse,
  LoginRequest,
  RegisterRequest,
  URLShortenRequest,
  URLShortenResponse,
  URLListItem,
  User,
} from "../types";

const API_BASE_URL = "http://localhost:8181";

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true, // Incluir cookies nas requisições
});

export const authService = {
  login: async (data: LoginRequest): Promise<AuthResponse> => {
    const response = await api.post("/auth/login", data);
    return response.data;
  },

  register: async (data: RegisterRequest): Promise<AuthResponse> => {
    const response = await api.post("/auth/register", data);
    return response.data;
  },

  logout: async (): Promise<void> => {
    await api.post("/auth/logout");
  },

  me: async (): Promise<{ user: User }> => {
    const response = await api.get("/auth/me");
    return response.data;
  },
};

export const urlService = {
  shortenUrl: async (data: URLShortenRequest): Promise<URLShortenResponse> => {
    const response = await api.post("/register", data);
    return response.data;
  },

  getUserUrls: async (): Promise<URLListItem[]> => {
    const response = await api.get("/urls");
    return response.data.urls || [];
  },

  getStats: async (): Promise<URLShortenResponse[]> => {
    const response = await api.get("/stats");
    return response.data;
  },
};

export default api;
