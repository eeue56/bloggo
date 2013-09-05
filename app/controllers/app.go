package controllers

import (
    "github.com/robfig/revel"
    "bloggo/app/models"
    "encoding/json"
    "time"
    "fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Create() revel.Result {
    var post models.BlogPost
    models.Connect();

    var out string;
    for k, v := range c.Params.Values {

        out += k;

        if v != nil && len(v) > 1{
            out += "="
            for l := range v {
                out += v[l];
            }
        }
        break;
    }

    fmt.Println(out)
    json.Unmarshal([]byte(out), &post)

    post.Time = int64(time.Now().Unix())
    post.SavePost()

    return c.RenderJson(true);
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
    if c.Session["logged_in"] == "false"{
        return c.Redirect("/")
    }
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
    var i = models.GetGlobalID();

    latest := []models.BlogPost{
        models.GetPost(i - 2),
        models.GetPost(i - 1),
        models.GetLatestPost()};


    c.RenderArgs["latest"] = latest
    c.RenderArgs["logged_in"] = c.Session["logged_in"]
	return c.Render()
}
