package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeIampolicyAssignmentCmd = &cobra.Command{
	Use:   "describe-iampolicy-assignment",
	Short: "Describes an existing IAM policy assignment, as specified by the assignment name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeIampolicyAssignmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeIampolicyAssignmentCmd).Standalone()

		quicksight_describeIampolicyAssignmentCmd.Flags().String("assignment-name", "", "The name of the assignment, also called a rule.")
		quicksight_describeIampolicyAssignmentCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the assignment that you want to describe.")
		quicksight_describeIampolicyAssignmentCmd.Flags().String("namespace", "", "The namespace that contains the assignment.")
		quicksight_describeIampolicyAssignmentCmd.MarkFlagRequired("assignment-name")
		quicksight_describeIampolicyAssignmentCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeIampolicyAssignmentCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_describeIampolicyAssignmentCmd)
}
