package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewChiRouter(healthHandler HealthHandler, followHandler FollowerHandler, userHandler UserHandler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))

	apiRouter := chi.NewRouter()
	apiRouter.Post("/users", userHandler.CreateUser)
	apiRouter.Get("/users", userHandler.GetUsers)
	apiRouter.Get("/users/{user_id}", userHandler.GetUserInfo)

	apiRouter.Post("/follow", followHandler.CreateFollow)
	apiRouter.Get("/followers/@me", followHandler.ListUserFollowers)
	apiRouter.Get("/followee/@me", followHandler.ListUserFollowee)

	apiRouter.Mount("/debug/", middleware.Profiler())

	apiHealthRouter := chi.NewRouter()
	apiHealthRouter.Get("/", healthHandler.IndexController)

	router.Use(Recovery)

	router.Mount("/health", apiHealthRouter)
	router.Mount("/api/v1/", apiRouter)

	return router
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				err, ok := rvr.(error)
				if ok {
					log.Error(err)
					jsonBody, _ := json.Marshal(map[string]string{
						"error": "Internal server error",
					})

					writer.Header().Set("Content-Type", "application/json")
					writer.WriteHeader(http.StatusInternalServerError)
					_, _ = writer.Write(jsonBody)
				}
			}
		}()

		next.ServeHTTP(writer, request)
	})
}
