package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"log"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"goji.io/pat"
)

const DBNAME = "harness"

/**
 * @api {post} /user Subscribe a new user
 * @apiName PostUser
 * @apiGroup User
 * @apiVersion 0.1.0
 *
 * @apiDescription Generate unique Users-ID and add new user account in the HArNESS data base
 *
 * @apiParamExample {json} The body of the request is in json format. Example:
 * {
 * 	"uid": "",
 *  "fhirid": "191021",
 * 	"name": "Jane",
 * 	"password": "ac65c0f32e6a9083b6a3f8dc8badd576",
 * 	"email": "name@example.com",
 * 	"address": "Geneva",
 * 	"phoneNumber": "004125466321",
 * 	"roles": ["patient", "doctor"],
 * 	"timestamp": "3600",
 * 	"currentip": ""
 * }
 *
 * @apiSuccess {string} uid Users-ID
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 201 (Created)
 *     {
 *       "uid": "12461"
 *     }
 *
 * @apiUse UserNameAlreadyExist
 */
func subscribe(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		var user User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}
		user.Uid = bson.NewObjectId()

		c := session.DB(DBNAME).C("users")

		err = c.Insert(user)
		if err != nil {
			if mgo.IsDup(err) {
				ErrorWithJSON(w, "FhirIdAlreadyExist", http.StatusBadRequest)
				return
			}

			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed insert user: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", r.URL.Path + "/" + user.Uid.Hex())
		w.WriteHeader(http.StatusCreated)

		response := struct {
			Uid string `json:"uid"`
		}{
			user.Uid.Hex(),
		}
		respBody, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(respBody)
	}
}

/**
 * @api {put} /beatHeart/:uid/ Beat Heart
 * @apiName BeatHeart
 * @apiGroup User
 * @apiVersion 0.1.0
 *
 * @apiDescription Update the current IP address of the user
 *
 * @apiParamExample {json} The body of the request is in json format. Example:
 * {
 *  "currentip": "10.1.2.3"
 * }
 *
 * @apiParam {string} uid User's id
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 204 No content
 *
 * @apiUse UidNotFoundError
 */
func beatHeart(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		userId := pat.Param(r, "uid")

		ip := struct {
			CurrentIp string `json:"currentip"`
		}{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&ip)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c, user := checkForUid(s, userId, w)
		if user.Uid == "" {
			return
		}

		user.CurrentIp = ip.CurrentIp

		err = c.Update(bson.M{"_id": bson.ObjectIdHex(userId)}, &user)
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed update user: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "UidNotFound", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

/**
 * @api {get} /user/getUid/:name Look up uid
 * @apiName GetUserByName
 * @apiGroup User
 * @apiVersion 0.1.0
 *
 * @apiDescription Get <code>uid</code> of the user with given <code>name</code>. If not found,
 * return a UserNameNotFoundError
 *
 * @apiParam {string} name User's name
 *
 * @apiSuccess {string} uid Users-ID
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       [{"uid": "12461"}]
 *     }
 *
 * @apiUse UserNameNotFoundError
 */
func lookUpUid(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		userName := pat.Param(r, "name")

		c := session.DB(DBNAME).C("users")

		var usersList []User
		err := c.Find(bson.M{"name": userName}).All(&usersList)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find users: ", err)
			return
		}

		if len(usersList) == 0 {
			ErrorWithJSON(w, "UserNameNotFound", http.StatusNotFound)
			return
		}

		type returnedStruct struct {
			Uid string `json:"uid"`
		}

		var uidList []returnedStruct
		for _, u := range usersList {
			uidList = append(uidList, returnedStruct{ u.Uid.Hex()})
		}

		respBody, err := json.MarshalIndent(uidList, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

/**
 * @api {get} /user/getIp/:uid Look up ip
 * @apiName GetUserIpByUid
 * @apiGroup User
 * @apiVersion 0.1.0
 *
 * @apiDescription Get <code>currentIp</code> of the user with given <code>uid</code>. If not found,
 * return a UidNotFoundError
 *
 * @apiParam {string} uid User's id
 *
 * @apiSuccess {string} currentIp User's current IP address
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "currentIp": "127.15.1.145"
 *     }
 *
 * @apiUse UidNotFoundError
 */
func lookUpIp(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		userId := pat.Param(r, "uid")

		_, user := checkForUid(session, userId, w)
		if user.Uid == "" {
			return 
		}

		respBody, err := json.MarshalIndent(struct {
			CurrentIp string `json:"current_ip"`
		}{user.CurrentIp}, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

/**
 * @api {post} /invitation/:srcUid/:destUid/:srcRole/:destRole Add invitation
 * @apiName PostInvitation
 * @apiGroup Invitation
 * @apiVersion 0.1.0
 *
 * @apiDescription Add new invitation. If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid destination user id
 * @apiParam {string} srcRole Sender role
 * @apiParam {string} destRole destination role
 *
 * @apiSuccess {string} src Sender user id
 * @apiSuccess {string} dest Target user id
 * @apiSuccess {string} status Invitation status
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 201 Created
 *
 * @apiUse UidNotFoundError
 */
func addInvitation(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		//srcUid := pat.Param(r, "srcUid")
		//destUid := pat.Param(r, "destUid")
		//srcRole := pat.Param(r, "srcRole")
		//destRole := pat.Param(r, "destRole")
		//
		//c, user := checkForUid(session, srcUid, w)
		//
		//c := session.DB(DBNAME).C("relations")
	}
}

/**
 * @api {get} /invitations/:uid Look up invitations
 * @apiName GetReceivedInvitations
 * @apiGroup Invitation
 * @apiVersion 0.1.0
 *
 * @apiDescription Get all pending invitation. If the <code>uid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} uid Destination user's id
 *
 * @apiSuccess {Relation} invitations List of the received invitation
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "invitations": [
 *        	{
 *       	 	"src": "556171",
 *       	 	"role": {
 *       	 		"srcRole": "doctor",
 *       	 		"destRole": "patient"
 *       	 	}
 *          },
 *        	{
 *       	 	"src": "17117",
 *       	 	"role": {
 *       	 		"srcRole": "pharmacist",
 *       	 		"destRole": "patient"
 *       	 	}
 *          }
 *       }
 *     }
 *
 * @apiUse UidNotFoundError
 */
func lookUpInvitations(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

/**
 * @api {delete} /invitation/:srcUid/:destUid/:srcRole/:destRole Remove invitation
 * @apiName DeleteInvitation
 * @apiGroup Invitation
 * @apiVersion 0.1.0
 *
 * @apiDescription Remove relation. If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid Destination user id
 * @apiParam {string} srcRole Sender user role
 * @apiParam {string} destRole Destination user role
 *
 * @apiUse UidNotFoundError
 */
func removeInvitation(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

/**
 * @api {post} /notification Add notification
 * @apiName PostNotification
 * @apiGroup Notification
 * @apiVersion 0.1.0
 *
 * @apiDescription Add new message notification. If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid Target user id
 *
 * @apiSuccess {string} src Sender user id
 * @apiSuccess {string} dest Target user id
 * @apiSuccess {string} dateTime Time of sending notification
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 Created
 *
 * @apiUse UidNotFoundError
 */
func addMsgNotif(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

/**
 * @api {get} /notifications/:uid Look up notifications
 * @apiName GetReceivedNotifications
 * @apiGroup Notification
 * @apiVersion 0.1.0
 *
 * @apiDescription Get all received message notification to user. If the <code>uid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} uid Target user's id
 *
 * @apiSuccess {[]Notif} notifications List of the received notifications
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "notifications": [
 *        	{
 *       	 	"src": "556171",
 *         	 	"dateTime": "2017-11-19 17:15:45"
 *          },
 *        	{
 *       	 	"src": "17117",
 *         	 	"dateTime": "2017-11-20 8:34:22"
 *          },
 *        	{
 *       	 	"src": "364346",
 *         	 	"dateTime": "2017-11-20 11:55:53"
 *          }
 *       ]
 *     }
 *
 * @apiUse UidNotFoundError
 */
func lookUpMsgNotif(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}



/**
	Check if the uid exist. If so, returns the corresponding user and mgo.Collection
 */
func checkForUid(session *mgo.Session, userId string, w http.ResponseWriter) (c *mgo.Collection, user User) {
	if len(userId) != 24 {
		ErrorWithJSON(w, "Error in uid format. Lenght must be 24", http.StatusNotFound)
		return
	}
	c = session.DB(DBNAME).C("users")
	err1 := c.FindId(bson.ObjectIdHex(userId)).One(&user)
	if err1 != nil {
		switch err1 {
		case mgo.ErrNotFound:
			ErrorWithJSON(w, "UidNotFound", http.StatusNotFound)
			return
		default :
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find user: ", err1)
			return
		}
	}
	return
}