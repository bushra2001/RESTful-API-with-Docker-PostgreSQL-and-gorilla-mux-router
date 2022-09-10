package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
    "strconv"
    // "net/url"
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
var Articles = []Article{
    Article{Id: "a",Title: "variable a", Content: 0},
    Article{Id: "b",Title: "variable b", Content: 0},
    Article{Id: "sum",Title: "sum of a and b", Content: 0},
}
func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()


}
