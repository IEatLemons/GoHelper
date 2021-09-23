package Email

import (
	"fmt"
	"testing"

	EmailModules "github.com/IEatLemons/GoHelper/helper/email/modules"
)

func InitDb() {
	Host := "smtp.163.com:25"
	User := "@163.com"
	Pwd := ""
	Reply := "xxx@163.com"

	Init(Host, User, Pwd, Reply)
}

func TestSmtp_Send(t *testing.T) {
	InitDb()
	Data := &EmailModules.Impact{
		User:    "syj88668@gmail.com",
		Subject: "Test Sender",
		Name:    "Test Sender",
		Number:  "0",
	}
	err := SMTP.Send(Data)
	fmt.Println("err", err)
}
