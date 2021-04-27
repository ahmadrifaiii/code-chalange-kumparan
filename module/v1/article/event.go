package article

import (
	"fmt"
	"kumparan/config/nats"

	"github.com/labstack/gommon/log"
)

// handle event create article
func (m *Module) EventCreateArticle() {
	nc, err := nats.NewConnection()
	if err != nil {
		log.Error(err)
		return
	}

	resp, err := nats.Subscription(nc, "create.article")
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println(resp)
}
