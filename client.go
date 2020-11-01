package spotify

//Client for sending authenticated requests
type Client struct {
	Token string
}

//New functions generates a Client
func New(token string) Client {
	return Client{
		Token: token,
	}
}