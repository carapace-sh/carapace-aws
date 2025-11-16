package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_describeActionTargetsCmd = &cobra.Command{
	Use:   "describe-action-targets",
	Short: "Returns a list of the custom action targets in Security Hub in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_describeActionTargetsCmd).Standalone()

	securityhub_describeActionTargetsCmd.Flags().String("action-target-arns", "", "A list of custom action target ARNs for the custom action targets to retrieve.")
	securityhub_describeActionTargetsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	securityhub_describeActionTargetsCmd.Flags().String("next-token", "", "The token that is required for pagination.")
	securityhubCmd.AddCommand(securityhub_describeActionTargetsCmd)
}
