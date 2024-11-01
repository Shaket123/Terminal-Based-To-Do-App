package handlers

import (
  "To-Do-App/models"
  "To-Do-App/storage"
  "To-Do-App/ui"
  "fmt"
  "time"
  "bufio"
  "os"
  "math/rand"
)

var maindata []models.Data

func Run() {
  var inp string
  var flag int
  maindata = storage.LoadData()
  for {
    ui.PrintWelcome()
    for{
      fmt.Scanln(&inp)
      if !validateInput(inp,&flag,"mainscreen"){
        break
      }
    }
    switch(flag){
      case 1:list()
      case 2:add()
      case 3:update()
      case 4:delete()
      case 5:exit()
    }
  }
}


func add(){
  var name string
  fmt.Println("Enter the Task Name to add")
  bf := bufio.NewReader(os.Stdin)
  name,err := bf.ReadString('\n')
  if(err != nil){
    panic(err)
  }
  var d models.Data
  d.Id = genId()
  d.Name = name
  d.Createdon = time.Now()
  maindata = append(maindata, d)
  fmt.Println("Task added successfully")
}

func update(){
  var Id int
  fmt.Println("Enter the Task Id to delete")
  fmt.Scanln(&Id)
  ind := findbyID(maindata,Id)
  if(ind == -1){
    fmt.Println("Task not found")
  }
  ui.PrintUpdateOptions()
  var flag1 int
  var updateopt string
  for{
    fmt.Scanln(&updateopt)
    if !validateInput(updateopt,&flag1,"updatescreen"){
      break
    }
  }

  switch(flag1){
    case 1: bf := bufio.NewReader(os.Stdin)
            name,err := bf.ReadString('\n')
            if(err != nil){
              panic(err)
            }
              maindata[ind].Name = name
    case 2: bf := bufio.NewReader(os.Stdin)
            name,err := bf.ReadString('\n')
            if(err != nil){
              panic(err)
            }
              maindata[ind].Task = name
    case 3:maindata[ind].Completed = !maindata[ind].Completed
    default :fmt.Println("Invalid option")
  }

  fmt.Println("Task updated successfully")
  //update name
  //update task
  //update status
}

func delete(){
  var Id int
  fmt.Println("Enter the Task Id to delete")
  fmt.Scanln(&Id)
  fmt.Println("Task deleted successfully")
  ind := findbyID(maindata,Id)
  if(ind == -1){
    fmt.Println("Task not found")
  }else{
    maindata = append(maindata[:ind],maindata[ind+1:]...)
  }
}

func list(){
  fmt.Println("List of Tasks")
  fmt.Println(maindata)
}

func exit(){
  storage.SaveData(maindata)
  fmt.Println("Thanks for using the application, Exiting !!!")
  os.Exit(0)
}

func findbyID(md []models.Data,target int) int{
  for i,ele:= range md{
    if  ele.Id== target{
      return i
    }
  }
  return -1
}

func genId() int{
  min := 1000
  max := 9999
  return rand.Intn(max-min+1) + min
}