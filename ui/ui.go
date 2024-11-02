package ui

import (
	"To-Do-App/models"
  "To-Do-App/utility"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"bufio"
)

func PrintHeader() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{"🟆 TO Do App 🟆"})
	table.Append([]string{"🚀 Version 1.0.0 🚀"})
	table.Append([]string{"Created by: Shaket Dahiwale 💜"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
	table.Render()
}
func PrintWelcome() {
	table2 := tablewriter.NewWriter(os.Stdout)
	table2.SetHeader([]string{"Select below options"})
	table2.SetRowLine(true)
	table2.Render()
	table1 := tablewriter.NewWriter(os.Stdout)
	table1.SetRowLine(true)
	table1.Append([]string{"1", "List Tasks"})
	table1.Append([]string{"2", "Add Tasks"})
	table1.Append([]string{"3", "Update Tasks"})
	table1.Append([]string{"4", "Delete Tasks"})
	table1.Append([]string{"5", "Exit App"})
	table1.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
	table1.Render()
}

func PrintUpdateOptions() {

	table2 := tablewriter.NewWriter(os.Stdout)
	table2.SetHeader([]string{"Select below update options"})
	table2.SetRowLine(true)
	table2.Render()
	table1 := tablewriter.NewWriter(os.Stdout)
	table1.SetRowLine(true)
	table1.Append([]string{"1", "Update Name"})
	table1.Append([]string{"2", "Update Task"})
	table1.Append([]string{"3", "Update Completion Status"})
	table1.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
	table1.Render()

}

func PrintExit() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{"💫 Thank You!!! 💫"})
	table.Append([]string{"👾 Catch you on the flip side 👾"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
	table.Render()
}

func PrintList(data []models.Data) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Created On", "Completed", "Completed On"})
	table.SetRowLine(true)
	for _, ele := range data {
		table.Append([]string{fmt.Sprintf("%d", ele.Id), ele.Name, ele.Createdon.Format("2006-01-02 15:04:05"), fmt.Sprintf("%t", ele.Completed), ele.Completedon.Format("2006-01-02 15:04:05")})
	}
	table.Render()
}

func PrintUserInpInt(caller string) int {
	fmt.Println("#################################################")
  var inp string
	var flag int
	for{
		if caller == "idscreen"{
			fmt.Printf("Enter Task ID : ")
		}else{
			fmt.Printf("Enter Options : ")
		}
		
		fmt.Scanln(&inp)
		if !utility.ValidateInput(inp,&flag,caller){
			break
		}
	}
	fmt.Println("#################################################")
	return flag
}

func PrintUserInpString(inpmsg string,outmsg string) string{
	var str string
	fmt.Printf("Enter %s : ",inpmsg)
	bf := bufio.NewReader(os.Stdin)
	str,err := bf.ReadString('\n')
	if(err != nil){
		panic(err)
	}
	fmt.Println("#################################################")
	fmt.Println(outmsg)
	fmt.Println("#################################################")
	return str
}

func PrintMsg(startmsg string){
	fmt.Println(startmsg)
	fmt.Println("#################################################")
}

func PrintWithouFormat(startmsg string){
	fmt.Println(startmsg)
}

func PrintError(errmsg string){
	fmt.Println(errmsg)
}