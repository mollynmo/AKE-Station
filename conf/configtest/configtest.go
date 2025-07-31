package configtest

import "github.com/mollynmo/AKE-Station/conf"

func SetupConfig() func() {
	oldValues := *conf.Server
	return func() {
		conf.Server = &oldValues
	}
}
