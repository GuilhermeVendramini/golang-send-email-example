package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	gomail "gopkg.in/gomail.v2"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", Contact)
	mux.POST("/contact/process", Process)
	http.ListenAndServe(":8080", mux)
}

// Contact form
func Contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	html := `<h1>Contact</h1>
		<form method="post" action="/contact/process" role="form">
		<div class="form-group">
			<input type="text" class="form-control" id="name" name="name" placeholder="Name" required>
		</div>
		<div class="form-group">
			<input type="text" class="form-control" id="email" name="email" placeholder="Email" required>
		</div>
		<div class="form-group">
			<input type="text" class="form-control" id="subject" name="subject" placeholder="Subject" required>
		</div>
		<div class="form-group">
			<textarea class="form-control" type="textarea" id="message" name="message" placeholder="Message" maxlength="140" rows="7"></textarea>
		</div>
		<div class="form-group">
			<button  type="submit" class="btn btn-primary">Submit</button>
		</div> 
	</form>`

	w.Write([]byte(fmt.Sprintf(html)))
}

// Process Email
func Process(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	subj := r.FormValue("subject")
	message := r.FormValue("message")

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", "to-test1@test.com", "to-test2@test.com")
	m.SetAddressHeader("Cc", "copy-test@test.com", "Name")
	m.SetHeader("Subject", subj)
	m.SetBody("text/html", "<b>Name:"+name+"</b><br>"+message)
	// m.Attach("/home/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 465, "your-gmail@gmail.com", "your-password")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
