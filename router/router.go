package router

import (
	"go-frond-end/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	// front user
	r.HandleFunc("/index", middleware.Home)
	r.HandleFunc("/about", middleware.About)
	r.HandleFunc("/posts", middleware.Posts)
	r.HandleFunc("/post/{id}", middleware.Post)

	// // front admin
	r.HandleFunc("/admin/index", middleware.HomeAdmin)
	r.HandleFunc("/admin/about", middleware.AboutAdmin)
	r.HandleFunc("/admin/posts", middleware.PostsAdmin)
	r.HandleFunc("/admin/posts/delete/{id}", middleware.DeletePostAdmin)
	r.HandleFunc("/admin/posts/showupdate/{id}", middleware.ShowUpdatePostAdmin)
	r.HandleFunc("/admin/posts/create", middleware.Create)
	r.HandleFunc("/admin/posts/store", middleware.Store)
	r.HandleFunc("/admin/posts/update", middleware.UpdatePostAdmin)

	// api post
	r.HandleFunc("/api/posts", middleware.ApiPosts).Methods("GET")
	r.HandleFunc("/api/posts", middleware.CreatePost).Methods("POST")
	r.HandleFunc("/api/post/{id}", middleware.ApiPostById).Methods("GET")
	r.HandleFunc("/api/post/{id}", middleware.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/post/{id}", middleware.UpdatePost).Methods("PUT")

	// api user
	r.HandleFunc("/api/users", middleware.ApiUsers).Methods("GET")
	r.HandleFunc("/api/users", middleware.CreateUsers).Methods("POST")
	r.HandleFunc("/api/user/{id}", middleware.ApiUserById).Methods("GET")
	r.HandleFunc("/api/user/{id}", middleware.DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/user/{id}", middleware.UpdateUser).Methods("PUT")

	return r
}
