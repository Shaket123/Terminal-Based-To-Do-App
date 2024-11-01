package handlers

import(
  "strconv"
  "fmt"
)

func validateInput(inp string,flag *int,caller string) (err_found bool){
  inp_int,err := strconv.Atoi(inp)
  if(err != nil){
    err_found = true
  }
  if(caller == "mainscreen"){
    if(inp_int<1 || inp_int >5){
      err_found = true
    }
  }

  if(caller == "updatescreen"){
    if(inp_int<1 || inp_int >3){
      err_found = true
    }
  }
  
  if(err_found){
    fmt.Println("Invalid Input",inp)
  }
  *flag = inp_int
  return
}