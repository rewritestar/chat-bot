package core_values

var (
	EnvDBUser          = "DB_USER"
	EnvDBPW            = "DB_PW"
	EnvDBHost          = "DB_HOST"
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

var (
	ChatHistoryLimit    = 20
	SummaryHistoryLimit = 30
)

var (
	SummaryPromptFrame = `
	You are updating the long-term memory for a chat room between a user and an AI assistant. 
	Your goal is to produce a concise MEMORY that the AI should carry forward for all future conversations in this room. 
	You will be given: 
	1) The existing memory summary for this room (may be empty). 
	2) The most recent conversation messages. Using BOTH, produce an UPDATED memory summary.
	3) Keep the memory concise but allow enough detail to preserve long-term important information.
	4) Do NOT lose previous memory summary if important.
	What to INCLUDE in the memory: 
	- User’s preferences, personality, and communication style
	- User’s important personal information(such as name)
	- User’s goals, ongoing projects, and intentions 
	- Important facts the user shared about themselves 
	- Important decisions or plans or feelings 
	- Context that would help the AI respond better in the future What to 
	EXCLUDE from the memory: 
	- One-time questions 
	- Small talk 
	- Detailed step-by-step problem solving logs 
	- Exact dialogue history Rules: 
	- The memory must be written in third person (e.g., "The user prefers...", not "You prefer...") 
	- Do not mention specific timestamps or message counts 
	- Do not refer to this task or the conversation itself 
	- Write as if this memory will be injected into the system prompt of future chats Return ONLY the updated memory summary. 
	------ Existing memory: 
	%s 
	------ Recent conversation: 
	%s
	`
)

func UintToPointer(input uint) *uint {
	return &input
}
