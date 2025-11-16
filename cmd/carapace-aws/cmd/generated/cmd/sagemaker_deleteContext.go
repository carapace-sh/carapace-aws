package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteContextCmd = &cobra.Command{
	Use:   "delete-context",
	Short: "Deletes an context.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteContextCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteContextCmd).Standalone()

		sagemaker_deleteContextCmd.Flags().String("context-name", "", "The name of the context to delete.")
		sagemaker_deleteContextCmd.MarkFlagRequired("context-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteContextCmd)
}
