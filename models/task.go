package models

import (
  "time"
)

type Data struct {
  Id int                   `json:"id"`
  Name string              `json:"name"`
  Task string              `json:"task"`
  Date time.Time           `json:"date"`
  Completed bool           `json:"completed"`
  Createdon time.Time      `json:"createdon"`
  Completedon time.Time    `json:"completedon"`
}