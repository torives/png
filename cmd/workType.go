package cmd

import "github.com/spf13/cobra"

var workTypeCmd = &cobra.Command{
	Use:   "worktype {add | list}",
	Short: "Creates a new work type or lists existing ones",
}

func init() {
	workTypeCmd.AddCommand(listWorkType)
}
