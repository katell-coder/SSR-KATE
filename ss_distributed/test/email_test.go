package test

import (
	"crypto/tls"
	"math/rand"
	"net/smtp"
	"strconv"
	"testing"
	"time"

	"github.com/jordan-wright/email"
)

var Code = make(map[string]int) //key为Email,value为随机码
func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <fucewei98@163.com>"
	e.To = []string{"1457310354@qq.com"}
	e.Subject = "验证码发送测试..."
	rand.Seed(time.Now().Unix())
	// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
	num := rand.Intn(10000)

	e.HTML = []byte("<h1>您的验证码: <b>" + strconv.Itoa(num) + "</b></h1>")

	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("",
		"fucewei98@163.com", "XQYORQXGAFNKCDLX", "smtp.163.com"), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.163.com",
	})
	if err != nil {
		t.Fatal(err)
	}
}
