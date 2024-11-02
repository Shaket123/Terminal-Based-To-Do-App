package utility

import(
  "strconv"
  "fmt"
)

func ValidateInput(inp string,flag *int,caller string) (err_found bool){
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

  if(caller == "idscreen"){
    if(inp_int<1000 || inp_int >9999){
      err_found = true
    }
  }

  if(err_found){
    if(caller == "idscreen"){
      fmt.Println("Invalid ID",inp)
    }else{
      fmt.Println("Invalid Input",inp)
    }
    
  }
  *flag = inp_int
  return
}