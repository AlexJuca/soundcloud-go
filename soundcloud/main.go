package main

import (
    "github.com/AlexJuca/soundcloud-go"
    "fmt")

func main() {
    client := soundclient.SoundCloud{ClientId: "042bd096125d480ea6d7f08cf4b092eb",
                                     ClientSecret: "042bd096125d480ea6d7f08cf4b092eb"}
                             
    // Retrieve a track by id as json object
    song := client.Tracks("13158")
    
    // Get track title
    title, _ := song.GetString("title")
    // Get track desccription
    description, _ := song.GetString("description")
    
    // Print track information
    fmt.Println("Title ->", title)
    fmt.Println("Description ->", description)
    
    // Retrieve a User by id as Json object
    user := client.Users("1")
    
    // Get the users full name and country
    fname, _ := user.GetString("full_name")
    country, _ := user.GetString("country")
    
    // Print user information
    fmt.Println("Full Name ->", fname)
    fmt.Println("COuntry ->", country)
    
}