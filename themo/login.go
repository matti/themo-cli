package themo

type Login struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"TokenType"`
	ExpiresIn   int    `json:"ExpiresIn"`
}

func (t *Themo) Login(username string, password string) (*Login, error) {

	var login *Login

	resp, err := t.client.R().SetFormData(map[string]string{
		"grant_type": "password",
		"username":   username,
		"password":   password,
	}).SetResult(&login).
		Post("/token")

	if resp.IsSuccess() {
		t.client.SetAuthToken(login.AccessToken)
		return login, nil
	} else {
		return nil, err
	}
}
