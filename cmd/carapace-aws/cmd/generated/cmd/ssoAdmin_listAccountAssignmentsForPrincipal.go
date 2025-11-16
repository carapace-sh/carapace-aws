package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listAccountAssignmentsForPrincipalCmd = &cobra.Command{
	Use:   "list-account-assignments-for-principal",
	Short: "Retrieves a list of the IAM Identity Center associated Amazon Web Services accounts that the principal has access to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listAccountAssignmentsForPrincipalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listAccountAssignmentsForPrincipalCmd).Standalone()

		ssoAdmin_listAccountAssignmentsForPrincipalCmd.Flags().String("filter", "", "Specifies an Amazon Web Services account ID number.")
		ssoAdmin_listAccountAssignmentsForPrincipalCmd.Flags().String("instance-arn", "", "Specifies the ARN of the instance of IAM Identity Center that contains the principal.")
		ssoAdmin_listAccountAssignmentsForPrincipalCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
		ssoAdmin_listAccountAssignmentsForPrincipalCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ssoAdmin_listAccountAssignmentsForPrincipalCmd.Flags().String("principal-id", "", "Specifies the principal for which you want to retrieve the list of account assignments.")
		ssoAdmin_listAccountAssignmentsForPrincipalCmd.Flags().String("principal-type", "", "Specifies the type of the principal.")
		ssoAdmin_listAccountAssignmentsForPrincipalCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_listAccountAssignmentsForPrincipalCmd.MarkFlagRequired("principal-id")
		ssoAdmin_listAccountAssignmentsForPrincipalCmd.MarkFlagRequired("principal-type")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listAccountAssignmentsForPrincipalCmd)
}
