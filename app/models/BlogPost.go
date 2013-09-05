package models

import (
    "menteslibres.net/gosexy/redis"
    "strconv"
)

type BlogPost struct {
    Content string
    Title string
    Time int64
    Id int64
}

var client *redis.Client

/**
 * Save a post
 */
func (post *BlogPost) SavePost() {
    
    //Get current id from redis
    var id string
    client.Incr("global:id")
    id, err := client.Get("global:id")

    if err != nil {
        id = "1"
        client.Set("global:id", 0)
    }

    //convert id to key
    key := "post:" + id

    if post.Time != 0 {
        //Store post with key
        client.HMSet(key, "title", post.Title, "content", post.Content, "time", post.Time)
    } else {
        post.Time = -1
        client.HMSet(key, "title", post.Title, "content", post.Content, "time", post.Time)
    }    
}

/**
 * Retieve the last recorded post at 2 strings, title & content
 */
func GetLatestPost() (BlogPost){
    return GetPost(GetGlobalID());
}

/**
 * Get post for a given post_id
 */
func GetPost(global_id int64) (BlogPost){
    post := BlogPost{}

    //Construct Key
    post.Id = global_id
    key := "post:" + strconv.FormatInt(global_id, 10)

    title, err := client.HMGet(key, "title")

    if err != nil {
        return post;
    }

    content, _ := client.HMGet(key, "content")
    time, _ := client.HMGet(key, "time")

    //Fetch data
    post.Title = title[0]
    post.Content = content[0]
    post.Time, _ = strconv.ParseInt(time[0], 10, 64)
    post.Id = global_id
    return post
}

/**
 * get the current value of global:id
 */
func GetGlobalID() (int64) {
    id, _ := client.Get("global:id")
    int_id, _ := strconv.ParseInt(id, 10, 64)
    return int_id
}

/**
 * Get connection to redis and retun client for a given host & port
 */
func GetConnection(host string, port uint) (*redis.Client, error) {
    var err error

    client = redis.New()

    err = client.Connect(host, port)

    return client, err
}

func Connect(){
    client, _ = GetConnection("127.0.0.1", uint(6379))
}