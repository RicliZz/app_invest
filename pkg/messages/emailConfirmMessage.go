package messages

import (
	"fmt"
	"github.com/RicliZz/app_invest/internal/models/authModel"
)

func EmailConfirmMessage(payload authModel.RequestSignUpPayload, url string) string {
	body := fmt.Sprintf(
		`Здравствуйте %s %s!
				
				Спасибо за регистрацию в нашем сервисе. Чтобы завершить процесс регистрации, подтвердите ваш Email, перейдя по ссылке ниже:
				%s
				Оно действует ровно 24 часа!

				Если вы не регистрировались в HaveIdea просто проигнорируйте это письмо.

																	С уважением, команда HaveIdea!`, payload.FirstName, payload.MiddleName, url)
	return body
}
