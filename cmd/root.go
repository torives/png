/*
Copyright © 2024 Victor Yves Crispim
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "png",
	Short: "Project Number Generator",
	Long: `PNG is a CLI program created to generate unique IDs for Harmony's 
project report sheets.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
