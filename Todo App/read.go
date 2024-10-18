package main

import(
	"github.com/spf13/cobra"
	"Todo_app/DBinit"
	"strconv"
	"fmt"
)


var CmdReadall= &cobra.Command{
    Use:   "readall",
    Short: "read all the data from your DB",
    Long: "",
    Run: func(cmd *cobra.Command, args []string) {
		DBinit.Read_all()
    },
  }


  var CmdRead= &cobra.Command{
    Use:   "read [ID to read]",
    Short: "read the data with the given ID",
    Long: "",
	Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0] , 10 , 64)
		if err != nil {
			fmt.Println("Please provide a valid ID.")
			return
		}
		DBinit.Read(id)
    },
  }