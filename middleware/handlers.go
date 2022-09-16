package middleware

import (
	"encoding/json"
	"go-frond-end/functions"
	"go-frond-end/model"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// user
func Home(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/main.html",
		"views/_header.html",
		"views/_footer.html",
		"views/nav.html",
		"views/home.html",
	))

	var err = tmpl.ExecuteTemplate(w, "index", "home")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/about.html",
		"views/_header.html",
		"views/_footer.html",
		"views/nav.html",
	))

	var err = tmpl.ExecuteTemplate(w, "about", "about")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Posts(w http.ResponseWriter, r *http.Request) {

	var posts = functions.GetPosts()
	var data = model.M{
		"title": "Posts",
		"posts": posts,
	}
	var tmpl = template.Must(template.ParseFiles(
		"views/posts.html",
		"views/_header.html",
		"views/_footer.html",
		"views/nav.html",
	))

	var err = tmpl.ExecuteTemplate(w, "posts", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/post.html",
		"views/_header.html",
		"views/_footer.html",
		"views/nav.html",
	))
	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	// get data from database
	var post = functions.GetPostById(id)

	var data = model.M{
		"title": "Post",
		"post":  post,
	}

	var err = tmpl.ExecuteTemplate(w, "post", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// admin
func HomeAdmin(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/admin/home_admin.html",
		"views/admin/_header.html",
		"views/admin/_footer.html",
		"views/admin/nav.html",
	))
	data := model.Page{Title: "Admin Home"}

	var err = tmpl.ExecuteTemplate(w, "home_admin", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutAdmin(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/admin/about_admin.html",
		"views/admin/_header.html",
		"views/admin/_footer.html",
		"views/admin/nav.html",
	))
	data := model.Page{Title: "About Admin"}

	var err = tmpl.ExecuteTemplate(w, "about_admin", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostsAdmin(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/admin/posts_admin.html",
		"views/admin/_header.html",
		"views/admin/_footer.html",
		"views/admin/nav.html",
	))
	var posts = functions.GetPosts()
	data := model.M{
		"Title": "Posts Admin",
		"Posts": posts,
	}

	var err = tmpl.ExecuteTemplate(w, "posts_admin", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeletePostAdmin(w http.ResponseWriter, r *http.Request) {

	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	functions.DeletedPost(id)

	http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
}

func ShowUpdatePostAdmin(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/admin/show_update.html",
		"views/admin/_header.html",
		"views/admin/_footer.html",
		"views/admin/nav.html",
	))

	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	var post = functions.GetPostById(id)
	data := model.M{
		"Title": "Update Posts Admin",
		"Post":  post,
	}

	var err = tmpl.ExecuteTemplate(w, "show_update", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/admin/new_post.html",
		"views/admin/_header.html",
		"views/admin/_footer.html",
		"views/admin/nav.html",
	))
	data := model.Page{Title: "New Post Admin"}

	var err = tmpl.ExecuteTemplate(w, "new_post", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Store(w http.ResponseWriter, r *http.Request) {
	// get data from formValue
	title := r.FormValue("title")
	content := r.FormValue("content")

	functions.CreatedPost(title, content)

	http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
}

// api
func ApiPosts(w http.ResponseWriter, r *http.Request) {
	var posts = functions.GetPosts()
	if posts == nil {
		http.Error(w, "data not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// get data from formValue
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post = functions.CreatedPost(title, content)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func ApiPostById(w http.ResponseWriter, r *http.Request) {
	// get id
	vars := mux.Vars(r)
	id := vars["id"]
	var post = functions.GetPostById(id)
	// condition id not found
	if post.ID == 0 {
		http.Error(w, "data not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	functions.DeletedPost(id)

	// show the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("deleted")
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	// get data from formValue
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post = functions.UpdatedPost(id, title, content)

	// show the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func UpdatePostAdmin(w http.ResponseWriter, r *http.Request) {

	// get data from formValue
	title := r.FormValue("title")
	id := r.FormValue("id")
	content := r.FormValue("content")

	functions.UpdatedPost(id, title, content)

	http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
}

// api user
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	// get data from formValue
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")

	var user = functions.CreatedUser(name, username, password)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func ApiUsers(w http.ResponseWriter, r *http.Request) {
	var users = functions.GetUsers()
	if users == nil {
		http.Error(w, "data not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func ApiUserById(w http.ResponseWriter, r *http.Request) {
	// get id
	vars := mux.Vars(r)
	id := vars["id"]
	var user = functions.GetUserById(id)
	// condition id not found
	if user.ID == 0 {
		http.Error(w, "data not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	functions.DeletedUser(id)

	// show the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	// get data from formValue
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")

	var user = functions.UpdatedUser(id, name, username, password)

	// show the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
