package Utils

import (
	"fmt"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/pkg/messages"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

func SendEmail(payload authModel.RequestSignUpPayload, emailToken uuid.UUID) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "ricliz7@yandex.ru")
	m.SetHeader("To", payload.Email)
	m.SetHeader("Subject", "Добро пожаловать в HaveIdea!")
	url := fmt.Sprintf("http://92.246.141.141:8080/api/v1/auth/verify/%s", emailToken.String())

	body := messages.EmailConfirmMessage(payload, url)

	m.SetBody("text/plain", body)
	d := gomail.NewDialer("smtp.yandex.ru", 587, "ricliz7@yandex.ru", os.Getenv("SMTP_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
