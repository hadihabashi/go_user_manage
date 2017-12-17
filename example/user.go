package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"

	"github.com/hadihabashi/go_user_manage"
	"gopkg.in/gomail.v2"

)

func mail(to string , msg string , subject string)error{

	m := gomail.NewMessage()
	m.SetHeader("From", "monitoring.ulin@gmail.com")
	m.SetHeader("To", to)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "monitoring.ulin@gmail.com", "#adi1369")

	// Send the email to Bob, Cora and Dan.
	err := d.DialAndSend(m)
	if err != nil{
		return err
	}else {
		return nil
	}

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandString() string {
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func register(w http.ResponseWriter, req *http.Request, userstate go_user_manage.IUserState) {
	fmt.Println("method:", req.Method) //get request method
	if req.Method == "GET" {
		fmt.Fprintf(w, "Register User Api") // write data to response
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()

		//user := req.Form["user"]
		pass := req.Form["pass"]
		email := req.Form["email"]
		user := email

		if (user != nil) && (pass!= nil) && (email != nil) {
			//Check User Register Earlier
			ch := userstate.HasUser(user[0])
			if ch {
				fmt.Fprintf(w, "User Exist \r\n")
			} else {
				confirmcode, err := userstate.GenerateUniqueConfirmationCode()

				if err != nil {
					fmt.Fprintf(w, "Can Not Create ConfirmationCode \r\n")
				}

				err1 := userstate.AddUser(user[0], pass[0], email[0])
				if (err1 != nil){
					fmt.Fprintf(w, "AddUser Error 1")
					fmt.Println(err1)
				}else {
					err2 := userstate.AddUnconfirmed(user[0], confirmcode)
					if (err1 != nil) {
						fmt.Fprintf(w, "AddUser Error 2")
						fmt.Println(err2)
					}else {

						answer := userstate.HasUser(user[0])

						if answer {
							fmt.Fprintf(w, "Please Check Your Mail For Confiram Link.")
							msg := "Confirmation Links :   \r\n\r\n  http://ge1.ulin.ir:3000/confirm?code=" + confirmcode + "&user=" + user[0] + ".\r\n"
							subject := "Confirm Registeration"
							err := mail(email[0], msg , subject )
							if err != nil {
								log.Fatal(err)
							}
						} else {
							fmt.Fprintf(w, "SomeThings Error To Add User \r\n")
						}
					}
				}
			}

		}
	}
}

func confirm(w http.ResponseWriter, req *http.Request, userstate go_user_manage.IUserState) {

	fmt.Println("method:", req.Method) //get request method
	w.Header().Set("Access-Control-Allow-Origin", "*")
	req.ParseForm()

	code := req.Form["code"]
	user := req.Form["user"]

	if (code != nil) && (user != nil){
		servercode , err := userstate.ConfirmationCode(user[0])
		if err != nil {
			fmt.Fprintf(w, "User Not Found Or Confirmation Code is Wrong")
		}else {
			if (servercode == code[0]){
				userstate.Confirm(user[0])
			}
		}

		fmt.Fprintf(w, "User " + user[0] + " was confirmed: %v\n", userstate.IsConfirmed(user[0]))
	}else {
		fmt.Fprintf(w, "Confirm User Api")
	}


}

func login(w http.ResponseWriter, req *http.Request, userstate go_user_manage.IUserState) {

	fmt.Println("method:", req.Method) //get request method
	if req.Method == "GET" {
		fmt.Fprintf(w, "Login User Api") // write data to response
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()
		user := req.Form["email"]
		pass := req.Form["pass"]

		if (user != nil) && (pass != nil) {
			if !userstate.HasUser(user[0]) {
				fmt.Fprintf(w, "User Not Found")
			}else {
				answer := userstate.CorrectPassword(user[0],pass[0])
				if !answer{
					fmt.Fprintf(w, "Wrong Password")
				}else{
					userstate.Login(w, user[0])
					fmt.Fprintf(w, "OK")
				}
			}
		}
	}

}



func logout(w http.ResponseWriter, req *http.Request, userstate go_user_manage.IUserState) {

	fmt.Println("method:", req.Method) //get request method
	if req.Method == "GET" {
		fmt.Fprintf(w, "LogOut User Api") // write data to response
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()

		user := req.Form["email"]

		if (user != nil) {
			err := userstate.Logout(w, user[0])
			if err != nil{
				fmt.Fprintf(w,"NOT")
			}else {
				fmt.Fprintf(w,"OK")
			}
		}
	}

}

func resetpass(w http.ResponseWriter, req *http.Request, userstate go_user_manage.IUserState) {

	fmt.Println("method:", req.Method) //get request method
	if req.Method == "GET" {
		fmt.Fprintf(w, "Reset Password User Api") // write data to response
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()

		email := req.Form["email"]
		user := email


		if (email != nil) && (user != nil) {
			orginalemail, err := userstate.Email(user[0])
			if (err != nil) {
				fmt.Fprintf(w, "Email Not Found")
			} else {
				if (orginalemail == email[0]) {
					newpass := RandString()
					userstate.SetPassword(user[0], newpass)

					msg := "To: "+ email[0] + "\r\n Subject: New PassWord!\r\n" + "\r\n" +   newpass + ".\r\n"
					subject := "Rest Password"
					err := mail(email[0],msg , subject)
					if err != nil{
						log.Fatal(err)
						fmt.Fprintf(w, "0")
					}else {
						fmt.Fprintf(w, "1")
					}

				}
			}
		}
	}

}

func changepass(w http.ResponseWriter, req *http.Request, userstate go_user_manage.IUserState) {

	fmt.Println("method:", req.Method) //get request method
	if req.Method == "GET" {
		fmt.Fprintf(w, "Reset Password User Api") // write data to response
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()

		user := req.Form["user"]
		pass := req.Form["pass"]
		oldpass := req.Form["oldpass"]

		if (pass != nil) && (oldpass != nil) && (user != nil) {
			if (!userstate.HasUser(user[0])) && (!userstate.IsLoggedIn(user[0])){
				fmt.Fprintf(w, "Username Not Found Or Not Login")
			} else {
				answer := userstate.CorrectPassword(user[0], oldpass[0])
				if (!answer) {
					fmt.Fprintf(w, "Old Pass Is Not Correct")
				} else {
					userstate.SetPassword(user[0], pass[0])
					fmt.Fprintf(w, "1")
				}
			}
		}

	}

}


func makeHandler(fn func(http.ResponseWriter, *http.Request, go_user_manage.IUserState),userstate go_user_manage.IUserState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := []string{"/register","/confirm" ,"/login","/logout","/resetpass","/changepass"}
		m := (r.URL.Path)
		c := 0
		for _,i := range path{
			if m == i {
				c++
			}

		}
		if c == 0 {
			http.NotFound(w, r)
			return
		}else if c == 1 {
			fn(w, r, userstate)
		} else {
			log.Fatal("BADSECTOR")
		}


	}
}

func rootPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UserManagement Api")
}


func main() {

	// New permissions middleware
	perm, err := go_user_manage.New(0, "", "SADFjv$a3f8admkasd")
	if err != nil {
		log.Fatalln(err)
	}

	// Get the userstate, used in the handlers below
	userstate := perm.UserState()

	http.HandleFunc("/", rootPath)



	// Admin Functions


	// User Functions
	http.HandleFunc("/register", makeHandler(register,userstate))
	http.HandleFunc("/confirm", makeHandler(confirm,userstate))
	http.HandleFunc("/login", makeHandler(login,userstate))
	http.HandleFunc("/logout", makeHandler(logout,userstate))
	http.HandleFunc("/resetpass", makeHandler(resetpass,userstate))
	http.HandleFunc("/changepass", makeHandler(changepass,userstate))


	// Serve
	http.ListenAndServe(":3000", nil)

}
