package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateIampolicyAssignmentCmd = &cobra.Command{
	Use:   "update-iampolicy-assignment",
	Short: "Updates an existing IAM policy assignment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateIampolicyAssignmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateIampolicyAssignmentCmd).Standalone()

		quicksight_updateIampolicyAssignmentCmd.Flags().String("assignment-name", "", "The name of the assignment, also called a rule.")
		quicksight_updateIampolicyAssignmentCmd.Flags().String("assignment-status", "", "The status of the assignment.")
		quicksight_updateIampolicyAssignmentCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the IAM policy assignment.")
		quicksight_updateIampolicyAssignmentCmd.Flags().String("identities", "", "The Amazon Quick Sight users, groups, or both that you want to assign the policy to.")
		quicksight_updateIampolicyAssignmentCmd.Flags().String("namespace", "", "The namespace of the assignment.")
		quicksight_updateIampolicyAssignmentCmd.Flags().String("policy-arn", "", "The ARN for the IAM policy to apply to the Amazon Quick Sight users and groups specified in this assignment.")
		quicksight_updateIampolicyAssignmentCmd.MarkFlagRequired("assignment-name")
		quicksight_updateIampolicyAssignmentCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateIampolicyAssignmentCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_updateIampolicyAssignmentCmd)
}
