package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getUploadJobPathCmd = &cobra.Command{
	Use:   "get-upload-job-path",
	Short: "This API retrieves the pre-signed URL and client token for uploading the file associated with the upload job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getUploadJobPathCmd).Standalone()

	customerProfiles_getUploadJobPathCmd.Flags().String("domain-name", "", "The unique name of the domain containing the upload job.")
	customerProfiles_getUploadJobPathCmd.Flags().String("job-id", "", "The unique identifier of the upload job to retrieve the upload path for.")
	customerProfiles_getUploadJobPathCmd.MarkFlagRequired("domain-name")
	customerProfiles_getUploadJobPathCmd.MarkFlagRequired("job-id")
	customerProfilesCmd.AddCommand(customerProfiles_getUploadJobPathCmd)
}
