package handlers

import (
  "To-Do-App/models"
  "To-Do-App/storage"
  "To-Do-App/ui"
  "time"
  "os"
  "math/rand"
)

var Maindata []models.Data

func Run() {
  Maindata = storage.LoadData()
  ui.PrintHeader()
  for {
    ui.PrintWelcome()
    flag := ui.PrintUserInpInt("mainscreen")
    switch(flag){
      case 1:list()
      case 2:add()
      case 3:update()
      case 4:delete()
      case 5:Exit()
    }
  }
}


func add(){
  name := ui.PrintUserInpString("Task Name","Task added Successfully")
  var d models.Data
  d.Id = genId()
  d.Name = name
  d.Createdon = time.Now()
  Maindata = append(Maindata, d)
}

func update(){
  var Id,ind int
  ind = -1
  ui.PrintWithouFormat("Enter the Task Id to Update")
  for{
    Id = ui.PrintUserInpInt("idscreen")
    ind = findbyID(Maindata,Id)
    if(ind == -1){
      ui.PrintError("Task not found")
    }else{
      break
    }
  }
  ui.PrintUpdateOptions()
  flag := ui.PrintUserInpInt("updatescreen")

  switch(flag){
    case 1: name := ui.PrintUserInpString("Enter the Task Name to Update","Task Name Updated Successfully")
            Maindata[ind].Name = name
    case 2: name := ui.PrintUserInpString("Enter the Task Description to Update","Task Description Updated Successfully")
            Maindata[ind].Task = name
    case 3:Maindata[ind].Completed = !Maindata[ind].Completed
    ui.PrintMsg("Task Compiliton Status Updated Successfully")
  }
}

func delete(){
  var Id int
  var ind int = -1
  ui.PrintWithouFormat("Enter the Task Id to delete")

  for{
    Id = ui.PrintUserInpInt("idscreen")
    ind = findbyID(Maindata,Id)

    if(ind == -1){
      ui.PrintError("Task not found")
    }else{
      break
    }
  }
    Maindata = append(Maindata[:ind],Maindata[ind+1:]...)
  ui.PrintMsg("Task deleted successfully")

}

func list(){
  ui.PrintList(Maindata)
  ui.PrintMsg("")
}

func Exit(){
  ui.PrintWithouFormat("")
  storage.SaveData(Maindata)
  ui.PrintExit()
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