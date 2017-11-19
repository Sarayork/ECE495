define({ "api": [
  {
    "type": "delete",
    "url": "/friendship/:uid1/:uid2",
    "title": "Remove friendship relation",
    "name": "DeleteFriendship",
    "group": "Friendship",
    "version": "0.1.0",
    "description": "<p>Remove friendship relation between two users. If the <code>user1</code> or <code>user2</code> do not exist, return UidNotFoundError</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "user1",
            "description": "<p>User1 id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "user2",
            "description": "<p>User2 id</p>"
          }
        ]
      }
    },
    "filename": "./doc/HarnessServer.go",
    "groupTitle": "Friendship",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/friendship/:uid1/:uid2",
    "title": "Create friendship",
    "name": "PostFriendship",
    "group": "Friendship",
    "version": "0.1.0",
    "description": "<p>Add new friendship relation between two users. If the <code>uid1</code> or <code>uid2</code> do not exist, return UidNotFoundError</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uid1",
            "description": "<p>User1 id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uid2",
            "description": "<p>User2 id</p>"
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
            "field": "uid1",
            "description": "<p>User1 id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "uid2",
            "description": "<p>User2 id</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"uid1\": \"12461\",\n  \"uid2\": \"16548\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
    "groupTitle": "Friendship",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "delete",
    "url": "/invitation/:srcUid/:destUid",
    "title": "Remove invitation",
    "name": "DeleteInvitation",
    "group": "Invitation",
    "version": "0.1.0",
    "description": "<p>Remove friendship invitation. If the <code>srcUid</code> or <code>destUid</code> do not exist, return UidNotFoundError</p>",
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
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
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
    "description": "<p>Get all received friendship invitation to user. If the <code>uid</code> do not exist, return UidNotFoundError</p>",
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
            "type": "[]Invitation",
            "optional": false,
            "field": "invitations",
            "description": "<p>List of the received invitation</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"invitations\": [\n   \t{\n  \t \t\"src\": \"556171\",\n   \t \t\"dest\": \"16548\",\n    \t \t\"status\": \"pending\"\n     },\n   \t{\n  \t \t\"src\": \"17117\",\n   \t \t\"dest\": \"16548\",\n    \t \t\"status\": \"pending\"\n     },\n   \t{\n  \t \t\"src\": \"364346\",\n   \t \t\"dest\": \"16548\",\n    \t \t\"status\": \"pending\"\n     }          \n  ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/addInvitation/:srcUid/:destUid",
    "title": "Add invitation",
    "name": "PostInvitation",
    "group": "Invitation",
    "version": "0.1.0",
    "description": "<p>Add new friendship invitation. If the <code>srcUid</code> or <code>destUid</code> do not exist, return UidNotFoundError</p>",
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
            "field": "status",
            "description": "<p>Invitation status</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"src\": \"12461\",\n  \"dest\": \"16548\",\n  \"status\": \"pending\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
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
          "content": "HTTP/1.1 200 OK\n{\n  \"notifications\": [\n   \t{\n  \t \t\"src\": \"556171\",\n   \t \t\"dest\": \"16548\",\n    \t \t\"dateTime\": \"2017-11-19 17:15:45\"\n     },\n   \t{\n  \t \t\"src\": \"17117\",\n   \t \t\"dest\": \"16548\",\n    \t \t\"dateTime\": \"2017-11-20 8:34:22\"\n     },\n   \t{\n  \t \t\"src\": \"364346\",\n   \t \t\"dest\": \"16548\",\n    \t \t\"dateTime\": \"2017-11-20 11:55:53\"\n     }          \n  ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/notification/:srcUid/:destUid",
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
          "content": "HTTP/1.1 200 OK\n{\n  \"src\": \"12461\",\n  \"dest\": \"16548\",\n  \"dateTime\": \"2017-11-19 17:15:45\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "put",
    "url": "/beatHeart/:uid/:currentIp",
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
            "field": "newIp",
            "description": "<p>User's current IP address</p>"
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
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "ip",
            "description": "<p>User's current IP address</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"uid\": \"12461\",\n  \"currentIp\": \"127.15.1.145\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/connect/:id",
    "title": "Connect a user",
    "name": "GETUser",
    "group": "User",
    "version": "0.1.0",
    "description": "<p>Connect user with his account. If the user exists, return his information. Otherwise, UidNotFoundError will be returned</p>",
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
            "field": "uid",
            "description": "<p>Users-ID</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>User's name</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>User's password (MD5 hash)</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>User's email</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "address",
            "description": "<p>User's address</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "phoneNumber",
            "description": "<p>User's phone number</p>"
          },
          {
            "group": "Success 200",
            "type": "[]string",
            "optional": false,
            "field": "roles",
            "description": "<p>User's roles (can be multiple)</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "timestamp",
            "description": "<p>Time interval to update user's IP address (in seconds)</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "currentIp",
            "description": "<p>User's current IP address (empty if the was offline before invoking the request)</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"uid\": \"12461\",\n  \"password\": \"054c1f6789da753c60b7fca1e48fcd13\"\n  \"name\": \"Bob\",\n  \"email\": \"bob@exemple.com\",\n  \"address\": \"Geneva...\",\n  \"phoneNumber\": \"0123456789\",\n  \"role\": [\"pateint\", \"doctor\"],\n  \"timestamp\": \"120\",\n  \"currentIp\": \"\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
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
            "field": "name",
            "description": "<p>User's name</p>"
          },
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
          "content": "HTTP/1.1 200 OK\n{\n  \"name\": \"raed\"\n  \"uid\": \"12461\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
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
            "field": "uid",
            "description": "<p>User's id</p>"
          },
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
          "content": "HTTP/1.1 200 OK\n{\n  \"uid\": \"12461\",\n  \"currentIp\": \"127.15.1.145\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
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
          "content": "HTTP/1.1 200 OK\n{\n  \"uid\": \"12461\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 500 Not Found\n{\n  \"error\": \"UserNameAlreadyExist\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "put",
    "url": "/disconnect/:uid",
    "title": "Disconnect user",
    "name": "PutUser",
    "group": "User",
    "version": "0.1.0",
    "description": "<p>Update the current IP address of the user to empty. If the <code>uid</code> do not exist, return UidNotFoundError</p>",
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
    "filename": "./doc/HarnessServer.go",
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
          "content": "HTTP/1.1 400 Not Found\n{\n  \"error\": \"UidNotFoundError\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "_home_raed_go_src_ECE495_raed_folder_doc_main_js",
    "groupTitle": "_home_raed_go_src_ECE495_raed_folder_doc_main_js",
    "name": ""
  }
] });
