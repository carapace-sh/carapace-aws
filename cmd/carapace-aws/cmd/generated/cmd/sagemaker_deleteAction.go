package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteActionCmd = &cobra.Command{
	Use:   "delete-action",
	Short: "Deletes an action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteActionCmd).Standalone()

	sagemaker_deleteActionCmd.Flags().String("action-name", "", "The name of the action to delete.")
	sagemaker_deleteActionCmd.MarkFlagRequired("action-name")
	sagemakerCmd.AddCommand(sagemaker_deleteActionCmd)
}
