package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"log"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"goji.io/pat"
	"fmt"
)

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

		c := session.DB(DBNAME).C(USERSCOLL)

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

		c := session.DB(DBNAME).C(USERSCOLL)

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

		type ReturnedStruct struct {
			Uid string `json:"uid"`
		}

		var uidList []ReturnedStruct
		for _, u := range usersList {
			uidList = append(uidList, ReturnedStruct{ u.Uid.Hex()})
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
 * @api {post} /relation/:srcUid/:destUid/:srcRole/:destRole Add relation
 * @apiName PostRelation
 * @apiGroup Relation
 * @apiVersion 0.1.0
 *
 * @apiDescription Add new relation with status "pending". If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid destination user id
 * @apiParam {string} srcRole Sender role
 * @apiParam {string} destRole destination role
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 201 Created
 *
 * @apiUse UidNotFoundError
 */
func addRelation(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		srcUid := pat.Param(r, "srcUid")
		destUid := pat.Param(r, "destUid")
		srcRole := pat.Param(r, "srcRole")
		destRole := pat.Param(r, "destRole")

		_, srcUser := checkForUid(session, srcUid, w)
		if srcUser.Uid == "" {
			return
		}
		_, destUser := checkForUid(session, destUid, w)
		if destUser.Uid == "" {
			return
		}
		if !userHaveRole(srcUser, srcRole) {
			ErrorWithJSON(w, "RoleNotFound", http.StatusNotFound)
			return
		}
		if !userHaveRole(destUser, destRole) {
			ErrorWithJSON(w, "RoleNotFound", http.StatusNotFound)
			return
		}

		c := session.DB(DBNAME).C(RELATIONSCOLL)
		var invitations []Relation
		err := c.Find(bson.M{
			"srcuid": srcUid,
			"destuid": destUid,
			"srcrole": srcRole,
			"destrole": destRole}).All(&invitations)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find invitations: ", err)
			return
		}
		if len(invitations) != 0 {
			ErrorWithJSON(w, "RelationAlreadyExist", http.StatusBadRequest)
			return
		}

		relation := Relation{
			SrcUid: srcUid,
			DestUid: destUid,
			SrcRole: srcRole,
			DestRole: destRole,
			Status: "pending",
		}
		err = c.Insert(relation)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed relation user: ", err)
			return
		}
	}
}

/**
 * @api {get} /invitations/:uid Look up invitations
 * @apiName GetPendingRelations
 * @apiGroup Relation
 * @apiVersion 0.1.0
 *
 * @apiDescription Get all pending relations. If the <code>uid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} uid Destination user's id
 *
 * @apiSuccess {Relation} invitations List of the relations with status "pending"
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "invitations": [
 *        	{
 *       	 	"src": "556171",
 *       	 	"roles": {
 *       	 		"srcRole": "doctor",
 *       	 		"destRole": "patient"
 *       	 	}
 *          },
 *        	{
 *       	 	"src": "17117",
 *       	 	"roles": {
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
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		destUid := pat.Param(r, "uid")

		_, user := checkForUid(session, destUid, w)
		if user.Uid == "" {
			return
		}

		c := session.DB(DBNAME).C(RELATIONSCOLL)

		var invitations []Relation
		err := c.Find(bson.M{"destuid": destUid, "status": "pending"}).All(&invitations)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find invitations: ", err)
			return
		}

		type Roles struct {
			SrcRole string `json:"srcRole"`
			DestRole string `json:"destRole"`
		}
		type ReturnedStruct struct {
			Src string `json:"src"`
			Roles Roles `json:"roles"`
		}

		var invitationsList []ReturnedStruct

		for _, i := range invitations {
			invitationsList = append(invitationsList,
				ReturnedStruct{
					Src: i.SrcUid,
					Roles: Roles{
						SrcRole: i.SrcRole,
						DestRole: i.DestRole,
					},
				})
		}

		respBody, err := json.MarshalIndent(invitationsList, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

/**
 * @api {put} /relation/:srcUid/:destUid/:srcRole/:destRole Accept relation
 * @apiName PutRelation
 * @apiGroup Relation
 * @apiVersion 0.1.0
 *
 * @apiDescription Change the relation status to "accepted". If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid Destination user id
 * @apiParam {string} srcRole Sender user role
 * @apiParam {string} destRole Destination user role
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 Ok
 *
 *
 * @apiUse UidNotFoundError
 */
func acceptRelation(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		srcUid := pat.Param(r, "srcUid")
		destUid := pat.Param(r, "destUid")
		srcRole := pat.Param(r, "srcRole")
		destRole := pat.Param(r, "destRole")

		c := session.DB(DBNAME).C(RELATIONSCOLL)

		var relation Relation
		err := c.Find(bson.M{
			"srcuid": srcUid,
			"destuid": destUid,
			"srcrole": srcRole,
			"destrole": destRole,
			"status": "pedding",
			}).One(&relation)

		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed update relation: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "relation not found", http.StatusNotFound)
				return
			}
		}

		relation.Status = "accepted"
		fmt.Println(relation.Id)

		err = c.Update(bson.M{"_id": relation.Id}, &relation)
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed update relation: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "relation not found", http.StatusNotFound)
				return
			}
		}

		ResponseWithNoContent(w, http.StatusOK)
	}
}

/**
 * @api {delete} /relation/:srcUid/:destUid/:srcRole/:destRole Remove relation
 * @apiName DeleteRelation
 * @apiGroup Relation
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
func removeRelation(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
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
	Check if the uid exist. If so, returns the corresponding user struct and mgo.Collection for user
 */
func checkForUid(session *mgo.Session, userId string, w http.ResponseWriter) (c *mgo.Collection, user User) {
	if len(userId) != 24 {
		ErrorWithJSON(w, "Error in uid format. Lenght must be 24", http.StatusNotFound)
		return
	}
	c = session.DB(DBNAME).C(USERSCOLL)
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

/**
	Check if the role is in the given list of roles. If so, returns true
 */
func userHaveRole(user User, role string) (b bool){
	b = true
	for _, e := range user.Roles {
		if e == role {
			return
		}
	}
	b = false
	return
}