package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateInsightCmd = &cobra.Command{
	Use:   "update-insight",
	Short: "Updates the Security Hub insight identified by the specified insight ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateInsightCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_updateInsightCmd).Standalone()

		securityhub_updateInsightCmd.Flags().String("filters", "", "The updated filters that define this insight.")
		securityhub_updateInsightCmd.Flags().String("group-by-attribute", "", "The updated `GroupBy` attribute that defines this insight.")
		securityhub_updateInsightCmd.Flags().String("insight-arn", "", "The ARN of the insight that you want to update.")
		securityhub_updateInsightCmd.Flags().String("name", "", "The updated name for the insight.")
		securityhub_updateInsightCmd.MarkFlagRequired("insight-arn")
	})
	securityhubCmd.AddCommand(securityhub_updateInsightCmd)
}
