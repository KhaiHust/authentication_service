package main

import (
	"github.com/KhaiHust/authen_service/public/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.All()).Run()
}
