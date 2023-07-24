# Introduction
The Kick Chat Room Logger is a lightweight application designed to log messages from Kick chat rooms. It utilizes txt files 
to efficiently store each message with relevant information.

# How it works
Upon execution, the application automatically joins all the Kick channels listed in the channels.json file. Additionally, it starts a small API built with Gin, 
which enables easy access to the logged messages. When a message is sent in a chat room, it is saved to the appropriate location in the following format: 
`/logs/users/{username}/{channelid}/logs.txt`. Each line in the log file contains the date, message, badges, and username color, all separated by the tab (\t) 
character, making it easy to parse and use through the API.

# Configuration
1. Create a file named channels.json in the root folder of the application.
2. Inside channels.json, define the Kick chat rooms you want to log as objects, where each object's key is the Kick chat room ID, and the properties include
   the username of the user associated with that channel. Optionally, you can include the stvID, which represents the 7tv profile ID used for frontend purposes
   (e.g., replacing words with emotes). Below is an example:

```
{
    "668": {
        "username": "xqc",
        "stvID": "612bbd809d60d44d5240b3dc"
    },
    "669": {
        "username": "example_user"
    }
}
```
