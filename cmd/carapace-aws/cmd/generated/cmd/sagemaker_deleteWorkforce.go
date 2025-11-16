package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteWorkforceCmd = &cobra.Command{
	Use:   "delete-workforce",
	Short: "Use this operation to delete a workforce.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteWorkforceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteWorkforceCmd).Standalone()

		sagemaker_deleteWorkforceCmd.Flags().String("workforce-name", "", "The name of the workforce.")
		sagemaker_deleteWorkforceCmd.MarkFlagRequired("workforce-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteWorkforceCmd)
}
