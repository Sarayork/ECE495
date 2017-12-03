/////////////////////////////////////////
///	errors definitions
/////////////////////////////////////////
/** 
 * @apiDefine UserNameAlreadyExist
 * @apiError (Error 5xx) UserNameAlreadyExist The <code>userName</code> already exists
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 400 Not Found
 *     {
 *       "error": "UserNameAlreadyExist"
 *     }
 */

/** 
 * @apiDefine UidNotFoundError
 * @apiError (Error 4xx) UidNotFound The <code>uid</code> is not found
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error": "UidNotFound"
 *     }
 */

/** 
 * @apiDefine UserNameNotFoundError
 * @apiError (Error 4xx) UerNameNotFound The <code>name</code> is not found
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error": "UerNameNotFound"
 *     }
 */
