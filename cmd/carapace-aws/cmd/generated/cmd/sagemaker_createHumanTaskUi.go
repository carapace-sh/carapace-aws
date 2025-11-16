package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createHumanTaskUiCmd = &cobra.Command{
	Use:   "create-human-task-ui",
	Short: "Defines the settings you will use for the human review workflow user interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createHumanTaskUiCmd).Standalone()

	sagemaker_createHumanTaskUiCmd.Flags().String("human-task-ui-name", "", "The name of the user interface you are creating.")
	sagemaker_createHumanTaskUiCmd.Flags().String("tags", "", "An array of key-value pairs that contain metadata to help you categorize and organize a human review workflow user interface.")
	sagemaker_createHumanTaskUiCmd.Flags().String("ui-template", "", "")
	sagemaker_createHumanTaskUiCmd.MarkFlagRequired("human-task-ui-name")
	sagemaker_createHumanTaskUiCmd.MarkFlagRequired("ui-template")
	sagemakerCmd.AddCommand(sagemaker_createHumanTaskUiCmd)
}
