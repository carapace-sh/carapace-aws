package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var costOptimizationHub_updateEnrollmentStatusCmd = &cobra.Command{
	Use:   "update-enrollment-status",
	Short: "Updates the enrollment (opt in and opt out) status of an account to the Cost Optimization Hub service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(costOptimizationHub_updateEnrollmentStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(costOptimizationHub_updateEnrollmentStatusCmd).Standalone()

		costOptimizationHub_updateEnrollmentStatusCmd.Flags().Bool("include-member-accounts", false, "Indicates whether to enroll member accounts of the organization if the account is the management account or delegated administrator.")
		costOptimizationHub_updateEnrollmentStatusCmd.Flags().Bool("no-include-member-accounts", false, "Indicates whether to enroll member accounts of the organization if the account is the management account or delegated administrator.")
		costOptimizationHub_updateEnrollmentStatusCmd.Flags().String("status", "", "Sets the account status.")
		costOptimizationHub_updateEnrollmentStatusCmd.Flag("no-include-member-accounts").Hidden = true
		costOptimizationHub_updateEnrollmentStatusCmd.MarkFlagRequired("status")
	})
	costOptimizationHubCmd.AddCommand(costOptimizationHub_updateEnrollmentStatusCmd)
}
