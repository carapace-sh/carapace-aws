package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getInsightsCmd = &cobra.Command{
	Use:   "get-insights",
	Short: "Lists and describes insights for the specified insight ARNs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getInsightsCmd).Standalone()

	securityhub_getInsightsCmd.Flags().String("insight-arns", "", "The ARNs of the insights to describe.")
	securityhub_getInsightsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
	securityhub_getInsightsCmd.Flags().String("next-token", "", "The token that is required for pagination.")
	securityhubCmd.AddCommand(securityhub_getInsightsCmd)
}
