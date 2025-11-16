package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeRemediationExceptionsCmd = &cobra.Command{
	Use:   "describe-remediation-exceptions",
	Short: "Returns the details of one or more remediation exceptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeRemediationExceptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeRemediationExceptionsCmd).Standalone()

		config_describeRemediationExceptionsCmd.Flags().String("config-rule-name", "", "The name of the Config rule.")
		config_describeRemediationExceptionsCmd.Flags().String("limit", "", "The maximum number of RemediationExceptionResourceKey returned on each page.")
		config_describeRemediationExceptionsCmd.Flags().String("next-token", "", "The `nextToken` string returned in a previous request that you use to request the next page of results in a paginated response.")
		config_describeRemediationExceptionsCmd.Flags().String("resource-keys", "", "An exception list of resource exception keys to be processed with the current request.")
		config_describeRemediationExceptionsCmd.MarkFlagRequired("config-rule-name")
	})
	configCmd.AddCommand(config_describeRemediationExceptionsCmd)
}
