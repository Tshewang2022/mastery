package mailer

import "embed"

// need to have the retry mechanism if the mailing failed
const (
	FromName            = "social"
	maxRetries          = 3
	UserWelcomeTemplate = "user_invitation.tmpl"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile string, username, email string, data any, isSandbox bool) error
}
