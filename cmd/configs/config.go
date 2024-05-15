package configs

// =================== Define Config globally =================== //
var Config = conf{
	Session: confSession{
		Token: "You bot toke here",
		Id:    "Your bot id here",
	},

	Handler: confHandler{
		Prefix: "!!",
	},

	Users: confUsers{
		Developers: [1]string{"144694124925288448"},
	},

	Logs: confLogs{
		Enable:     true,
		Guild:      "",
		GlitchLogs: "",
		LeftLogs:   "",
	},

	Messages: confMessages{
		DeveloperMessage:    "This is developer only commands",
		CooldownMessage:     "Your are on cooldown. Please wait",
		OwnerOnlyMessage:    "Only owners can use this command",
		NoPermissionMessage: "You do not have permission to use this command",
	},

	Emojies: confEmojis{
		Success:   "",
		Attention: "",
	},
}

////////////////////////////////////////////////////////////////////

type conf struct {
	Session  confSession
	Handler  confHandler
	Users    confUsers
	Logs     confLogs
	Messages confMessages
	Emojies  confEmojis
}

type confSession struct {
	Token string
	Id    string
}

type confHandler struct {
	Prefix string
}

type confUsers struct {
	Developers [1]string
}

type confLogs struct {
	Enable     bool
	Guild      string
	GlitchLogs string
	LeftLogs   string
}

type confMessages struct {
	DeveloperMessage    string
	CooldownMessage     string
	OwnerOnlyMessage    string
	NoPermissionMessage string
}

type confEmojis struct {
	Success   string
	Attention string
}
