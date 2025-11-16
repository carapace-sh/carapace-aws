package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listAccountIntegrationsCmd = &cobra.Command{
	Use:   "list-account-integrations",
	Short: "Lists all of the integrations associated to a specific URI in the AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listAccountIntegrationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listAccountIntegrationsCmd).Standalone()

		customerProfiles_listAccountIntegrationsCmd.Flags().String("include-hidden", "", "Boolean to indicate if hidden integration should be returned.")
		customerProfiles_listAccountIntegrationsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
		customerProfiles_listAccountIntegrationsCmd.Flags().String("next-token", "", "The pagination token from the previous ListAccountIntegrations API call.")
		customerProfiles_listAccountIntegrationsCmd.Flags().String("uri", "", "The URI of the S3 bucket or any other type of data source.")
		customerProfiles_listAccountIntegrationsCmd.MarkFlagRequired("uri")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listAccountIntegrationsCmd)
}
