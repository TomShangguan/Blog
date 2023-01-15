package dao

import (
	"blog/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	Register(user *model.User)
	Login(username string) model.User

	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(uid int) model.Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to open db", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}

func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}

func (mgr *manager) GetPost(uid int) model.Post {
	var post model.Post
	mgr.db.First(&post, uid)
	return post
}
