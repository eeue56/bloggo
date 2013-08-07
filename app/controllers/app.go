package controllers

import (
    "github.com/robfig/revel"
    "bloggo/app/models"
)

type App struct {
	*revel.Controller
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
