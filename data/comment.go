// リスト5.12用
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Image    string
	Created  string
	Updated  string
}

type Tweet struct {
	Id      int
	Content string
}
type Greeting struct {
	Messege string
	Target  string
}

var tweet_id int
var content string

func top(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/top.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/about.html", "templates/header.html", "templates/menu.html", "templates/footer.html")
	t.Execute(w, nil)
}

var db *sql.DB
var err error

func signup(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/signup.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)

	db, err = sql.Open("mysql", "gouser:gopass@/twitter")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err = db.Exec("INSERT INTO users(name, email, password) VALUES(?, ?, ?)", name, email, password)
	fmt.Println("-------------------")
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(password)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/login.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)

	if r.Method != "POST" {
		http.ServeFile(w, r, "login.html")
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	var databaseEmail string
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM users WHERE email=?", email).Scan(&databaseEmail, &databasePassword)

	fmt.Println(databaseEmail)

	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}

	// r.Write([]byte(databaseEmail))
}

func posts(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/posts.html", "templates/header.html", "templates/menu.html", "templates/tweets.html")

	db, err = sql.Open("mysql", "gouser:gopass@/twitter")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select content from tweets")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{})
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var a = 0

	tweets := make([]*Tweet, 10)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string

		for i, col := range values {

			a += 1

			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
			fmt.Println(a)
			fmt.Println(value)

			// t.ExecuteTemplate(w, "posts")

			tweets[a] = &Tweet{
				a,
				value,
			}
			//
			t.Execute(w, tweets)
		}

		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

}

func posts_new(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/posts_new.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)

	db, err = sql.Open("mysql", "gouser:gopass@/twitter")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	user_id := r.FormValue("user_id")
	content := r.FormValue("content")

	_, err = db.Exec("INSERT INTO tweets(user_id, content) VALUES(?, ?)", user_id, content)
	fmt.Println("-------------------")
	fmt.Println(user_id)
	fmt.Println(content)

}

func posts_edit(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/posts_edit.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)
}

func posts_show(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/posts_show.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)
}

func user(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/users_index.html", "templates/header.html", "templates/menu.html", "templates/footer.html")
	t.Execute(w, nil)
}

func user_edit(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/users_edit.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)

	db, err = sql.Open("mysql", "gouser:gopass@/twitter")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	image := r.FormValue("image")

	fmt.Println("アカウント編集")
	fmt.Println("id:", id)
	fmt.Println("name:", name)
	fmt.Println("email:", email)
	fmt.Println("image:", image)
	fmt.Println("pasword", password)

	// Update文発行
	_, err = db.Exec("UPDATE users SET name = ? ,email = ?, password = ?, image = ? WHERE id = ?", name, email, password, image, id)
	if err != nil {
		panic(err.Error())
	}

}

func user_show(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/users_show.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)
}

func user_show_like(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/users_show_like.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)
}

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.HandleFunc("/", top)
	http.HandleFunc("/about", about)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/posts", posts)
	http.HandleFunc("/posts_new", posts_new)
	http.HandleFunc("/posts_edit", posts_edit)
	http.HandleFunc("/posts_show", posts_show)

	http.HandleFunc("/user", user)
	http.HandleFunc("/user_edit", user_edit)
	http.HandleFunc("/user_show", user_show)
	http.HandleFunc("/user_show_like", user_show_like)

	http.ListenAndServe(":8888", nil)
}
