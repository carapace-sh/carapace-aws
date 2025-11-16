package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getEnrollmentStatusesForOrganizationCmd = &cobra.Command{
	Use:   "get-enrollment-statuses-for-organization",
	Short: "Returns the Compute Optimizer enrollment (opt-in) status of organization member accounts, if your account is an organization management account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getEnrollmentStatusesForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getEnrollmentStatusesForOrganizationCmd).Standalone()

		computeOptimizer_getEnrollmentStatusesForOrganizationCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of account enrollment statuses.")
		computeOptimizer_getEnrollmentStatusesForOrganizationCmd.Flags().String("max-results", "", "The maximum number of account enrollment statuses to return with a single request.")
		computeOptimizer_getEnrollmentStatusesForOrganizationCmd.Flags().String("next-token", "", "The token to advance to the next page of account enrollment statuses.")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getEnrollmentStatusesForOrganizationCmd)
}
