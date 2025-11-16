package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var costOptimizationHub_listEnrollmentStatusesCmd = &cobra.Command{
	Use:   "list-enrollment-statuses",
	Short: "Retrieves the enrollment status for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(costOptimizationHub_listEnrollmentStatusesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(costOptimizationHub_listEnrollmentStatusesCmd).Standalone()

		costOptimizationHub_listEnrollmentStatusesCmd.Flags().String("account-id", "", "The account ID of a member account in the organization.")
		costOptimizationHub_listEnrollmentStatusesCmd.Flags().String("include-organization-info", "", "Indicates whether to return the enrollment status for the organization.")
		costOptimizationHub_listEnrollmentStatusesCmd.Flags().String("max-results", "", "The maximum number of objects that are returned for the request.")
		costOptimizationHub_listEnrollmentStatusesCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	})
	costOptimizationHubCmd.AddCommand(costOptimizationHub_listEnrollmentStatusesCmd)
}
