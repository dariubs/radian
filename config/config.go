package config

type CONFIG struct {
	Title       string
	Description string
	Logo        string
	File        FILE
	Server      SERVER
	User        USER
}

type FILE struct {
	Storage string
	Imgmime []string
	Default string
}

type SERVER struct {
	Port string
}

type USER struct {
	Accesskey  string
	Privatekey string
}
