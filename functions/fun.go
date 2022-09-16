package functions

import (
	"go-frond-end/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("post.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})

	return db
}

func GetPosts() *[]model.Post {
	// connect to database
	db := ConnectDb()

	// get data from database
	var posts []model.Post
	db.Find(&posts)

	return &posts
}

func GetPostById(id string) *model.Post {
	// connect db
	db := ConnectDb()

	// get data from database
	var post model.Post
	db.First(&post, id)
	return &post
}

func CreatedPost(title string, content string) *model.Post {
	// parsing data to post
	var post model.Post
	post.Title = title
	post.Content = content

	// add to database
	db := ConnectDb()
	db.Create(&post)
	return &post
}

func DeletedPost(id string) {
	// connect to database
	db := ConnectDb()

	// delete data from database
	var post model.Post
	db.Delete(&post, id)
}

func UpdatedPost(id string, title string, content string) *model.Post {
	// connect to database
	db := ConnectDb()

	// parsing data to post
	var post model.Post
	post.Title = title
	post.Content = content

	// update data to database
	db.Model(&post).Where("id = ?", id).Updates(post)
	return &post
}

func CreatedUser(name string, username string, password string) *model.User {
	// parsing data to user
	var user model.User
	user.Name = name
	user.Username = username
	user.Password = password

	// add to database
	db := ConnectDb()
	db.Create(&user)
	return &user
}

func GetUsers() *[]model.User {
	// connect to database
	db := ConnectDb()

	// get data from database
	var users []model.User
	db.Find(&users)

	return &users
}

func GetUserById(id string) *model.User {
	// connect db
	db := ConnectDb()

	// get data from database
	var user model.User
	db.First(&user, id)
	return &user
}

func DeletedUser(id string) {
	// connect to database
	db := ConnectDb()

	// delete data from database
	var user model.User
	db.Delete(&user, id)
}

func UpdatedUser(id string, name string, username string, password string) *model.User {
	// connect to database
	db := ConnectDb()

	// parsing data to user
	var user model.User
	user.Name = name
	user.Username = username
	user.Password = password

	// update data to database
	db.Model(&user).Where("id = ?", id).Updates(user)
	return &user
}
