package models

type EnviromentModel struct {
	Line   line
	Email  email
	Server server
}

type line struct {
	Token string
}

type email struct {
	User     string
	Password string
	Smtp     string
	Host     string
	Sender   string
	To       []string
	Auth     bool
}

type server struct {
	Mode string
	Port string
}
