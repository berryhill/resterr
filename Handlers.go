package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "io/ioutil"
   "io"
   "github.com/gorilla/mux"
   "strconv"
//   "decoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintln(w, "GO AWAY!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   if err := json.NewEncoder(w).Encode(todos); err != nil {
      panic(err)
   }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   todoId := vars["todoId"]

   todoIdInt, _ := strconv.ParseInt(todoId, 10, 64)
   todoStruct := RepoFindTodo(int(todoIdInt))

   fmt.Fprintln(w, "Todo show:", todoId)
   fmt.Fprintln(w, "Struct from DB:", todoStruct)
   fmt.Fprintln(w, "Name:", todoStruct.Name)
   fmt.Fprintln(w, "Finished?:", todoStruct.Completed)
   fmt.Fprintln(w, "Time:", todoStruct.Due)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
   var todo Todo
   body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
   if err != nil {
      panic(err)
   }
   if err := r.Body.Close(); err != nil {
      panic(err)
   }
   if err := json.Unmarshal(body, &todo); err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(422) // unprocessable entity
      if err := json.NewEncoder(w).Encode(err); err != nil {
         panic(err)
      }
   }

   t := RepoCreateTodo(todo)
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   w.WriteHeader(http.StatusCreated)
   if err := json.NewEncoder(w).Encode(t); err != nil {
      panic(err)
   }
}
