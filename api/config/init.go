package config

var ENV = &Env{}

func Init() {
	ENV.Initenv()
}
