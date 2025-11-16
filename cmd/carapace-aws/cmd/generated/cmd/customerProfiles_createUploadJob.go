package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createUploadJobCmd = &cobra.Command{
	Use:   "create-upload-job",
	Short: "Creates an Upload job to ingest data for segment imports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createUploadJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_createUploadJobCmd).Standalone()

		customerProfiles_createUploadJobCmd.Flags().String("data-expiry", "", "The expiry duration for the profiles ingested with the job.")
		customerProfiles_createUploadJobCmd.Flags().String("display-name", "", "The unique name of the upload job.")
		customerProfiles_createUploadJobCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_createUploadJobCmd.Flags().String("fields", "", "The mapping between CSV Columns and Profile Object attributes.")
		customerProfiles_createUploadJobCmd.Flags().String("unique-key", "", "The unique key columns for de-duping the profiles used to map data to the profile.")
		customerProfiles_createUploadJobCmd.MarkFlagRequired("display-name")
		customerProfiles_createUploadJobCmd.MarkFlagRequired("domain-name")
		customerProfiles_createUploadJobCmd.MarkFlagRequired("fields")
		customerProfiles_createUploadJobCmd.MarkFlagRequired("unique-key")
	})
	customerProfilesCmd.AddCommand(customerProfiles_createUploadJobCmd)
}
