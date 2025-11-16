package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_deleteSamplingRuleCmd = &cobra.Command{
	Use:   "delete-sampling-rule",
	Short: "Deletes a sampling rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_deleteSamplingRuleCmd).Standalone()

	xray_deleteSamplingRuleCmd.Flags().String("rule-arn", "", "The ARN of the sampling rule.")
	xray_deleteSamplingRuleCmd.Flags().String("rule-name", "", "The name of the sampling rule.")
	xrayCmd.AddCommand(xray_deleteSamplingRuleCmd)
}
