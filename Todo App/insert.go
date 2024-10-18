package main

import(
	"github.com/spf13/cobra"
	"fmt"
	"strings"
	"Todo_app/DBinit"
  "strconv"
)

func RemoveFromStr(s[]string , index int)[]string{
  s = append(s[:index], s[index+1:]...)
  x := len(s)
  s[x-1] = ""
  return s
}

var CmdInsert = &cobra.Command{
    Use:   "insert [a unique ID] [string to insert]",
    Short: "Create a row in your DB",
    Long: "",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
    id, err := strconv.ParseUint(args[0] , 10 , 64) 
		if err != nil {
			fmt.Println("Please provide a valid ID.")
			return
		}
    RemoveFromStr(args , 0)
	  DBinit.Insert(id ,strings.Join(args, " "))
	  fmt.Printf("Log %v has been inserted\n" , strings.Join(args, " "))
    },
  }