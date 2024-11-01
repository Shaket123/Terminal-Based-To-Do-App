package storage

import (
	"To-Do-App/models"
	"encoding/json"
	"os"
    "fmt"
)

func LoadData() []models.Data{
  file, err := os.OpenFile("db.json",os.O_CREATE|os.O_RDONLY, 0644)
  if err != nil {
      panic(err)
  }
stat,err1 := file.Stat()
if err1 != nil{
  panic(err1)
}
if stat.Size() == 0{
    return []models.Data{}
}


var data []models.Data
decode := json.NewDecoder(file)
err2 := decode.Decode(&data)
if err2 != nil {
    fmt.Println(stat)
    panic(err2)
}
    file.Close()
return data
    
}

func SaveData(data []models.Data){
  file, err := os.OpenFile("db.json",os.O_WRONLY|os.O_TRUNC,0644)
  if err != nil {
      panic(err)
  }
  defer file.Close()

  encode := json.NewEncoder(file)
  err = encode.Encode(data)
  if err != nil {
      panic(err)
  }
}