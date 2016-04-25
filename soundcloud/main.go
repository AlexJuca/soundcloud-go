package main

import (
    "github.com/AlexJuca/soundcloud-go"
    "fmt")

func main() {
    client := soundclient.SoundCloud{ClientId: "YOUR_CLIENT_ID",
                                     ClientSecret: "YOUR_SECRET"}
                             
    // Retrieve a track by id as json object
    song := client.Tracks("13158")
    
    // Get track title
    title, _ := song.GetString("title")
    // Get track desccription
    description, err := song.GetString("description")
    if err == nil {
        songOwner, _ := song.GetString("id")
         fmt.Println("songOwner id ->", songOwner)
    }
    // Print track information
    fmt.Println("Title ->", title)
    fmt.Println("Description ->", description)
    
    // Retrieve a User by id as Json object
    user := client.Users("36991")
    
    // Get the users full name and country
    fname, _ := user.GetString("full_name")
    country, _ := user.GetString("country")
    
    // Print user information
    fmt.Println("Full Name ->", fname)
    fmt.Println("Country ->", country)
    
    // Retrieve a comment for a specific track
    comment := client.Comments("13158")
    commentBody, _ := comment.GetString("body")
    
    fmt.Println("Comment -> ", commentBody)
    
}