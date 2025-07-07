import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";
import { urlService } from "../services/api";
import { URLShortenResponse } from "../types";

export const Dashboard: React.FC = () => {
  const [originalUrl, setOriginalUrl] = useState("");
  const [shortenedUrl, setShortenedUrl] = useState<URLShortenResponse | null>(
    null
  );
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const { user, logout } = useAuth();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setSuccess("");
    setIsLoading(true);

    try {
      const response = await urlService.shortenUrl({
        url: originalUrl,
      });

      setShortenedUrl(response);
      setSuccess("URL encurtada com sucesso!");
      setOriginalUrl("");
    } catch (err) {
      setError("Erro ao encurtar URL. Verifique se a URL é válida.");
    } finally {
      setIsLoading(false);
    }
  };

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
    setSuccess("URL copiada para a área de transferência!");
  };

  return (
    <div className="min-h-screen bg-gray-900">
      {/* Header */}
      <div className="bg-gray-800 shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-6">
            <div>
              <h1 className="text-3xl font-bold text-gray-100">
                URL Shortener
              </h1>
              <p className="text-gray-300">Bem-vindo, {user?.name}!</p>
            </div>
            <div className="flex items-center space-x-4">
              <Link
                to="/my-urls"
                className="bg-gray-700 text-gray-200 px-4 py-2 rounded-md hover:bg-gray-600 transition-colors"
              >
                Minhas URLs
              </Link>
              <button
                onClick={logout}
                className="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700 transition-colors"
              >
                Sair
              </button>
            </div>
          </div>
        </div>
      </div>

      <div className="max-w-4xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
        {/* Form para encurtar URL */}
        <div className="bg-gray-800 rounded-lg shadow-md p-6 mb-8">
          <h2 className="text-2xl font-bold text-gray-100 mb-6">
            Encurtar URL
          </h2>

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label
                htmlFor="originalUrl"
                className="block text-sm font-medium text-gray-300 mb-2"
              >
                URL para encurtar *
              </label>
              <input
                id="originalUrl"
                type="url"
                required
                className="w-full px-3 py-2 border border-gray-600 bg-gray-700 text-gray-100 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="https://exemplo.com/url-muito-longa"
                value={originalUrl}
                onChange={(e) => setOriginalUrl(e.target.value)}
              />
            </div>

            <button
              type="submit"
              disabled={isLoading}
              className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {isLoading ? "Encurtando..." : "Encurtar URL"}
            </button>
          </form>

          {error && (
            <div className="mt-4 p-4 bg-red-900 border border-red-700 rounded-md">
              <p className="text-red-200 text-sm">{error}</p>
            </div>
          )}

          {success && (
            <div className="mt-4 p-4 bg-green-900 border border-green-700 rounded-md">
              <p className="text-green-200 text-sm">{success}</p>
            </div>
          )}
        </div>

        {/* Resultado da URL encurtada */}
        {shortenedUrl && (
          <div className="bg-gray-800 rounded-lg shadow-md p-6">
            <h3 className="text-lg font-semibold text-gray-100 mb-4">
              URL Encurtada
            </h3>

            <div className="space-y-3">
              <div>
                <label className="block text-sm font-medium text-gray-300">
                  URL Original:
                </label>
                <p className="text-sm text-gray-400 break-all">
                  {shortenedUrl.originalUrl}
                </p>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-300">
                  URL Encurtada:
                </label>
                <div className="flex items-center space-x-2 mt-1">
                  <p className="text-sm text-blue-400 break-all flex-1">
                    {shortenedUrl.shortUrl}
                  </p>
                  <button
                    onClick={() => copyToClipboard(shortenedUrl.shortUrl)}
                    className="bg-gray-700 text-gray-200 px-3 py-1 rounded text-sm hover:bg-gray-600 transition-colors"
                  >
                    Copiar
                  </button>
                </div>
              </div>

              <div className="flex space-x-4 text-sm text-gray-500">
                <span>Cliques: {shortenedUrl.clickCount}</span>
                <span>
                  Criado em:{" "}
                  {new Date(shortenedUrl.createdAt).toLocaleDateString("pt-BR")}
                </span>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};
