package inn_auth

// Auth endpoints
const (
	LOGIN           = "/login"
	LOGOUT          = "/logout"
	REGISTER        = "/register"
	EMAIL_RESEND    = "/auth/email/resend"
	PASSWORD_CHANGE = "/auth/password/send"
	PASSWORD_FORGOT = "/auth/password/forgot"
	PASSWORD_RESET  = "/auth/password/reset"
	USERNAME_FORGOT = "/auth/username/forgot"
)
