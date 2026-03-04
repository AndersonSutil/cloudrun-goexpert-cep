# Cloud Run Weather API - Go

Este sistema recebe um CEP, identifica a cidade correspondente e retorna o clima atual (Celsius, Fahrenheit e Kelvin).

## URL de Acesso (Google Cloud Run)
A aplicação está implantada no Google Cloud Run e pode ser acessada através da seguinte URL:
**[https://weather-api-524282631878.us-central1.run.app/weather?zipcode=84350000](https://cloudrun-weather-api-qax7333zha-uc.a.run.app/weather?zipcode=01001000)**

*(Substitua `01001000` pelo CEP desejado)*

## Como rodar localmente

### Pré-requisitos
- Go 1.24 instalado (opcional, se rodar via Docker)
- Docker instalado
- API Key da [WeatherAPI](https://www.weatherapi.com/)

### Configuração
Defina a variável de ambiente `WEATHER_API_KEY`:
```bash
export WEATHER_API_KEY=sua_api_key_aqui
```

### Rodando com Docker
1. Construa a imagem:
   ```bash
   docker build -t cloudrun-weather-api .
   ```
2. Execute o container:
   ```bash
   docker run -p 8080:8080 -e WEATHER_API_KEY=sua_api_key_aqui cloudrun-weather-api
   ```
3. Acesse em: `http://localhost:8080/weather?zipcode=01001000`

### Rodando os Testes
Para rodar os testes automatizados, utilize o comando:
```bash
go test ./internal/...
```

### Deploy no Google Cloud Run

Para realizar o deploy da aplicação no Google Cloud Run, siga os passos abaixo:

1. **Autenticação no Google Cloud:**
   ```bash
   gcloud auth login
   gcloud auth configure-docker
   ```

2. **Configuração do Projeto:**
   ```bash
   gcloud config set project [NOME_DO_PROJETO]
   ```

3. **Habilitar APIs Necessárias:**
   ```bash
   gcloud services enable run.googleapis.com \
       artifactregistry.googleapis.com \
       cloudbuild.googleapis.com
   ```

4. **Criar Repositório no Artifact Registry:**
   ```bash
   gcloud artifacts repositories create cloudrun-repo \
       --repository-format=docker \
       --location=us-central1
   ```

5. **Build e Push da Imagem:**
   ```bash
   docker build -t us-central1-docker.pkg.dev/[NOME_DO_PROJETO]/cloudrun-repo/weather-api .
   docker push us-central1-docker.pkg.dev/[NOME_DO_PROJETO]/cloudrun-repo/weather-api
   ```

6. **Deploy no Cloud Run:**
   ```bash
   gcloud run deploy weather-api \
       --image us-central1-docker.pkg.dev/[NOME_DO_PROJETO]/cloudrun-repo/weather-api \
       --platform managed \
       --region us-central1 \
       --allow-unauthenticated \
       --set-env-vars WEATHER_API_KEY=[SUA_WEATHER_API_KEY]
   ```

Após o deploy, a URL de acesso será exibida no terminal.

## Requisitos Funcionais
- **Entrada:** CEP válido de 8 dígitos.
- **Saída (Sucesso - 200):**
  ```json
  { "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.65 }
  ```
- **Falha (Formato Inválido - 422):** `invalid zipcode`
- **Falha (CEP Não Encontrado - 404):** `can not find zipcode`

## Estrutura do Projeto
- `internal/model`: Estruturas de dados (JSON responses).
- `internal/service`: Lógica de negócio, conversões e clientes de API.
- `internal/handler`: Handlers HTTP.
- `main.go`: Ponto de entrada da aplicação.
