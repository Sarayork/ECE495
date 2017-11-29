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
	roles [] string `json:"roles"`
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
 * @apiParam {string} name User's name
 * @apiParam {string} password User's password
 * @apiParam {string} email User's email
 * @apiParam {string} address User's address
 * @apiParam {string} phoneNumber User's phoneNumber
 * @apiParam {string} roles User's roles
 * @apiParam {string} timestamp User's timestamp
 * @apiParam {string} currentIp User's currentIp
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
func subscribe(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {put} /beatHeart/:uid Beat Heart
 * @apiName BeatHeart
 * @apiGroup User
 * @apiVersion 0.1.0
 * 
 * @apiDescription Update the current IP address of the user
 *
 * @apiParam {string} uid User's id
 * @apiParam {string} currentIp User's current IP address
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 204 No content
 *     
 * @apiUse UidNotFoundError
 */
func beatHeart(w http.ResponseWriter, req *http.Request) {
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
 * @apiSuccess {string} uid Users-ID
 * 
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "uid": "12461"
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
func lookUpIp(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {post} /addInvitation Add invitation
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
func addInvitation(w http.ResponseWriter, req *http.Request) {
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
func lookUpInvitations(w http.ResponseWriter, req *http.Request) {
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
func removeInvitation(w http.ResponseWriter, req *http.Request) {
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
func lookUpMsgNotif(w http.ResponseWriter, req *http.Request) {
}
