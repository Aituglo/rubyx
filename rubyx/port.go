package rubyx

import (
	"fmt"

	"github.com/aituglo/rubyx/pkg"
)

func AddPort(config pkg.Config, port int, ip_id int) {
	var data = []byte(fmt.Sprintf(`{
		"ip_id": %d,
		"port": %d
	}`, ip_id, port))

	pkg.Post(config, "port", data)
}
