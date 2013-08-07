package models

type BlogPost struct {
    Content string
    Title string
    Time int64
}

func GetLatestPost() BlogPost {
    return BlogPost{"Hello! This is a test!", "Some title", 1993}
}