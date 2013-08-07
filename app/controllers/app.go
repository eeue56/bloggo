package controllers

import (
    "github.com/robfig/revel"
    "bloggo/app/models"
    "fmt"
    "encoding/json"
)

type App struct {
	*revel.Controller
}

func (c App) Create() revel.Result {
    var post models.BlogPost

    for k, _ := range c.Params.Values {
        json.Unmarshal([]byte(k), &post)
        break
    }
    fmt.Println("here")
    fmt.Println(post.Title)
    fmt.Println(post.Content)

    return c.Redirect("/")
}

func (c App) NextPost(id int) revel.Result {
    c.RenderArgs(models.GetLatestPost());
}

func (c App) PreviousPost(id int) revel.Result {
    c.RenderArgs(models.GetLatestPost());
}

func (c App) Write() revel.Result {

    return c.Render()
}

func (c App) Login(user *models.User) revel.Result {
    user.Validate(c.Validation)

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect("/")
    }

    c.Session["logged_in"] = "true";

    return c.Redirect("/") 
}

func (c App) Logout() revel.Result {
    c.Session["logged_in"] = "";

    return c.Redirect("/") 
}

func (c App) Index() revel.Result {
    c.RenderArgs["latest"] = models.GetLatestPost();
    c.RenderArgs["logged_in"] = c.Session["logged_in"]
	return c.Render()
}
