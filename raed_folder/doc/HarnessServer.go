package main

import(
	"net/http"
)

type User struct{
	id string `json:"uid"`
	name string `json:"name"`
	password string `json:"password"`
	email string `json:"email"`
	address string `json:"address"`
	phoneNumber string `json:"phoneNumber"`
	roles []string `json:"roles"`
	timestamp string `json:"timestamp"`
	currentIp string `json:"currentIp"`
}

/**
 * @api {post} /user Subscribe a new user
 * @apiName PostUser
 * @apiGroup User
 * @apiVersion 0.1.0
 * 
 * @apiDescription Generate unique Users-ID and add new user account in the HArNESS data base
 * 
 * @apiSuccess {string} uid Users-ID
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "uid": "12461",
 *     }
 *     
 * @apiUse UserNameAlreadyExist
 */
func subscribe(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {get} /connect/:id Connect a user
 * @apiName GETUser
 * @apiGroup User
 * @apiVersion 0.1.0
 * 
 * @apiDescription Connect user with his account. If the user exists, return his information.
 * Otherwise, UidNotFoundError will be returned
 *
 * @apiParam {string} uid User's id
 * 
 * @apiSuccess {string} uid Users-ID
 * @apiSuccess {string} name User's name
 * @apiSuccess {string} password User's password (MD5 hash)
 * @apiSuccess {string} email User's email
 * @apiSuccess {string} address User's address
 * @apiSuccess {string} phoneNumber User's phone number
 * @apiSuccess {[]string} roles User's roles (can be multiple)
 * @apiSuccess {string} timestamp Time interval to update user's IP address (in seconds)
 * @apiSuccess {string} currentIp User's current IP address (empty if the was offline before invoking
 * the request)
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "uid": "12461",
 *       "password": "054c1f6789da753c60b7fca1e48fcd13"
 *       "name": "Bob",
 *       "email": "bob@exemple.com",
 *       "address": "Geneva...",
 *       "phoneNumber": "0123456789",
 *       "role": ["pateint", "doctor"],
 *       "timestamp": "120",
 *       "currentIp": ""
 *     }
 *     
 * @apiUse UidNotFoundError
 */
func connect(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {put} /beatHeart/:uid/:currentIp Beat Heart
 * @apiName BeatHeart
 * @apiGroup User
 * @apiVersion 0.1.0
 * 
 * @apiDescription Update the current IP address of the user
 *
 * @apiParam {string} uid User's id
 * @apiParam {string} newIp User's current IP address
 * 
 * @apiSuccess {string} uid Users-ID
 * @apiSuccess {string} ip User's current IP address
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "uid": "12461",
 *       "currentIp": "127.15.1.145"
 *     }
 *     
 * @apiUse UidNotFoundError
 */
func beatHeart(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {put} /disconnect/:uid Disconnect user
 * @apiName PutUser
 * @apiGroup User
 * @apiVersion 0.1.0
 * 
 * @apiDescription Update the current IP address of the user to empty. If the <code>uid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} uid User's id
 *     
 * @apiUse UidNotFoundError
 */
func disconnect(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {get} /user/:name Look up uid
 * @apiName GetUserByName
 * @apiGroup User
 * @apiVersion 0.1.0
 * 
 * @apiDescription Get <code>uid</code> of the user with given <code>name</code>. If not found,
 * return a UidNotFoundError
 *
 * @apiParam {string} name User's name
 * 
 * @apiSuccess {string} name User's name
 * @apiSuccess {string} uid Users-ID
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "name": "raed"
 *       "uid": "12461",
 *     }
 *     
 * @apiUse UidNotFoundError
 */
func lookUpUid(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {get} /user/:uid Look up ip
 * @apiName GetUserIpByUid
 * @apiGroup User
 * @apiVersion 0.1.0
 * 
 * @apiDescription Get <code>currentIp</code> of the user with given <code>uid</code>. If not found,
 * return a UidNotFoundError
 *
 * @apiParam {string} uid User's id
 * 
 * @apiSuccess {string} uid User's id
 * @apiSuccess {string} currentIp User's current IP address
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "uid": "12461",
 *       "currentIp": "127.15.1.145"
 *     }
 *     
 * @apiUse UidNotFoundError
 */
func lookUpIp(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {post} /addInvitation/:srcUid/:destUid Add invitation
 * @apiName PostInvitation
 * @apiGroup Invitation
 * @apiVersion 0.1.0
 * 
 * @apiDescription Add new friendship invitation. If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid Target user id
 * 
 * @apiSuccess {string} src Sender user id
 * @apiSuccess {string} dest Target user id
 * @apiSuccess {string} status Invitation status
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "src": "12461",
 *       "dest": "16548",
 *       "status": "pending"
 *     }
 *     
 * @apiUse UidNotFoundError
 */
func addInvitation(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {get} /invitations/:uid Look up invitations
 * @apiName GetReceivedInvitations
 * @apiGroup Invitation
 * @apiVersion 0.1.0
 * 
 * @apiDescription Get all received friendship invitation to user. If the <code>uid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} uid Target user's id
 * 
 * @apiSuccess {[]Invitation} invitations List of the received invitation
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "invitations": [
 *        	{
 *       	 	"src": "556171",
 *        	 	"dest": "16548",
 *         	 	"status": "pending"
 *          },
 *        	{
 *       	 	"src": "17117",
 *        	 	"dest": "16548",
 *         	 	"status": "pending"
 *          },
 *        	{
 *       	 	"src": "364346",
 *        	 	"dest": "16548",
 *         	 	"status": "pending"
 *          }          
 *       ]
 *     }
 *     
 * @apiUse UidNotFoundError
 */
func lookUpInvitations(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {delete} /invitation/:srcUid/:destUid Remove invitation
 * @apiName DeleteInvitation
 * @apiGroup Invitation
 * @apiVersion 0.1.0
 * 
 * @apiDescription Remove friendship invitation. If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid Target user id
 *     
 * @apiUse UidNotFoundError
 */
func removeInvitation(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {post} /friendship/:uid1/:uid2 Create friendship
 * @apiName PostFriendship
 * @apiGroup Friendship
 * @apiVersion 0.1.0
 * 
 * @apiDescription Add new friendship relation between two users. If the <code>uid1</code> or <code>uid2</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} uid1 User1 id
 * @apiParam {string} uid2 User2 id
 * 
 * @apiSuccess {string} uid1 User1 id
 * @apiSuccess {string} uid2 User2 id
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "uid1": "12461",
 *       "uid2": "16548"
 *     }
 *     
 * @apiUse UidNotFoundError
 */
func addFriendship(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {delete} /friendship/:uid1/:uid2 Remove friendship relation
 * @apiName DeleteFriendship
 * @apiGroup Friendship
 * @apiVersion 0.1.0
 * 
 * @apiDescription Remove friendship relation between two users. If the <code>user1</code> or <code>user2</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} user1 User1 id
 * @apiParam {string} user2 User2 id
 *     
 * @apiUse UidNotFoundError
 */
func removeFriendship(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {post} /notification/:srcUid/:destUid Add notification
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
 *     HTTP/1.1 200 OK
 *     {
 *       "src": "12461",
 *       "dest": "16548",
 *       "dateTime": "2017-11-19 17:15:45"
 *     }
 *     
 * @apiUse UidNotFoundError
 */
func addMsgNotif(w http.ResponseWriter, req *http.Request) {
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
 *        	 	"dest": "16548",
 *         	 	"dateTime": "2017-11-19 17:15:45"
 *          },
 *        	{
 *       	 	"src": "17117",
 *        	 	"dest": "16548",
 *         	 	"dateTime": "2017-11-20 8:34:22"
 *          },
 *        	{
 *       	 	"src": "364346",
 *        	 	"dest": "16548",
 *         	 	"dateTime": "2017-11-20 11:55:53"
 *          }          
 *       ]
 *     }
 *     
 * @apiUse UidNotFoundError
 */
func lookUpMsgNotif(w http.ResponseWriter, req *http.Request) {
}