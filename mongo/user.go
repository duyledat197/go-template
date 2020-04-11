package mongo

import (
	"github.com/duyledat197/go-template/models/domain"
	mongoI "github.com/duyledat197/go-template/mongo/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type userRepository struct {
	coll mongoI.CtxCollection
}

func (r *userRepository) Create(user *domain.User) error {
	_, err := r.coll.InsertOne(user)
	return err
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var result domain.User
	err := r.coll.FindOne(bson.M{"email": email}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUnknowUser
		}
	}
	return &result, nil
}

func (r *userRepository) FindByUserID(userID string) (*domain.User, error) {
	var result *domain.User
	if err := r.coll.FindOne(bson.M{"_id": userID}).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) FindAll() ([]*domain.User, error) {
	var result []*domain.User
	cur, err := r.coll.Find(bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.Decode(&result)
	return result, nil
}

func (r *userRepository) Update(userID string, user *domain.User) error {
	return nil
}

func (r *userRepository) Delete(userID string) error {
	return nil
}

// NewUserRepository ...
func NewUserRepository(coll mongoI.CtxCollection) domain.UserRepository {
	return &userRepository{
		coll: coll,
	}
}
