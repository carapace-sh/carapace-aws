package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeHumanTaskUiCmd = &cobra.Command{
	Use:   "describe-human-task-ui",
	Short: "Returns information about the requested human task user interface (worker task template).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeHumanTaskUiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeHumanTaskUiCmd).Standalone()

		sagemaker_describeHumanTaskUiCmd.Flags().String("human-task-ui-name", "", "The name of the human task user interface (worker task template) you want information about.")
		sagemaker_describeHumanTaskUiCmd.MarkFlagRequired("human-task-ui-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeHumanTaskUiCmd)
}
