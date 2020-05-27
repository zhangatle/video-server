package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zhangatle/video_server/scheduler/dbops"
	"net/http"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, 400, "video id shoud not by empty")
		return
	}
	err := dbops.AddvideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}
	sendResponse(w, 200, "")
	return
}
