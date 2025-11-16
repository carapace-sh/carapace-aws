package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getComplianceSummaryByResourceTypeCmd = &cobra.Command{
	Use:   "get-compliance-summary-by-resource-type",
	Short: "Returns the number of resources that are compliant and the number that are noncompliant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getComplianceSummaryByResourceTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getComplianceSummaryByResourceTypeCmd).Standalone()

		config_getComplianceSummaryByResourceTypeCmd.Flags().String("resource-types", "", "Specify one or more resource types to get the number of resources that are compliant and the number that are noncompliant for each resource type.")
	})
	configCmd.AddCommand(config_getComplianceSummaryByResourceTypeCmd)
}
