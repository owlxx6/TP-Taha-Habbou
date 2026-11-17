package main

import "net/http"

func getCat(req *http.Request) (int, any) {

	catId := req.PathValue("catId")
	Logger.Info("Searching for " + catId)

	aCat, found := catsDatabase[catId]

	if found {
		return http.StatusOK, aCat
	} else {
		Logger.Info("Cat not found")
		return http.StatusNotFound, "Cat not found"
	}
}
func deleteCat(req *http.Request) (int, any) {

	catId := req.PathValue("catId")
	Logger.Info("Searching for " + catId)

	unChat, found := catsDatabase[catId]
	unChat.miaule()

	if found {
		delete(catsDatabase, catId)
		return http.StatusNoContent, nil
	} else {
		Logger.Info("Cat not found")
		return http.StatusNotFound, "Cat not found"
	}
}
