import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";
import { urlService } from "../services/api";
import { URLListItem } from "../types";

export const MyUrls: React.FC = () => {
  const [urls, setUrls] = useState<URLListItem[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const { user, logout } = useAuth();

  useEffect(() => {
    loadUrls();
  }, []);

  const loadUrls = async () => {
    try {
      setIsLoading(true);
      const userUrls = await urlService.getUserUrls();
      setUrls(userUrls);
    } catch (err) {
      setError("Erro ao carregar URLs");
      console.error("Error loading URLs:", err);
    } finally {
      setIsLoading(false);
    }
  };

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
    setSuccess("URL copiada para a área de transferência!");
    setTimeout(() => setSuccess(""), 3000);
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("pt-BR", {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  return (
    <div className="min-h-screen bg-gray-900">
      {/* Header */}
      <div className="bg-gray-800 shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-6">
            <div>
              <h1 className="text-3xl font-bold text-gray-100">Minhas URLs</h1>
              <p className="text-gray-300">Gerencie suas URLs encurtadas</p>
            </div>
            <div className="flex items-center space-x-4">
              <Link
                to="/dashboard"
                className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors"
              >
                Nova URL
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

      <div className="max-w-7xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
        {error && (
          <div className="mb-4 p-4 bg-red-900 border border-red-700 rounded-md">
            <p className="text-red-200 text-sm">{error}</p>
          </div>
        )}

        {success && (
          <div className="mb-4 p-4 bg-green-900 border border-green-700 rounded-md">
            <p className="text-green-200 text-sm">{success}</p>
          </div>
        )}

        {isLoading ? (
          <div className="flex justify-center items-center py-12">
            <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-600"></div>
          </div>
        ) : urls.length === 0 ? (
          <div className="text-center py-12">
            <h3 className="text-xl font-medium text-gray-400 mb-4">
              Nenhuma URL encontrada
            </h3>
            <p className="text-gray-500 mb-6">
              Você ainda não criou nenhuma URL encurtada.
            </p>
            <Link
              to="/dashboard"
              className="bg-blue-600 text-white px-6 py-3 rounded-md hover:bg-blue-700 transition-colors"
            >
              Criar primeira URL
            </Link>
          </div>
        ) : (
          <div className="bg-gray-800 rounded-lg shadow-md overflow-hidden">
            <div className="px-6 py-4 border-b border-gray-700">
              <h2 className="text-xl font-semibold text-gray-100">
                Suas URLs ({urls.length})
              </h2>
            </div>

            <div className="divide-y divide-gray-700">
              {urls.map((url) => (
                <div
                  key={url.ID}
                  className="p-6 hover:bg-gray-700 transition-colors"
                >
                  <div className="flex flex-col md:flex-row md:items-center md:justify-between space-y-4 md:space-y-0">
                    <div className="flex-1 min-w-0">
                      <div className="mb-2">
                        <label className="block text-sm font-medium text-gray-300">
                          URL Original:
                        </label>
                        <p className="text-sm text-gray-400 break-all">
                          {url.UrlOriginal}
                        </p>
                      </div>

                      <div className="mb-2">
                        <label className="block text-sm font-medium text-gray-300">
                          URL Encurtada:
                        </label>
                        <div className="flex items-center space-x-2">
                          <p className="text-sm text-blue-400 break-all flex-1">
                            {url.UrlShortened}
                          </p>
                          <button
                            onClick={() => copyToClipboard(url.UrlShortened)}
                            className="bg-gray-700 text-gray-200 px-3 py-1 rounded text-sm hover:bg-gray-600 transition-colors"
                          >
                            Copiar
                          </button>
                        </div>
                      </div>

                      <div className="text-xs text-gray-500">
                        Criado em: {formatDate(url.CreatedAt)}
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
};
