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
 * @apiError (Error 4xx) UidNotFoundError The <code>uid</code> is not found
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error": "UidNotFoundError"
 *     }
 */