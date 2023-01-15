package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)

	c.Redirect(301, "/")
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "Username doesn't exist")
		fmt.Println("Username doesn't exist")
	} else {
		if u.Password != password {
			fmt.Println("Wrong password")
			c.HTML(200, "login.html", "Wrong password")
		} else {
			fmt.Println("Login Sucess")
			c.Redirect(301, "/")
		}
	}
}

func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func PostDetail(c *gin.Context) {
	s := c.Query("uid")
	uid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(uid)

	content := blackfriday.Run([]byte(p.Content))

	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
