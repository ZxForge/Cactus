package emailresponce

import email_shema "cactus/internal/schema/email-shema"

type GetEmailsResponse struct {
	Emails []email_shema.Email `json:"emails"`
}
