package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteHumanTaskUiCmd = &cobra.Command{
	Use:   "delete-human-task-ui",
	Short: "Use this operation to delete a human task user interface (worker task template).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteHumanTaskUiCmd).Standalone()

	sagemaker_deleteHumanTaskUiCmd.Flags().String("human-task-ui-name", "", "The name of the human task user interface (work task template) you want to delete.")
	sagemaker_deleteHumanTaskUiCmd.MarkFlagRequired("human-task-ui-name")
	sagemakerCmd.AddCommand(sagemaker_deleteHumanTaskUiCmd)
}
