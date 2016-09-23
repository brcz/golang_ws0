package main

import (
      "testing"
        "fmt"
        
		  "net/http"
        "io/ioutil"
        "net/http/httptest"
        "github.com/abiosoft/river"
        
)

// define dbDriver mockup
type mockDB struct {}
func (db *mockDB) Create(t Task) error {
return nil
}
func (db *mockDB) ReadById(id *int64) (TaskList, error) {
return nil, nil
}
func (db *mockDB) ReadByAlias(alias *string) (TaskList, error) {
    return nil, nil
}
func (db *mockDB) Update(t Task) error {
return nil
}
func (db *mockDB) Delete(t Task) error {
return nil
}
func (db *mockDB) ReadAll() (TaskList, error) {
    tasks := TaskList{Task{Alias:"task1"}, Task{Alias:"task2"}}
    return tasks, nil
}

type myCo struct {
    *river.Context
}
func (mc myCo) Init(responser http.ResponseWriter) {
    fmt.Printf("content: %v",mc)
    //mc.rw = responser
}

func TestGetTODOList(t *testing.T) {
    /* * /
	req, err := http.NewRequest("GET", "/todo", nil)
    if err != nil {
        t.Fatal(err)
    }
    /* */

    //rec := httptest.NewRecorder()
    
    //create mockup db object
    mockDB := &mockDB{}

    rv := river.New()
    TODOHandler := river.NewEndpoint().Get("/", getTODOList)
    TODOHandler.Register(TODOModel(mockDB))
    rv.Handle("/todo", TODOHandler)
    
    
    server := httptest.NewServer(rv)
    defer server.Close()
    
    //fmt.Println("url=",server.URL)
    resp, err := http.Get(fmt.Sprintf("%s/todo", server.URL))
    if err != nil {
        t.Fatal(err)
    }
    if resp.StatusCode != 200 {
        t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
    }
    //expected := fmt.Sprintf("Visitor count: %d.", i)
    actual, err := ioutil.ReadAll(resp.Body)
    fmt.Println("got result:", string(actual))
    if err != nil {
        t.Fatal(err)
    }
    
}
