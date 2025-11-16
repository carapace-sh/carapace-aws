package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateActionCmd = &cobra.Command{
	Use:   "update-action",
	Short: "Updates an action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateActionCmd).Standalone()

	sagemaker_updateActionCmd.Flags().String("action-name", "", "The name of the action to update.")
	sagemaker_updateActionCmd.Flags().String("description", "", "The new description for the action.")
	sagemaker_updateActionCmd.Flags().String("properties", "", "The new list of properties.")
	sagemaker_updateActionCmd.Flags().String("properties-to-remove", "", "A list of properties to remove.")
	sagemaker_updateActionCmd.Flags().String("status", "", "The new status for the action.")
	sagemaker_updateActionCmd.MarkFlagRequired("action-name")
	sagemakerCmd.AddCommand(sagemaker_updateActionCmd)
}
