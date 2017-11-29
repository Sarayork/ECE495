define({ "api": [
  {
    "type": "delete",
    "url": "/invitation/:srcUid/:destUid/:srcRole/:destRole",
    "title": "Remove invitation",
    "name": "DeleteInvitation",
    "group": "Invitation",
    "version": "0.1.0",
    "description": "<p>Remove relation. If the <code>srcUid</code> or <code>destUid</code> do not exist, return UidNotFoundError</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "srcUid",
            "description": "<p>Sender user id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "destUid",
            "description": "<p>Destination user id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "srcRole",
            "description": "<p>Sender user role</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "destRole",
            "description": "<p>Destination user role</p>"
          }
        ]
      }
    },
    "filename": "./HarnessServer.go",
    "groupTitle": "Invitation",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UidNotFoundError",
            "description": "<p>The <code>uid</code> is not found</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/invitations/:uid",
    "title": "Look up invitations",
    "name": "GetReceivedInvitations",
    "group": "Invitation",
    "version": "0.1.0",
    "description": "<p>Get all pending invitation. If the <code>uid</code> do not exist, return UidNotFoundError</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uid",
            "description": "<p>Destination user's id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Relation",
            "optional": false,
            "field": "invitations",
            "description": "<p>List of the received invitation</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"invitations\": [\n   \t{\n  \t \t\"src\": \"556171\",\n  \t \t\"role\": {\n  \t \t\t\"srcRole\": \"doctor\",\n  \t \t\t\"destRole\": \"patient\"\n  \t \t}\n     },\n   \t{\n  \t \t\"src\": \"17117\",\n  \t \t\"role\": {\n  \t \t\t\"srcRole\": \"pharmacist\",\n  \t \t\t\"destRole\": \"patient\"\n  \t \t}\n     }        \n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./HarnessServer.go",
    "groupTitle": "Invitation",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UidNotFoundError",
            "description": "<p>The <code>uid</code> is not found</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/addInvitation",
    "title": "Add invitation",
    "name": "PostInvitation",
    "group": "Invitation",
    "version": "0.1.0",
    "description": "<p>Add new invitation. If the <code>srcUid</code> or <code>destUid</code> do not exist, return UidNotFoundError</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "srcUid",
            "description": "<p>Sender user id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "destUid",
            "description": "<p>destination user id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "srcRole",
            "description": "<p>Sender role</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "destRole",
            "description": "<p>destination role</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "src",
            "description": "<p>Sender user id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "dest",
            "description": "<p>Target user id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "status",
            "description": "<p>Invitation status</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 201 Created",
          "type": "json"
        }
      ]
    },
    "filename": "./HarnessServer.go",
    "groupTitle": "Invitation",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UidNotFoundError",
            "description": "<p>The <code>uid</code> is not found</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/notifications/:uid",
    "title": "Look up notifications",
    "name": "GetReceivedNotifications",
    "group": "Notification",
    "version": "0.1.0",
    "description": "<p>Get all received message notification to user. If the <code>uid</code> do not exist, return UidNotFoundError</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uid",
            "description": "<p>Target user's id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "[]Notif",
            "optional": false,
            "field": "notifications",
            "description": "<p>List of the received notifications</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"notifications\": [\n   \t{\n  \t \t\"src\": \"556171\",\n    \t \t\"dateTime\": \"2017-11-19 17:15:45\"\n     },\n   \t{\n  \t \t\"src\": \"17117\",\n    \t \t\"dateTime\": \"2017-11-20 8:34:22\"\n     },\n   \t{\n  \t \t\"src\": \"364346\",\n    \t \t\"dateTime\": \"2017-11-20 11:55:53\"\n     }          \n  ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./HarnessServer.go",
    "groupTitle": "Notification",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UidNotFoundError",
            "description": "<p>The <code>uid</code> is not found</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/notification",
    "title": "Add notification",
    "name": "PostNotification",
    "group": "Notification",
    "version": "0.1.0",
    "description": "<p>Add new message notification. If the <code>srcUid</code> or <code>destUid</code> do not exist, return UidNotFoundError</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "srcUid",
            "description": "<p>Sender user id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "destUid",
            "description": "<p>Target user id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "src",
            "description": "<p>Sender user id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "dest",
            "description": "<p>Target user id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "dateTime",
            "description": "<p>Time of sending notification</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 Created",
          "type": "json"
        }
      ]
    },
    "filename": "./HarnessServer.go",
    "groupTitle": "Notification",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UidNotFoundError",
            "description": "<p>The <code>uid</code> is not found</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "put",
    "url": "/beatHeart/:uid",
    "title": "Beat Heart",
    "name": "BeatHeart",
    "group": "User",
    "version": "0.1.0",
    "description": "<p>Update the current IP address of the user</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uid",
            "description": "<p>User's id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "currentIp",
            "description": "<p>User's current IP address</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 204 No content",
          "type": "json"
        }
      ]
    },
    "filename": "./HarnessServer.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UidNotFoundError",
            "description": "<p>The <code>uid</code> is not found</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/user/:name",
    "title": "Look up uid",
    "name": "GetUserByName",
    "group": "User",
    "version": "0.1.0",
    "description": "<p>Get <code>uid</code> of the user with given <code>name</code>. If not found, return a UidNotFoundError</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>User's name</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "uid",
            "description": "<p>Users-ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"uid\": \"12461\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./HarnessServer.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UidNotFoundError",
            "description": "<p>The <code>uid</code> is not found</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/user/:uid",
    "title": "Look up ip",
    "name": "GetUserIpByUid",
    "group": "User",
    "version": "0.1.0",
    "description": "<p>Get <code>currentIp</code> of the user with given <code>uid</code>. If not found, return a UidNotFoundError</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uid",
            "description": "<p>User's id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "currentIp",
            "description": "<p>User's current IP address</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"currentIp\": \"127.15.1.145\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./HarnessServer.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UidNotFoundError",
            "description": "<p>The <code>uid</code> is not found</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/user",
    "title": "Subscribe a new user",
    "name": "PostUser",
    "group": "User",
    "version": "0.1.0",
    "description": "<p>Generate unique Users-ID and add new user account in the HArNESS data base</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>User's name</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>User's password</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>User's email</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "address",
            "description": "<p>User's address</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "phoneNumber",
            "description": "<p>User's phoneNumber</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "roles",
            "description": "<p>User's roles</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "timestamp",
            "description": "<p>User's timestamp</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "currentIp",
            "description": "<p>User's currentIp</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "uid",
            "description": "<p>Users-ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 201 (Created)\n{\n  \"uid\": \"12461\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./HarnessServer.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 5xx": [
          {
            "group": "Error 5xx",
            "optional": false,
            "field": "UserNameAlreadyExist",
            "description": "<p>The <code>userName</code> already exists</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UserNameAlreadyExist\"\n}",
          "type": "json"
        }
      ]
    }
  }
] });
