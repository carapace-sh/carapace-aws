package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeComplianceByResourceCmd = &cobra.Command{
	Use:   "describe-compliance-by-resource",
	Short: "Indicates whether the specified Amazon Web Services resources are compliant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeComplianceByResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeComplianceByResourceCmd).Standalone()

		config_describeComplianceByResourceCmd.Flags().String("compliance-types", "", "Filters the results by compliance.")
		config_describeComplianceByResourceCmd.Flags().String("limit", "", "The maximum number of evaluation results returned on each page.")
		config_describeComplianceByResourceCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_describeComplianceByResourceCmd.Flags().String("resource-id", "", "The ID of the Amazon Web Services resource for which you want compliance information.")
		config_describeComplianceByResourceCmd.Flags().String("resource-type", "", "The types of Amazon Web Services resources for which you want compliance information (for example, `AWS::EC2::Instance`).")
	})
	configCmd.AddCommand(config_describeComplianceByResourceCmd)
}
