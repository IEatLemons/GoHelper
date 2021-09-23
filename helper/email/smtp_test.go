package Email

import (
	"fmt"
	"testing"
	"time"

	EmailModules "github.com/IEatLemons/GoHelper/helper/email/modules"
	"github.com/IEatLemons/GoHelper/language"
)

func InitDb() {
	Host := "smtp.163.com:465"
	User := "xxx@163.com"
	Pwd := ""
	Reply := "xxx@163.com"

	Init(Host, User, Pwd, Reply)
}

func TestSmtp_Send(t *testing.T) {
	InitDb()
	Data := &EmailModules.Registered{
		User:       "syj88668@gmail.com",
		Platform:   "sofa",
		Expiration: time.Now(),
		Language:   language.ZhCnLan,
	}
	err := SMTP.Send(Data)
	fmt.Println("err", err)
}
