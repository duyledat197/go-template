package mongo

import (
	"github.com/stamp-server/models"
	"gopkg.in/mgo.v2"
)

type walletRepository struct {
	session        *mgo.Session
	db             string
	collectionName string
}

func (r *walletRepository) Create(wallet *models.Wallet) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collectionName)
	err := c.Insert(wallet)
	return err
}

// NewWalletRepository ...
func NewWalletRepository(db string, session *mgo.Session, collectionName string) models.WalletRepository {
	return &walletRepository{
		db:             db,
		session:        session,
		collectionName: collectionName,
	}
}
