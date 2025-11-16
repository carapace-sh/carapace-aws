package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var costOptimizationHub_updatePreferencesCmd = &cobra.Command{
	Use:   "update-preferences",
	Short: "Updates a set of preferences for an account in order to add account-specific preferences into the service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(costOptimizationHub_updatePreferencesCmd).Standalone()

	costOptimizationHub_updatePreferencesCmd.Flags().String("member-account-discount-visibility", "", "Sets the \"member account discount visibility\" preference.")
	costOptimizationHub_updatePreferencesCmd.Flags().String("preferred-commitment", "", "Sets the preferences for how Reserved Instances and Savings Plans cost-saving opportunities are prioritized in terms of payment option and term length.")
	costOptimizationHub_updatePreferencesCmd.Flags().String("savings-estimation-mode", "", "Sets the \"savings estimation mode\" preference.")
	costOptimizationHubCmd.AddCommand(costOptimizationHub_updatePreferencesCmd)
}
