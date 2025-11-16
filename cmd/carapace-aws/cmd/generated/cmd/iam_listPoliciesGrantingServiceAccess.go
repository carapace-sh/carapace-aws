package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listPoliciesGrantingServiceAccessCmd = &cobra.Command{
	Use:   "list-policies-granting-service-access",
	Short: "Retrieves a list of policies that the IAM identity (user, group, or role) can use to access each specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listPoliciesGrantingServiceAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listPoliciesGrantingServiceAccessCmd).Standalone()

		iam_listPoliciesGrantingServiceAccessCmd.Flags().String("arn", "", "The ARN of the IAM identity (user, group, or role) whose policies you want to list.")
		iam_listPoliciesGrantingServiceAccessCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listPoliciesGrantingServiceAccessCmd.Flags().String("service-namespaces", "", "The service namespace for the Amazon Web Services services whose policies you want to list.")
		iam_listPoliciesGrantingServiceAccessCmd.MarkFlagRequired("arn")
		iam_listPoliciesGrantingServiceAccessCmd.MarkFlagRequired("service-namespaces")
	})
	iamCmd.AddCommand(iam_listPoliciesGrantingServiceAccessCmd)
}
