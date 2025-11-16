package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listApplicationAssignmentsForPrincipalCmd = &cobra.Command{
	Use:   "list-application-assignments-for-principal",
	Short: "Lists the applications to which a specified principal is assigned.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listApplicationAssignmentsForPrincipalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listApplicationAssignmentsForPrincipalCmd).Standalone()

		ssoAdmin_listApplicationAssignmentsForPrincipalCmd.Flags().String("filter", "", "Filters the output to include only assignments associated with the application that has the specified ARN.")
		ssoAdmin_listApplicationAssignmentsForPrincipalCmd.Flags().String("instance-arn", "", "Specifies the instance of IAM Identity Center that contains principal and applications.")
		ssoAdmin_listApplicationAssignmentsForPrincipalCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
		ssoAdmin_listApplicationAssignmentsForPrincipalCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ssoAdmin_listApplicationAssignmentsForPrincipalCmd.Flags().String("principal-id", "", "Specifies the unique identifier of the principal for which you want to retrieve its assignments.")
		ssoAdmin_listApplicationAssignmentsForPrincipalCmd.Flags().String("principal-type", "", "Specifies the type of the principal for which you want to retrieve its assignments.")
		ssoAdmin_listApplicationAssignmentsForPrincipalCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_listApplicationAssignmentsForPrincipalCmd.MarkFlagRequired("principal-id")
		ssoAdmin_listApplicationAssignmentsForPrincipalCmd.MarkFlagRequired("principal-type")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listApplicationAssignmentsForPrincipalCmd)
}
