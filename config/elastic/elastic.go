package elastic

import (
	"kumparan/config/env"
	"sync"

	"github.com/olivere/elastic/v7"
)

var (
	lockElastic sync.Locker
)

func NewElastic() (e *elastic.Client, err error) {

	lockElastic.Lock()
	defer lockElastic.Unlock()

	if lockElastic == nil {
		e, err = elastic.NewClient(
			elastic.SetURL(env.Conf.ElasticHost),
			elastic.SetSniff(false),
		)
	}

	return
}
