package main

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"io/ioutil"
	"net/http/httptest"
	"github.com/abiosoft/river"
)

const (
	samplePayload  = `{"alias":"go-dms-workshop","desc":"Create app and try it with different DMS", "type":"important", "ts":1473837996,"tags":["Golang","Workshop","DMS"],"etime":"4h","rtime":"8h","reminders":["3h", "15m"]}`
	samplePayload2 = `{"alias":"go-dms-workshop3","desc":"Create app and try it with different DMS2", "type":"important", "ts":1473837999,"tags":["Golang","Workshop1","DMS"],"etime":"2h","rtime":"4h","reminders":["1h", "35m"]}`
)

var counter int64

// define dbDriver mockup
type mockDB struct{}

func (db *mockDB) Create(t Task) error {
	return nil
}
func (db *mockDB) ReadById(id *int64) (TaskList, error) {
	fmt.Println("mockDB.ReadById")
	return nil, errors.New("no such id")
}
func (db *mockDB) ReadByAlias(alias *string) (TaskList, error) {
	fmt.Println("mockDB.ReadByAlias")
	task := mockTask(Task{Alias: "go-dms-workshop"}) //,Tags:["Golang", "Workshop", "DMS"]
	return TaskList{task}, nil
}
func (db *mockDB) Update(t Task) error {
	return nil
}
func (db *mockDB) Delete(t Task) error {
	return errors.New("Delete not supported")
}
func (db *mockDB) ReadAll() (TaskList, error) {
	tasks := TaskList{mockTask(Task{Alias: "task1"}), mockTask(Task{Alias: "task2"})}
	return tasks, nil
}

func mockTask(t Task) Task {
	counter++
	task := t //Task{Id:counter}
	//if t.Alias =
	task.Id = counter
	return task
}

func mockEmptyTask(t Task) Task {
	task := Task{Id: 1, Alias: "task1"}
	return task
}

/*
func TestGetTODORecord(t *testing.T) {
    //create mockup db object
    mockDB := &mockDB{}

    rv := river.New()
    TODOHandler := river.NewEndpoint().Get("/:id", getTODORecordExt(TODOModel(mockDB)))
    TODOHandler.Register(TODOModel(mockDB))
    rv.Handle("/todo", TODOHandler)


    server := httptest.NewServer(rv)
    defer server.Close()

    //fmt.Println("url=",server.URL)
    resp, err := http.Get(fmt.Sprintf("%s/todo/1", server.URL))
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

	 resp, err = http.Get(fmt.Sprintf("%s/todo/go-dms-workshop", server.URL))
    if err != nil {
        t.Fatal(err)
    }
    if resp.StatusCode != 200 {
        t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
    }
    //expected := fmt.Sprintf("Visitor count: %d.", i)
    actual, err = ioutil.ReadAll(resp.Body)
    fmt.Println("got result:", string(actual))
    if err != nil {
        t.Fatal(err)
    }

}
/* */
func TestGetTODOListViaServer(t *testing.T) {
	//create mockup db object
	mockDB := &mockDB{}

	rv := river.New()
	TODOHandler := river.NewEndpoint().Get("/", getTODOListExt(TODOModel(mockDB)))
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
	actual, err := ioutil.ReadAll(resp.Body)
	fmt.Println("got result:", string(actual))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTODOListExt(t *testing.T) {
	mockDB := &mockDB{}
	homeHandle := getTODOListExt(TODOModel(mockDB))
	req, _ := http.NewRequest("GET", "/todo", nil)
	w := httptest.NewRecorder()
	homeHandle.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
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
	//expected := fmt.Sprintf("count: %d.", i)
	actual, err := ioutil.ReadAll(resp.Body)
	fmt.Println("got result:", string(actual))
	if err != nil {
		t.Fatal(err)
	}
}
func BenchmarkGetTODOList(b *testing.B) {
	mockDB := &mockDB{}

	rv := river.New()
	TODOHandler := river.NewEndpoint().Get("/", getTODOList)
	TODOHandler.Register(TODOModel(mockDB))
	rv.Handle("/todo", TODOHandler)

	server := httptest.NewServer(rv)
	defer server.Close()

	// run the  function b.N times
	for n := 0; n < b.N; n++ {
		requestGetTODOList(server.URL)
	}
}

//func requestGetTODOList(url string) {
    /* * /
	resp, err := http.Get(fmt.Sprintf("%s/todo", url))


	   if err != nil {
	       //t.Fatal(err)
	   }
	   if resp.StatusCode != 200 {
	       //t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	   }

	   actual, err := ioutil.ReadAll(resp.Body)
	   fmt.Println("got result:", string(actual))
	   if err != nil {
	       //t.Fatal(err)
	   }
	/* */
//}

func BenchmarkGetTODOListExt(b *testing.B) {
	mockDB := &mockDB{}
	homeHandle := getTODOListExt(TODOModel(mockDB))
	req, _ := http.NewRequest("GET", "/todo", nil)
	w := httptest.NewRecorder()
	b.ResetTimer()
	// run the  function b.N times
	for n := 0; n < b.N; n++ {
		homeHandle.ServeHTTP(w, req)
	}
}

func requestGetTODOListX() {
	mockDB := &mockDB{}
	homeHandle := getTODOListExt(TODOModel(mockDB))
	req, _ := http.NewRequest("GET", "/todo", nil)
	w := httptest.NewRecorder()
	homeHandle.ServeHTTP(w, req)
}

