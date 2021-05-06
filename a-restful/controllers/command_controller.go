package controller

import (
	"net/http"
	robotsdk "robot-api-golang/a-restful/interface"
	util "robot-api-golang/a-restful/utils"
)

var GetCommandStatus = func(w http.ResponseWriter, r *http.Request) {
	status := robotsdk.Robot.CurrentState()
	resp := util.Message(true, status)
	resp["status"] = "robot status is g"
	util.Respond(w, resp)
}

var PostCommands = func(w http.ResponseWriter, r *http.Request) {
	robotsdk.Robot.EnqueueTask()
	resp := util.Message(true, "Successful - queue command")
	util.Respond(w, resp)
}

var PostCancelCommands = func(w http.ResponseWriter, r *http.Request) {
	robotsdk.Robot.CancelTask()
	resp := util.Message(true, "Successful - canceled command")
	util.Respond(w, resp)
}
