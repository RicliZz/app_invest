package Utils

import (
	"fmt"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
	"log"
)

func SendEmail(payload authModel.RequestSignUpPayload, emailToken uuid.UUID) {
	m := gomail.NewMessage()
	log.Println(payload.Email)
	m.SetHeader("From", "ricliz7@yandex.ru")
	m.SetHeader("To", payload.Email)
	m.SetHeader("Subject", "Добро пожаловать в HaveIdea!")
	url := fmt.Sprintf("http://localhost:8080/api/v1/auth/verify/%s", emailToken.String())

	body := fmt.Sprintf(
		`Здравствуйте %s %s!
				
				Спасибо за регистрацию в нашем сервисе. Чтобы завершить процесс регистрации, подтвердите ваш Email, перейдя по ссылке ниже:
				%s

				Если вы не регистрировались в HaveIdea просто проигнорируйте это письмо

												С уважением, команда HaveIdea!`, payload.FirstName, payload.MiddleName, url)

	m.SetBody("text/plain", body)
	d := gomail.NewDialer("smtp.yandex.ru", 587, "ricliz7@yandex.ru", "lwztvnjvrviwqqsg")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
