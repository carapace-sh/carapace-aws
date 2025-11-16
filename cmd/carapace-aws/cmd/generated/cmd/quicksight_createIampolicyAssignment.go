package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createIampolicyAssignmentCmd = &cobra.Command{
	Use:   "create-iampolicy-assignment",
	Short: "Creates an assignment with one specified IAM policy, identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createIampolicyAssignmentCmd).Standalone()

	quicksight_createIampolicyAssignmentCmd.Flags().String("assignment-name", "", "The name of the assignment, also called a rule.")
	quicksight_createIampolicyAssignmentCmd.Flags().String("assignment-status", "", "The status of the assignment.")
	quicksight_createIampolicyAssignmentCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account where you want to assign an IAM policy to Amazon Quick Sight users or groups.")
	quicksight_createIampolicyAssignmentCmd.Flags().String("identities", "", "The Amazon Quick Sight users, groups, or both that you want to assign the policy to.")
	quicksight_createIampolicyAssignmentCmd.Flags().String("namespace", "", "The namespace that contains the assignment.")
	quicksight_createIampolicyAssignmentCmd.Flags().String("policy-arn", "", "The ARN for the IAM policy to apply to the Amazon Quick Sight users and groups specified in this assignment.")
	quicksight_createIampolicyAssignmentCmd.MarkFlagRequired("assignment-name")
	quicksight_createIampolicyAssignmentCmd.MarkFlagRequired("assignment-status")
	quicksight_createIampolicyAssignmentCmd.MarkFlagRequired("aws-account-id")
	quicksight_createIampolicyAssignmentCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_createIampolicyAssignmentCmd)
}
