export interface User {
  id: string;
  email: string;
  name: string;
}

export interface AuthContextType {
  user: User | null;
  login: (email: string, password: string) => Promise<void>;
  register: (name: string, email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  isLoading: boolean;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface AuthResponse {
  token: string;
  user: User;
}

export interface URLShortenRequest {
  url: string;
}

export interface URLShortenResponse {
  shortUrl: string;
  originalUrl: string;
  clickCount: number;
  createdAt: string;
}

export interface URLStats {
  id: string;
  originalUrl: string;
  shortUrl: string;
  clickCount: number;
  createdAt: string;
}

export interface URLListItem {
  ID: string;
  UrlOriginal: string;
  UrlShortened: string;
  Slug: string;
  CreatedAt: string;
}
