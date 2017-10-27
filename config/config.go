package config

type CONFIG struct {
	Title       string
	Description string
	Logo        string
	File        FILE
	Server      SERVER
}

type FILE struct {
	Storage string
	Imgmime []string
}

type SERVER struct {
	Port string
}
