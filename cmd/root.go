/*
Copyright © 2024 Victor Yves Crispim
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/torives/png/model"
)

//FIXME: refactor "project" references to make it clear that it is a report sheet code for a project

var databaseDsn = "file:png.sqlite"

type ErrOpenDatabase struct {
	dbErr error
}

func (e ErrOpenDatabase) Error() string {
	return fmt.Sprintf("failed to open database: %s", e.dbErr)
}

var rootCmd = &cobra.Command{
	Use:   "png",
	Short: "Project Number Generator",
	Long: `PNG is a CLI program created to generate unique IDs for Harmony's 
project report sheets.`,
}

func Execute() {
	// if the program name is the only parameter
	if len(os.Args) == 1 {
		err := runInteractivePrompt()
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// Starts an interactive prompt that will add the missing program arguments
// necessary to invoke invoke a command. Should only be used if the user started
// the program without providing any parameters.
func runInteractivePrompt() error {
	defaultPromptTemplate := &promptui.PromptTemplates{
		Valid:           "{{ . | green }} ",
		Invalid:         "{{ . | yellow }} ",
		ValidationError: "{{ . | red }} ",
		Success:         fmt.Sprintf("%s {{ . | faint }} ", promptui.IconGood),
	}

	topLevelActions := []string{
		addProject.Short,
		listProjects.Short,
		addTeam.Short,
		listTeams.Short,
		addWorkType.Short,
		listWorkTypes.Short,
	}
	selectPrompt := promptui.Select{
		Label:     "What would you like to do?",
		Items:     topLevelActions,
		Size:      len(topLevelActions),
		Templates: &promptui.SelectTemplates{Label: "{{ . }}"},
	}

	_, action, err := selectPrompt.Run()
	if err != nil {
		return err
	}

	switch action {
	case addProject.Short:
		teamPrompt := promptui.Prompt{
			Label:     "Which team is responsible for the project?",
			Validate:  model.ValidateTeamName,
			Templates: defaultPromptTemplate,
		}
		team, err := teamPrompt.Run()
		if err != nil {
			return err
		}

		workTypePrompt := promptui.Prompt{Label: "What is the work type?",
			Validate:  model.ValidateWorkTypeName,
			Templates: defaultPromptTemplate,
		}
		workType, err := workTypePrompt.Run()
		if err != nil {
			return err
		}

		os.Args = append(os.Args, []string{"project", "add", "-t", team, "-w", workType}...)
	case listProjects.Short:
		os.Args = append(os.Args, []string{"project", "list"}...)
	case addTeam.Short:
		namePrompt := promptui.Prompt{
			Label:     "Name:",
			Validate:  model.ValidateTeamName,
			Templates: defaultPromptTemplate,
		}
		name, err := namePrompt.Run()
		if err != nil {
			return err
		}

		os.Args = append(os.Args, []string{"team", "add", name}...)
	case listTeams.Short:
		os.Args = append(os.Args, []string{"team", "list"}...)
	case addWorkType.Short:
		namePrompt := promptui.Prompt{
			Label:     "Name:",
			Validate:  model.ValidateWorkTypeName,
			Templates: defaultPromptTemplate,
		}
		name, err := namePrompt.Run()
		if err != nil {
			return err
		}

		os.Args = append(os.Args, []string{"worktype", "add", name}...)
	case listWorkTypes.Short:
		os.Args = append(os.Args, []string{"worktype", "list"}...)
	}

	return nil
}

func init() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(teamCmd)
	rootCmd.AddCommand(workTypeCmd)
}
