package repository

import (
	"context"
	"encoding/json"
	"log"
	"main/config"
	"main/model"
	"main/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var (
	collection = config.ConnectDB()
	content    = utils.ContentType
	appjson    = utils.ApplicationJson
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, appjson)

	var users []model.User

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		utils.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var user model.User
		// & character returns the memory address of the following variable.
		err := cur.Decode(&user) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users) // encode similar to serialize process.
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, appjson)

	var user model.User
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, appjson)

	var user model.User

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&user)

	// insert our user model.
	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, appjson)

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var user model.User

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&user)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"nickname", user.NickName},
			{"fullname", user.FullName},
			{"freelance", user.Freelance},
			{"email", user.Email},
			{"address1", user.Address1},
			{"address2", user.Address2},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	user.ID = id

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, appjson)

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
