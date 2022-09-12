//Created by Bushra Ashfaque
package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
    "strconv"
    "database/sql"
    _ "github.com/lib/pq"
    // "reflect"
)


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range Articles {
        
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
    

}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    

    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func gVariable(w http.ResponseWriter, r *http.Request) {    
    fmt.Println("Endpoint hit read")
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    // fmt.Println(vars)
    key := vars["id"]
    // fmt.Println(key)
    for _, article := range Articles {
        if (article.Id == key){
            if (article.Id == "sum") {
                sum_value := Articles[0].Content + Articles[1].Content 
                response :="Sum of a and b is " + strconv.Itoa(sum_value)
                err := json.NewEncoder(w).Encode(&response)
                fmt.Println(response)
                insertDynStmt:=`UPDATE "Articles" SET "Content" = $1 WHERE "Id" = $2`
                var _,e = db.Exec(insertDynStmt,sum_value,article.Id)
                CheckError(e)
                if err != nil {
                    log.Fatalln("There was an error encoding the initialized struct")
                }
            } else {
                val_read :="Value of "+ article.Id + " is " + strconv.Itoa(article.Content)
                error := json.NewEncoder(w).Encode(&article.Content)
                fmt.Println(val_read)
                if error != nil {
                    log.Fatalln("There was an error encoding the initialized struct")
                }
                }        
        }
    }

}

func wVariable(w http.ResponseWriter, r *http.Request) {    
    fmt.Println("Endpoint hit write")
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    // fmt.Println(vars)
    id_key := vars["id"]
    // fmt.Println(id_key)
    content := vars["content"]
    content_key,_ := strconv.Atoi(content)
    // fmt.Println(content_key)
    var i = 0
    // fmt.Println(key)
    for _, article := range Articles {
        if article.Id == id_key {
            
            Articles[i].Content = content_key
            response :="Updated value of "+ article.Id + " is " + strconv.Itoa(Articles[i].Content)
            fmt.Println(response)
            err := json.NewEncoder(w).Encode(&Articles[i].Content)
            insertDynStmt:=`UPDATE "Articles" SET "Content" = $1 WHERE "Id" = $2`
            var _,e = db.Exec(insertDynStmt,Articles[i].Content,article.Id)
            CheckError(e)
            if err != nil {
                log.Fatalln("There was an error encoding the initialized struct")
        }   
        }
        i+=1
    }
}

func dVariable(w http.ResponseWriter, r *http.Request) {    
    fmt.Println("Endpoint hit delete")
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    // fmt.Println(vars)
    id_key := vars["id"]
    var i = 0
    // fmt.Println(key)
    for _, article := range Articles {
        if article.Id == id_key {
            
            Articles[i].Content = 0
            response :="Updated value of "+ article.Id + " is " + strconv.Itoa(Articles[i].Content)
            fmt.Println(response)
            err := json.NewEncoder(w).Encode(&Articles[i].Content)
            insertDynStmt:=`UPDATE "Articles" SET "Content" = $1 WHERE "Id" = $2`
            var _,e = db.Exec(insertDynStmt,Articles[i].Content,article.Id)
            CheckError(e)
            if err != nil {
                log.Fatalln("There was an error encoding the initialized struct")
        }   
        }
        i+=1
    }
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    // myRouter.HandleFunc("/", gVariable("a"))
    myRouter.HandleFunc("/articles", returnAllArticles)
    // NOTE: Ordering is important here! This has to be defined before
    // the other `/article` endpoint. 
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    myRouter.HandleFunc("/read{id}", gVariable).Methods("GET")
    myRouter.HandleFunc("/delete{id}", dVariable).Methods("DELETE")
    myRouter.HandleFunc("/update{id}{content}", wVariable).Methods("POST")
    log.Fatal(http.ListenAndServe(":10000", myRouter))


}



type Article struct {
    Id string `json:"Id"`
    Title string `json:"Title"`
    Content int `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database

// var Articles = []Article{
//     Article{Id: "a",Title: "variable a", Content: 0},
//     Article{Id: "b",Title: "variable b", Content: 0},
//     Article{Id: "sum",Title: "sum of a and b", Content: 0},
// }

const (
    host = "127.0.0.1"
    port = 5432
    user = "postgres"
    password = "new_password"
    dbname = "RestApi_db"
)
var psqlconn = fmt.Sprintf("host= %s port=%d user= %s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)
var db,e = sql.Open("postgres", psqlconn)
var readDb = `Select * from "Articles"`
var row, err = db.Exec(readDb)
var Articles []Article 
// CheckError(err)
func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")

    CheckError(err)
    defer db.Close()
 
    get_val := `Select "Title", "Content" FROM "Articles" WHERE "Id"=$1;`
    arr := [3]string{"a","b","sum"}
    
    for j:= 0; j < 3; j++{
    // fmt.P
    rows,error := db.Query(get_val, arr[j])
    // rows, err := db.Query(get_val,"a")
    CheckError(error)
    // fmt.Println(reflect.TypeOf(rows))
    if rows == nil {
        // panic(err)
        Articles = []Article{
        Article{Id: "a",Title: "variable a", Content: 0},
        Article{Id: "b",Title: "variable b", Content: 0},
        Article{Id: "sum",Title: "sum of a and b", Content: 0},
        }
    }

    
    defer rows.Close()
    for rows.Next() {
    var r_Title string
    var r_Content int
 
    err = rows.Scan(&r_Title, &r_Content)
    CheckError(err)
    CheckError(error)

    
    
    // fmt.Println(r_Title, r_Content)
    Articles = append(Articles,Article{Id: arr[j],Title: r_Title , Content: r_Content})
    }

    }
// fmt.Println(Articles)

 
            
    insertStmt := `insert into "Articles"("Id", "Title", "Content") values('a', 'variable a', 0),('b', 'variable b', 0), ('sum', 'sum of a and b', 0) ON CONFLICT DO NOTHING`
    _, e := db.Exec(insertStmt)
    CheckError(e)
    handleRequests()
    
}
func CheckError(err error){
        if err != nil {
            panic(err)
        }
}

