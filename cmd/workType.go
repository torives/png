package cmd

import "github.com/spf13/cobra"

var workTypeCmd = &cobra.Command{
	Use:   "worktype {add | list}",
	Short: "create a new work type or list existing ones",
	Long:  "Creates a new work type or list existing ones",
}

func init() {
	workTypeCmd.AddCommand(listWorkTypes)
	workTypeCmd.AddCommand(addWorkType)
}
