package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_startUploadJobCmd = &cobra.Command{
	Use:   "start-upload-job",
	Short: "This API starts the processing of an upload job to ingest profile data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_startUploadJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_startUploadJobCmd).Standalone()

		customerProfiles_startUploadJobCmd.Flags().String("domain-name", "", "The unique name of the domain containing the upload job to start.")
		customerProfiles_startUploadJobCmd.Flags().String("job-id", "", "The unique identifier of the upload job to start.")
		customerProfiles_startUploadJobCmd.MarkFlagRequired("domain-name")
		customerProfiles_startUploadJobCmd.MarkFlagRequired("job-id")
	})
	customerProfilesCmd.AddCommand(customerProfiles_startUploadJobCmd)
}
