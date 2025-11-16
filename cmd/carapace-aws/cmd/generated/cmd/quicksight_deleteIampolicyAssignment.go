package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteIampolicyAssignmentCmd = &cobra.Command{
	Use:   "delete-iampolicy-assignment",
	Short: "Deletes an existing IAM policy assignment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteIampolicyAssignmentCmd).Standalone()

	quicksight_deleteIampolicyAssignmentCmd.Flags().String("assignment-name", "", "The name of the assignment.")
	quicksight_deleteIampolicyAssignmentCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID where you want to delete the IAM policy assignment.")
	quicksight_deleteIampolicyAssignmentCmd.Flags().String("namespace", "", "The namespace that contains the assignment.")
	quicksight_deleteIampolicyAssignmentCmd.MarkFlagRequired("assignment-name")
	quicksight_deleteIampolicyAssignmentCmd.MarkFlagRequired("aws-account-id")
	quicksight_deleteIampolicyAssignmentCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_deleteIampolicyAssignmentCmd)
}
