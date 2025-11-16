package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_stopUploadJobCmd = &cobra.Command{
	Use:   "stop-upload-job",
	Short: "This API stops the processing of an upload job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_stopUploadJobCmd).Standalone()

	customerProfiles_stopUploadJobCmd.Flags().String("domain-name", "", "The unique name of the domain containing the upload job to stop.")
	customerProfiles_stopUploadJobCmd.Flags().String("job-id", "", "The unique identifier of the upload job to stop.")
	customerProfiles_stopUploadJobCmd.MarkFlagRequired("domain-name")
	customerProfiles_stopUploadJobCmd.MarkFlagRequired("job-id")
	customerProfilesCmd.AddCommand(customerProfiles_stopUploadJobCmd)
}
