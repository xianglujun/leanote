package controllers

import (
	"github.com/revel/revel"
//	"strings"
//	"time"
	//	"encoding/json"
//	"github.com/leanote/leanote/app/info"
//	. "github.com/leanote/leanote/app/lea"
//	"github.com/leanote/leanote/app/lea/blog"
//	"gopkg.in/mgo.v2/bson"
	//	"github.com/leanote/leanote/app/types"
	//	"io/ioutil"
	//	"math"
	//	"os"
	//	"path"
)

type Preview struct {
	Blog
}

// 得到要预览的主题绝对路径
func (c Preview) getPreviewThemeAbsolutePath(themeId string) bool {
	if themeId != "" {
		c.Session["themeId"] = themeId // 存到session中, 下次的url就不能带了
	} else {
		themeId = c.Session["themeId"] // 直接从session中获取
	}
	if themeId == "" {
		return false
	}
	theme := themeService.GetTheme(c.GetUserId(), themeId)
	
	c.RenderArgs["isPreview"] = true
	c.RenderArgs["themePath"] = theme.Path
	if theme.Path == "" {
		return false
	}
	return true
}

func (c Preview) Index(userIdOrEmail string, themeId string) revel.Result {
	if !c.getPreviewThemeAbsolutePath(themeId) {
		return c.E404()
	}
	return c.Blog.Index(c.GetUserId())
//	return blog.RenderTemplate("index.html", c.RenderArgs, c.getPreviewThemeAbsolutePath(themeId))
}

func (c Preview) Tag(userIdOrEmail, tag string) revel.Result {
	if !c.getPreviewThemeAbsolutePath("") {
		return c.E404()
	}
	return c.Blog.Tag(c.GetUserId(), tag)
}
func (c Preview) Tags(userIdOrEmail string) revel.Result {
	if !c.getPreviewThemeAbsolutePath("") {
		return c.E404()
	}
	return c.Blog.Tags(c.GetUserId())
//	if tag == "" {
//		return blog.RenderTemplate("tags.html", c.RenderArgs, c.getPreviewThemeAbsolutePath(""))
//	}
//	return blog.RenderTemplate("tag_posts.html", c.RenderArgs, c.getPreviewThemeAbsolutePath(""))
}
func (c Preview) Archive(userIdOrEmail string, notebookId string) revel.Result {
	if !c.getPreviewThemeAbsolutePath("") {
		return c.E404()
	}
	return c.Blog.Archive(c.GetUserId(), notebookId)
//	return blog.RenderTemplate("archive.html", c.RenderArgs, c.getPreviewThemeAbsolutePath(""))
}
func (c Preview) Cate(notebookId string) revel.Result {
	if !c.getPreviewThemeAbsolutePath("") {
		return c.E404()
	}
	return c.Blog.Cate(notebookId)
//	return blog.RenderTemplate("cate.html", c.RenderArgs, c.getPreviewThemeAbsolutePath(""))
}
func (c Preview) Post(noteId string) revel.Result {
	if !c.getPreviewThemeAbsolutePath("") {
		return c.E404()
	}
	return c.Blog.Post(noteId)
//	return blog.RenderTemplate("view.html", c.RenderArgs, c.getPreviewThemeAbsolutePath(""))
}
func (c Preview) Single(singleId string) revel.Result {
	if !c.getPreviewThemeAbsolutePath("") {
		return c.E404()
	}
	return c.Blog.Single(singleId)
//	return blog.RenderTemplate("single.html", c.RenderArgs, c.getPreviewThemeAbsolutePath(""))
}
func (c Preview) Search(userIdOrEmail, keywords string) revel.Result {
	if !c.getPreviewThemeAbsolutePath("") {
		return c.E404()
	}
	return c.Blog.Search(c.GetUserId(), keywords)
//	return blog.RenderTemplate("search.html", c.RenderArgs, c.getPreviewThemeAbsolutePath(""))
}