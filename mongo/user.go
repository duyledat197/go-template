package mongo

import (
	"github.com/stamp-server/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userRepository struct {
	session        *mgo.Session
	db             string
	collectionName string
}

func (r *userRepository) Create(user *models.User) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collectionName)
	err := c.Insert(user)
	return err
}

func (r *userRepository) FindByUserName(userName string) (*models.User, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collectionName)
	var result models.User
	if err := c.Find(bson.M{"userName": userName}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, models.ErrUnknowUser
		}
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) FindByUserID(userID string) (*models.User, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collectionName)
	var result models.User
	if err := c.Find(bson.M{"_id": userID}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, models.ErrUnknowUser
		}
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) FindAll() ([]*models.User, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collectionName)
	var result []*models.User
	if err := c.Find(bson.M{}).All(&result); err != nil {
		return []*models.User{}, err
	}
	return result, nil
}

func (r *userRepository) Update(userID string, user *models.User) error {
	return nil
}

func (r *userRepository) Delete(userID string) error {
	return nil
}

// NewUserRepository ...
func NewUserRepository(db string, session *mgo.Session, collectionName string) models.UserRepository {
	return &userRepository{
		db:             db,
		session:        session,
		collectionName: collectionName,
	}
}
