// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"ghn-create-go-app/pkg/registry"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"ghn-create-go-app/pkg/cga"
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(
		&useCustomTemplate,
		"template", "t", false,
		"enables to use custom backend templates",
	)
}

// createCmd represents the `create` command.
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create a new project via interactive UI",
	Long:    "\nCreate a new project via interactive UI.",
	RunE:    runCreateCmd,
}

// runCreateCmd represents runner for the `create` command.
func runCreateCmd(cmd *cobra.Command, args []string) error {
	// Start message.
	cga.ShowMessage(
		"",
		fmt.Sprintf(
			"Create a new project via Create Go App CLI v%v...",
			registry.CLIVersion,
		),
		true, true,
	)

	// Default survey.
	if err := survey.Ask(
		registry.CreateQuestions,
		&createAnswers,
		survey.WithIcons(surveyIconsConfig),
	); err != nil {
		return cga.ShowError(err.Error())
	}

	// Define variables for better display.
	backend = fmt.Sprintf(
		"github.com/create-go-app/%v-go-template",
		strings.ReplaceAll(createAnswers.Backend, "/", "_"),
	)

	// Catch the cancel action (hit "n" in the last question).
	if (!createAnswers.AgreeCreation && !useCustomTemplate) || (!customCreateAnswers.AgreeCreation && useCustomTemplate) {
		cga.ShowMessage(
			"",
			"Oh no! You said \"no\", so I won't create anything. Hope to see you soon!",
			true, true,
		)
		return nil
	}

	// Start timer.
	startTimer := time.Now()

	/*
		The project's backend part creation.
	*/
	// Clone backend files from git repository.
	if err := cga.GitClone("backend", backend); err != nil {
		return cga.ShowError(err.Error())
	}

	// Show success report.
	cga.ShowMessage(
		"success",
		fmt.Sprintf("Backend was created with template `%v`!", backend),
		true, false,
	)

	/*
		Cleanup project.
	*/
	cga.RemoveFolders("backend", []string{".git", ".github"})

	// Stop timer.
	stopTimer := cga.CalculateDurationTime(startTimer)
	cga.ShowMessage(
		"info",
		fmt.Sprintf("Completed in %v seconds!", stopTimer),
		true, true,
	)
	cga.ShowMessage(
		"",
		"Have a happy new project! :)",
		false, true,
	)

	return nil
}
