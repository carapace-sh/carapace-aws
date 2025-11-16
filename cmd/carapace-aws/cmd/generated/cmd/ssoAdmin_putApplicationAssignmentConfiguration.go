package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_putApplicationAssignmentConfigurationCmd = &cobra.Command{
	Use:   "put-application-assignment-configuration",
	Short: "Configure how users gain access to an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_putApplicationAssignmentConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_putApplicationAssignmentConfigurationCmd).Standalone()

		ssoAdmin_putApplicationAssignmentConfigurationCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
		ssoAdmin_putApplicationAssignmentConfigurationCmd.Flags().String("assignment-required", "", "If `AssignmentsRequired` is `true` (default value), users don’t have access to the application unless an assignment is created using the [CreateApplicationAssignment API](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_CreateApplicationAssignment.html).")
		ssoAdmin_putApplicationAssignmentConfigurationCmd.MarkFlagRequired("application-arn")
		ssoAdmin_putApplicationAssignmentConfigurationCmd.MarkFlagRequired("assignment-required")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_putApplicationAssignmentConfigurationCmd)
}
