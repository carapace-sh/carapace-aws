package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listUploadJobsCmd = &cobra.Command{
	Use:   "list-upload-jobs",
	Short: "This API retrieves a list of upload jobs for the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listUploadJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listUploadJobsCmd).Standalone()

		customerProfiles_listUploadJobsCmd.Flags().String("domain-name", "", "The unique name of the domain to list upload jobs for.")
		customerProfiles_listUploadJobsCmd.Flags().String("max-results", "", "The maximum number of upload jobs to return per page.")
		customerProfiles_listUploadJobsCmd.Flags().String("next-token", "", "The pagination token from the previous call to retrieve the next page of results.")
		customerProfiles_listUploadJobsCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listUploadJobsCmd)
}
