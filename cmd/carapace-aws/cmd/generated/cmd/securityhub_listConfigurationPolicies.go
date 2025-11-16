package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listConfigurationPoliciesCmd = &cobra.Command{
	Use:   "list-configuration-policies",
	Short: "Lists the configuration policies that the Security Hub delegated administrator has created for your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listConfigurationPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_listConfigurationPoliciesCmd).Standalone()

		securityhub_listConfigurationPoliciesCmd.Flags().String("max-results", "", "The maximum number of results that's returned by `ListConfigurationPolicies` in each page of the response.")
		securityhub_listConfigurationPoliciesCmd.Flags().String("next-token", "", "The NextToken value that's returned from a previous paginated `ListConfigurationPolicies` request where `MaxResults` was used but the results exceeded the value of that parameter.")
	})
	securityhubCmd.AddCommand(securityhub_listConfigurationPoliciesCmd)
}
