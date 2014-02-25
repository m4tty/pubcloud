package pubsDataMgr

type PubsDataManager interface {
	GetPubs() (results []*Pub, err error)
}
