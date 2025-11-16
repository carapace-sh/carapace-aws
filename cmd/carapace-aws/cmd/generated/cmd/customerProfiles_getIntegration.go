package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getIntegrationCmd = &cobra.Command{
	Use:   "get-integration",
	Short: "Returns an integration for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getIntegrationCmd).Standalone()

	customerProfiles_getIntegrationCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_getIntegrationCmd.Flags().String("uri", "", "The URI of the S3 bucket or any other type of data source.")
	customerProfiles_getIntegrationCmd.MarkFlagRequired("domain-name")
	customerProfiles_getIntegrationCmd.MarkFlagRequired("uri")
	customerProfilesCmd.AddCommand(customerProfiles_getIntegrationCmd)
}
