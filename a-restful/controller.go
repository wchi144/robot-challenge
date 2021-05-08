package main

import (
	util "lindachin/robot-api-golang/utils"
	"net/http"
)

var GetCommandStatus = func(w http.ResponseWriter, r *http.Request) {
	var s Robot = Spot{}
	status := s.CurrentState()
	resp := util.Message(true, status.GetString())
	resp["status"] = "robot status is g"
	util.Respond(w, resp)
}

var PostCommands = func(w http.ResponseWriter, r *http.Request) {
	var s Robot = Spot{}
	s.EnqueueTask("N E")
	resp := util.Message(true, "Successful - queue command")
	util.Respond(w, resp)
}

var PostCancelCommands = func(w http.ResponseWriter, r *http.Request) {
	var s Robot = Spot{}
	s.CancelTask("123")
	resp := util.Message(true, "Successful - canceled command")
	util.Respond(w, resp)
}
