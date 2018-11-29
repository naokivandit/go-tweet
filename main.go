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
	// Created  string
	// Updated  string
}

type Tweet struct {
	Id       int
	Content  string
	Name     string
	Image    string
	Created  string
	Tweet_id int
	// Updated  string
}

type Page struct {
	Title string
	Body  string
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

	rows, err := db.Query("select id, content, name, image, tweet_id from tweets left outer join users on tweets.user_id = users.id order by tweets.created desc")
	if err != nil {
		panic(err.Error())
	}

	tweets := make([]*Tweet, 10)

	var a = 0
	for rows.Next() {
		tw := Tweet{}
		err = rows.Scan(&tw.Id, &tw.Content, &tw.Name, &tw.Image, &tw.Tweet_id)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(tw.Id, tw.Name, tw.Content, tw.Image, tw.Tweet_id)
		tweets[a] = &Tweet{
			tw.Id,
			tw.Content,
			tw.Name,
			tw.Image,
			tw.Content,
			tw.Tweet_id,
		}
		a += 1
	}
	t.Execute(w, tweets)
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

	db, err = sql.Open("mysql", "gouser:gopass@/twitter")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	get_tweet_id := r.URL.RawQuery[9:]
	fmt.Println(get_tweet_id)
	tw := Tweet{}
	err := db.QueryRow(`
		select tweet_id, content
		from tweets
		where tweet_id = ?
		`, get_tweet_id).Scan(&tw.Id, &tw.Content)
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, tw)
}

func posts_show(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/posts_show.html", "templates/header.html", "templates/menu.html")
	db, err = sql.Open("mysql", "gouser:gopass@/twitter")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	get_tweet_id := r.URL.RawQuery[9:]
	tw := Tweet{}
	err := db.QueryRow(`
		select tweet_id, content, name, image, tweets.created from tweets left outer join users on tweets.user_id = users.id where tweet_id = ?
		`, get_tweet_id).Scan(&tw.Id, &tw.Content, &tw.Name, &tw.Image, &tw.Created)
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, tw)

}

func user(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/users_index.html", "templates/header.html", "templates/menu.html", "templates/footer.html")

	db, err = sql.Open("mysql", "gouser:gopass@/twitter")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select id, name, email, password, image from users")
	if err != nil {
		panic(err.Error())
	}

	users := make([]*User, 7)
	var a = 0
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Image)
		if err != nil {
			panic(err.Error())
		}

		// fmt.Println(u.Id, u.Name, u.Email, u.Password, u.Image)
		users[a] = &User{
			u.Id,
			u.Name,
			u.Email,
			u.Password,
			u.Image,
			// u.Created,
			// u.Updated,
		}
		a += 1
	}
	t.Execute(w, users)
}

func user_edit(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/users_edit.html", "templates/header.html", "templates/menu.html")
	t.Execute(w, nil)

	db, err = sql.Open("mysql", "gouser:gopass@/twitter")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	get_user_id := r.URL.RawQuery[8:]
	fmt.Println(get_user_id)
	u := User{}
	err := db.QueryRow(`
		select id, name, email, password, image
		from users
		where id = ?
		`, get_user_id).Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Image)
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, u)

	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	// image := r.FormValue("image")
	fmt.Println(name, email, password, id)
	// Update文発行
	_, err = db.Exec("UPDATE users SET name = ? ,email = ?, password = ? WHERE id = ?", name, email, password, id)
	if err != nil {
		panic(err.Error())
	}

}

func user_show(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/users_show.html", "templates/header.html", "templates/menu.html")

	db, err = sql.Open("mysql", "gouser:gopass@/twitter")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	get_user_id := r.URL.RawQuery[8:]
	fmt.Println(get_user_id)
	u := User{}
	err := db.QueryRow(`
		select id, name, email, image
		from users
		where id = ?
		`, get_user_id).Scan(&u.Id, &u.Name, &u.Email, &u.Image)
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, u)
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

	http.ListenAndServe(":8000", nil)
}
