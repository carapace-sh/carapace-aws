package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteIntegrationCmd = &cobra.Command{
	Use:   "delete-integration",
	Short: "Removes an integration from a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteIntegrationCmd).Standalone()

	customerProfiles_deleteIntegrationCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_deleteIntegrationCmd.Flags().String("uri", "", "The URI of the S3 bucket or any other type of data source.")
	customerProfiles_deleteIntegrationCmd.MarkFlagRequired("domain-name")
	customerProfiles_deleteIntegrationCmd.MarkFlagRequired("uri")
	customerProfilesCmd.AddCommand(customerProfiles_deleteIntegrationCmd)
}
