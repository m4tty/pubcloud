package pubsDomain

import (
	"github.com/m4tty/pubcloud/data/pubs"
	"github.com/m4tty/pubcloud/web/resources"
)

type PubsMgr struct {
	pubsDataMgr.PubsDataManager
}

func NewPubsMgr(bdm pubsDataMgr.PubsDataManager) *PubsMgr {
	return &PubsMgr{bdm}
}

func (dm PubsMgr) GetPubs() (pubs []*resources.PubsResource, err error) {
	// dPubs, err := dm.PubsDataManager.GetPubs()
	// if err != nil {
	// 	return nil, err
	// }

	// pubs = make([]*resources.PubsResource, len(dPubs))
	// for j, pubs := range dPubs {
	// 	var pubsResource *resources.PubsResource = new(resources.PubsResource)
	// 	mapDataToResource(pubs, pubsResource)
	// 	pubs[j] = pubsResource
	// }
	return nil, nil
}

// mapper...
func mapResourceToData(pubsResource *resources.PubsResource, pubsData *pubsDataMgr.Pub) {
	pubsData.Name = pubsResource.Name

}

func mapDataToResource(pubsData *pubsDataMgr.Pub, pubsResource *resources.PubsResource) {

	pubsResource.Name = pubsData.Name

}
