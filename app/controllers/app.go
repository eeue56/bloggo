package controllers

import (
    "github.com/robfig/revel"
    "bloggo/app/models"
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

    post.SavePost()

    return c.Redirect("/")
}

func (c App) NextPost(id int) revel.Result {
    models.Connect();
    return c.RenderJson(models.GetLatestPost());
}

func (c App) PreviousPost(id int) revel.Result {
    models.Connect();
    return c.RenderJson(models.GetLatestPost());
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
    models.Connect();
    c.RenderArgs["latest"] = models.GetLatestPost();
    c.RenderArgs["logged_in"] = c.Session["logged_in"]
	return c.Render()
}
