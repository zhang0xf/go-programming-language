package chapter11_2

import (
	"fmt"
	"log"
	"net/smtp"
)

// 现在我们可以在测试中用伪邮件发送函数替代真实的邮件发送函数。它只是简单记录要通知的用户和邮件的内容。

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendEmail(%s) failed: %s", username, err)
	}
}

func CheckQuota2(username string) {
	fmt.Printf("call func CheckQuota2()\n")
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	fmt.Printf("msg : %s\n", msg)
	notifyUser(username, msg)
}
