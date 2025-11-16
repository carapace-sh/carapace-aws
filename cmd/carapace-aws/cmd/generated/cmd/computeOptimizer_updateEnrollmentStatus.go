package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_updateEnrollmentStatusCmd = &cobra.Command{
	Use:   "update-enrollment-status",
	Short: "Updates the enrollment (opt in and opt out) status of an account to the Compute Optimizer service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_updateEnrollmentStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_updateEnrollmentStatusCmd).Standalone()

		computeOptimizer_updateEnrollmentStatusCmd.Flags().String("include-member-accounts", "", "Indicates whether to enroll member accounts of the organization if the account is the management account of an organization.")
		computeOptimizer_updateEnrollmentStatusCmd.Flags().String("status", "", "The new enrollment status of the account.")
		computeOptimizer_updateEnrollmentStatusCmd.MarkFlagRequired("status")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_updateEnrollmentStatusCmd)
}
