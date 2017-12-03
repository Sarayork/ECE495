package main

import "gopkg.in/mgo.v2/bson"

type User struct{
	Uid         bson.ObjectId   `bson:"_id,omitempty" json:"uid"`
	FhirId		string    `json:"fhirid"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phoneNumber"`
	Roles       [] string `json:"roles"`
	Timestamp   string    `json:"timestamp"`
	CurrentIp   string    `json:"currentIp"`
}