package message

type ResponseSignin struct {
	WorkerID uint   `json:"workerId"`
	Email    string `json:"email"`
}
