// DOES NOT COMPILE!

/*
You create two handler functions, setMessage and showMessage, 
and attach them to /set_message and /show_message, respectively. 
Let’s start with setMessage, which is straightforward.
*/

package main

import (
    "encoding/base64"
    "fmt"
    "net/http"
    "time"
)

// Setting message
func setMessage(w http.ResponseWriter, r *http.Request) { 
	msg := []byte("Hello World!")
	c := http.Cookie{
        Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg), 
	}
	http.SetCookie(w, &c) 
}


/*
This isn’t much different from the setCookie handler function from earlier, 
except this time you do a Base64 URL encoding of the message. 
You do so because the cookie values need to be URL encoded in the header. 
You managed to get away with it earlier because you didn’t have any special characters 
like a space or the percentage sign, but you can’t get away with here because 
messages will eventually need to have them.

Now let’s look at the showMessage function:
*/


func showMessage(w http.ResponseWriter, r *http.Request) { 
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie { 
			fmt.Fprintln(w, "No message found")
		} else {
			rc := http.Cookie{
				Name: "flash", 
				MaxAge: -1,
				Expires: time.Unix(1, 0),
			}

		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value) 
		fmt.Fprintln(w, string(val))		
	}
}


func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_message", setMessage) 
	http.HandleFunc("/show_message", showMessage) 

	server.ListenAndServe()
}

/*
First, you get the cookie. If you can’t find the cookie (err will have a value of http.ErrNoCookie), 
you’ll show the message “No message found.”
If you find the message, you have to do two things:
1 Create a cookie with the same name, but with MaxAge set to a negative number and an Expires value that’s in the past.
2 Send the cookie to the browser with SetCookie.
Here you’re replacing the existing cookie, essentially removing it altogether because 
the MaxAge field is a negative number and the Expires field is in the past. 
Once you do that, you can decode your string and show the value.

Go back to the Web Inspector and look at the cookies. 
Your cookie is gone! Setting a cookie with the same name to the browser will replace 
the old cookie with the new cookie of the same name. 
Because the new cookie has a negative number for MaxAge and expires in some time in the past, 
this tells the browser to remove the cookie, which means the earlier cookie you set is removed.
This is what you’ll see in the browser:
        No message found
*/