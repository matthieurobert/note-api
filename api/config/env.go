package config

import (
	"os"
	"strconv"
)

type Env struct {
	ApiPort int
}

func (env *Env) Initenv() {
	env.ApiPort, _ = strconv.Atoi(os.Getenv("API_PORT"))
}
