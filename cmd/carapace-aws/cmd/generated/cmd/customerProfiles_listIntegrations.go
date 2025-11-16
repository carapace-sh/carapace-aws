package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listIntegrationsCmd = &cobra.Command{
	Use:   "list-integrations",
	Short: "Lists all of the integrations in your domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listIntegrationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listIntegrationsCmd).Standalone()

		customerProfiles_listIntegrationsCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_listIntegrationsCmd.Flags().String("include-hidden", "", "Boolean to indicate if hidden integration should be returned.")
		customerProfiles_listIntegrationsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
		customerProfiles_listIntegrationsCmd.Flags().String("next-token", "", "The pagination token from the previous ListIntegrations API call.")
		customerProfiles_listIntegrationsCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listIntegrationsCmd)
}
