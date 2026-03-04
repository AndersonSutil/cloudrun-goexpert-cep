# Cloud Run Weather API - Go

Este sistema recebe um CEP, identifica a cidade correspondente e retorna o clima atual (Celsius, Fahrenheit e Kelvin).

## URL de Acesso (Google Cloud Run)
A aplicação está implantada no Google Cloud Run e pode ser acessada através da seguinte URL:
**[https://cloudrun-weather-api-qax7333zha-uc.a.run.app/weather?zipcode=01001000](https://cloudrun-weather-api-qax7333zha-uc.a.run.app/weather?zipcode=01001000)**

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
