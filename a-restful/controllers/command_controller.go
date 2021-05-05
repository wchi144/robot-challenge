package controller

import (
	"net/http"
)

var GetCommandStatus = func(w http.ResponseWriter, r *http.Request) {
	status := robotsdk.Robot.CurrentState()
	// resp := util.Message(true, status)
	// resp["data"] = users
	// util.Respond(w, resp)
}

var PostCommands = func(w http.ResponseWriter, r *http.Request) {
	status := robotsdk.Robot.EnqueueTask()
	// resp := util.Message(true, "Successful")
	// resp["data"] = users
	// util.Respond(w, resp)
}

var PostCancelCommands = func(w http.ResponseWriter, r *http.Request) {
	status := robotsdk.Robot.CancelTask()
	// resp := util.Message(true, "Successful")
	// resp["data"] = users
	// util.Respond(w, resp)
}
