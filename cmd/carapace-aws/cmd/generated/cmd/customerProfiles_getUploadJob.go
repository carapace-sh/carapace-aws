package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getUploadJobCmd = &cobra.Command{
	Use:   "get-upload-job",
	Short: "This API retrieves the details of a specific upload job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getUploadJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getUploadJobCmd).Standalone()

		customerProfiles_getUploadJobCmd.Flags().String("domain-name", "", "The unique name of the domain containing the upload job.")
		customerProfiles_getUploadJobCmd.Flags().String("job-id", "", "The unique identifier of the upload job to retrieve.")
		customerProfiles_getUploadJobCmd.MarkFlagRequired("domain-name")
		customerProfiles_getUploadJobCmd.MarkFlagRequired("job-id")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getUploadJobCmd)
}
