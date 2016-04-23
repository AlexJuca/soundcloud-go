# SoundCloud Go API Wrapper

## Description
This is the SoundCloud API Wrapper written in Go. It provides methods to handle
authorization and execution of HTTP Calls.

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
    desccription, _ := song.GetString("description")
    
    # Print track information
    fmt.Println("Title ->", title)
    fmt.Println("Description ->", desccription)
    
    
## Developed by 
* Alexandre A.Juca <alexandre.juca@bitfyr.com>

## License

MIT © [Alexandre A.Juca]<alexandre.juca@bitfyr.com>

Copyright (c) 2015-2016 Bitfyr, Inc. http://www.bitfyr.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

