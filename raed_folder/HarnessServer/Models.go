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

type Relation struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	SrcUid   string `bson:"srcuid,omitempty" json:"srcuid"`
	DestUid  string `bson:"destuid,omitempty" json:"destuid"`
	SrcRole  string `bson:"srcrole,omitempty" json:"srcrole"`
	DestRole string `bson:"destrole,omitempty" json:"destrole"`
	Status   string `bson:"status" json:"status"`
}