package ui
import(
  "fmt"
)
func PrintWelcome() {
  fmt.Printf("Select below options\n1) List Tasks\n2) Add Task\n3) Update Task\n4) Delete Task\n5) Exit\n")
}

func PrintUpdateOptions() {
  fmt.Printf("Select below options\n1) Update Name\n2) Update Task\n3) Update Status\n")
}