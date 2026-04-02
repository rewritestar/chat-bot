package core_values

var (
	EnvJwtSecret       = "JWT_SECRET"
	EnvVapidPublicKey  = "VAPID_PUBLIC_KEY"
	EnvVapidPrivateKey = "VAPID_PRIVATE_KEY"
	EnvClientUrl       = "CLIENT_URL"
	EnvAdminEmail      = "ADMIN_EMAIL"
	WorkerIDKey        = "workerId"
)

var (
	EnvOllamaApiKey = "OLLAMA_API_KEY"
	OllamaApiURL    = "https://ollama.com/api"
	OllamaModel     = "gpt-oss:20b"

	OllamaRoleSystem    = "system"
	OllamaRoleUser      = "user"
	OllamaRoleAssistant = "assistant"
)
