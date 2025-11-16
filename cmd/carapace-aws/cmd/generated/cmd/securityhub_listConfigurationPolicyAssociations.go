package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listConfigurationPolicyAssociationsCmd = &cobra.Command{
	Use:   "list-configuration-policy-associations",
	Short: "Provides information about the associations for your configuration policies and self-managed behavior.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listConfigurationPolicyAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_listConfigurationPolicyAssociationsCmd).Standalone()

		securityhub_listConfigurationPolicyAssociationsCmd.Flags().String("filters", "", "Options for filtering the `ListConfigurationPolicyAssociations` response.")
		securityhub_listConfigurationPolicyAssociationsCmd.Flags().String("max-results", "", "The maximum number of results that's returned by `ListConfigurationPolicies` in each page of the response.")
		securityhub_listConfigurationPolicyAssociationsCmd.Flags().String("next-token", "", "The `NextToken` value that's returned from a previous paginated `ListConfigurationPolicyAssociations` request where `MaxResults` was used but the results exceeded the value of that parameter.")
	})
	securityhubCmd.AddCommand(securityhub_listConfigurationPolicyAssociationsCmd)
}
