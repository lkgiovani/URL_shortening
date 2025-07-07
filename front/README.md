# URL Shortening Frontend

Frontend em React + TypeScript + Tailwind CSS para o sistema de encurtamento de URLs.

## 🚀 Tecnologias

- **React 18** - Biblioteca principal
- **TypeScript** - Tipagem estática
- **Tailwind CSS** - Estilização
- **React Router DOM** - Roteamento
- **Axios** - Requisições HTTP
- **Vite** - Build tool

## 📦 Instalação

```bash
# Instalar dependências
npm install

# Executar em modo desenvolvimento
npm run dev

# Build para produção
npm run build
```

## 🔧 Configuração

O projeto está configurado para se conectar com a API backend em `http://localhost:8080`.

Para alterar esta configuração, edite o arquivo `src/services/api.ts`:

```typescript
const API_BASE_URL = "http://localhost:8080";
```

## 📁 Estrutura do Projeto

```
src/
├── components/          # Componentes reutilizáveis
├── contexts/            # Contextos React (AuthContext)
├── pages/               # Páginas da aplicação
│   ├── Login.tsx        # Página de login
│   ├── Register.tsx     # Página de cadastro
│   └── Dashboard.tsx    # Página principal
├── services/            # Serviços e API
├── types/               # Definições de tipos TypeScript
├── App.tsx              # Componente principal
├── main.tsx             # Ponto de entrada
└── index.css            # Estilos globais
```

## 🎨 Funcionalidades

### 🔐 Autenticação

- Login de usuário
- Cadastro de usuário
- Logout
- Proteção de rotas

### 🔗 Encurtamento de URLs

- Encurtar URLs longas
- Alias personalizado (opcional)
- Copiar URL encurtada
- Exibir estatísticas (cliques, data de criação)

## 📱 Páginas

### Login (`/login`)

- Formulário de login com email e senha
- Validação de campos
- Redirecionamento automático após login

### Cadastro (`/register`)

- Formulário de registro
- Validação de senha (mínimo 6 caracteres)
- Confirmação de senha

### Dashboard (`/dashboard`)

- Formulário para encurtar URLs
- Exibição da URL encurtada
- Botão para copiar URL
- Estatísticas da URL

## 🎯 Como Usar

1. **Cadastre-se** ou faça **login**
2. Na página principal, **insira a URL** que deseja encurtar
3. Opcionalmente, adicione um **alias personalizado**
4. Clique em **"Encurtar URL"**
5. **Copie** a URL encurtada gerada
6. Compartilhe a URL encurtada!

## 🔗 Integração com Backend

O frontend se integra com os seguintes endpoints:

- `POST /auth/login` - Login
- `POST /auth/register` - Cadastro
- `POST /register` - Encurtar URL
- `GET /:urlShortened` - Acessar URL encurtada

## 🎨 Estilização

O projeto usa **Tailwind CSS** para estilização com:

- Design responsivo
- Tema consistente de cores
- Componentes reutilizáveis
- Feedback visual (loading states, erros, sucessos)

## 🚀 Deploy

```bash
# Build para produção
npm run build

# Preview do build
npm run preview
```

Os arquivos de produção ficam na pasta `dist/`.

## 📝 Notas

- O projeto está configurado para funcionar com o backend Go
- As rotas são protegidas com autenticação JWT
- O estado da aplicação é gerenciado com React Context
- Todas as requisições incluem o token de autenticação automaticamente
