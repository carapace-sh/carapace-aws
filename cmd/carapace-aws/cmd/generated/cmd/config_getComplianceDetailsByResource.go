package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getComplianceDetailsByResourceCmd = &cobra.Command{
	Use:   "get-compliance-details-by-resource",
	Short: "Returns the evaluation results for the specified Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getComplianceDetailsByResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getComplianceDetailsByResourceCmd).Standalone()

		config_getComplianceDetailsByResourceCmd.Flags().String("compliance-types", "", "Filters the results by compliance.")
		config_getComplianceDetailsByResourceCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_getComplianceDetailsByResourceCmd.Flags().String("resource-evaluation-id", "", "The unique ID of Amazon Web Services resource execution for which you want to retrieve evaluation results.")
		config_getComplianceDetailsByResourceCmd.Flags().String("resource-id", "", "The ID of the Amazon Web Services resource for which you want compliance information.")
		config_getComplianceDetailsByResourceCmd.Flags().String("resource-type", "", "The type of the Amazon Web Services resource for which you want compliance information.")
	})
	configCmd.AddCommand(config_getComplianceDetailsByResourceCmd)
}
