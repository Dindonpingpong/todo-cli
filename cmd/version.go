package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of TodoCli",
  Long:  `All software has versions. This is Todo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Todo cli v1 -- HEAD")
  },
}