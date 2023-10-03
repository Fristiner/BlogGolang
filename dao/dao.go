package dao

import (
	"Blog/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	Register(user *model.User)
	Login(username string) model.User

	// 博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
}

type manager struct {
	db *gorm.DB
}

func (mgr *manager) Login(username string) model.User {
	var user model.User

	mgr.db.Where("user_name=?", username).First(&user)
	return user
}

func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}

var Mgr Manager

func init() {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	_ = db.AutoMigrate(&model.User{})
	_ = db.AutoMigrate(&model.Post{})
}
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}
func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}
func (mgr *manager) GetPost(pid int) model.Post {
	var post model.Post
	mgr.db.First(&post, pid)
	return post
}
