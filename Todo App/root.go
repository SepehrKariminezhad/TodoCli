package main
import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{Use: "todo"}


func Execute(){
rootCmd.AddCommand(CmdInsert , CmdDelete , CmdRead , CmdReadall , CmdUpdate)
rootCmd.Execute()
}