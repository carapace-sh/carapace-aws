package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getEnrollmentStatusCmd = &cobra.Command{
	Use:   "get-enrollment-status",
	Short: "Returns the enrollment (opt in) status of an account to the Compute Optimizer service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getEnrollmentStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getEnrollmentStatusCmd).Standalone()

	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getEnrollmentStatusCmd)
}
