package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listApplicationAssignmentsCmd = &cobra.Command{
	Use:   "list-application-assignments",
	Short: "Lists Amazon Web Services account users that are assigned to an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listApplicationAssignmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listApplicationAssignmentsCmd).Standalone()

		ssoAdmin_listApplicationAssignmentsCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
		ssoAdmin_listApplicationAssignmentsCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
		ssoAdmin_listApplicationAssignmentsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ssoAdmin_listApplicationAssignmentsCmd.MarkFlagRequired("application-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listApplicationAssignmentsCmd)
}
