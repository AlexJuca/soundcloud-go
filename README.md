# SoundCloud Go API Wrapper

## Description
This is the SoundCloud API Wrapper written in Go. It provides methods to handle
authorization and execution of HTTP Calls. I built this after my 5th day learning Go so there might be lots of silly 
beginner mistakes, if you do find one please share some of your wisdom with me. :)


# Examples


```go
    # Register a client with your YOUR_CLIENT_ID as ClientId
    client := soundclient.SoundCloud{ ClientId: YOUR_CLIENT_ID,
                                      ClientSecret: YOUR_SECRET_KEY}
    # Get a track 
    song := client.Tracks("13158")
    
    # Get track title
    title, _ := song.GetString("title")
    # Get track desccription
    description, _ := song.GetString("description")
    
    # Print track information
    fmt.Println("Title ->", title)
    fmt.Println("Description ->", description)
```
    
    
## Developed by 
* Alexandre A.Juca <alexandre.juca@bitfyr.com>

