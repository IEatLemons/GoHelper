package EmailModules

type Impact struct {
	User    string `json:"user"`
	Subject string `json:"subject"`
	Name    string `json:"name"`
	Number  string `json:"number"`
}

func (M *Impact) GetTo() string {
	return M.User
}

func (M *Impact) GetSubject() string {
	return M.Subject
}

func (M *Impact) GetBody() string {
	body := "name : " + M.Name + "; number : " + M.Number
	return body
}

func (M *Impact) GetMailType() string {
	return "text"
}
