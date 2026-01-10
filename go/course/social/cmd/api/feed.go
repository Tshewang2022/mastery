package main

import "net/http"

// func getUserFeedHandler(req, res)=>{}
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: pagination, filters
	ctx := r.Context()
	feed, err := app.store.Posts.GetUserFeed(ctx, int64(103))

	if err != nil {
		app.internalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
	}
}
