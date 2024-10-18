package main

import(
	"github.com/spf13/cobra"
	"fmt"
	"Todo_app/DBinit"
	"strconv"
	"strings"
)


var CmdUpdate = &cobra.Command{
    Use:   "update [ID to Update] [Updated log]",
    Short: "Update the log for the given ID",
    Long: "",
	Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
    id, err := strconv.ParseUint(args[0] , 10 , 64) 
		if err != nil {
			fmt.Println("Please provide a valid ID.")
			return
		}
    RemoveFromStr(args , 0)
	DBinit.Update(id ,strings.Join(args, " "))
	fmt.Println("Update Complete")
    },
  }

