package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_createInsightCmd = &cobra.Command{
	Use:   "create-insight",
	Short: "Creates a custom insight in Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_createInsightCmd).Standalone()

	securityhub_createInsightCmd.Flags().String("filters", "", "One or more attributes used to filter the findings included in the insight.")
	securityhub_createInsightCmd.Flags().String("group-by-attribute", "", "The attribute used to group the findings for the insight.")
	securityhub_createInsightCmd.Flags().String("name", "", "The name of the custom insight to create.")
	securityhub_createInsightCmd.MarkFlagRequired("filters")
	securityhub_createInsightCmd.MarkFlagRequired("group-by-attribute")
	securityhub_createInsightCmd.MarkFlagRequired("name")
	securityhubCmd.AddCommand(securityhub_createInsightCmd)
}
