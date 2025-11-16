package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_listBackendJobsCmd = &cobra.Command{
	Use:   "list-backend-jobs",
	Short: "Lists the jobs for the backend of an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_listBackendJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_listBackendJobsCmd).Standalone()

		amplifybackend_listBackendJobsCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_listBackendJobsCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_listBackendJobsCmd.Flags().String("job-id", "", "The ID for the job.")
		amplifybackend_listBackendJobsCmd.Flags().String("max-results", "", "The maximum number of results that you want in the response.")
		amplifybackend_listBackendJobsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		amplifybackend_listBackendJobsCmd.Flags().String("operation", "", "Filters the list of response objects to include only those with the specified operation name.")
		amplifybackend_listBackendJobsCmd.Flags().String("status", "", "Filters the list of response objects to include only those with the specified status.")
		amplifybackend_listBackendJobsCmd.MarkFlagRequired("app-id")
		amplifybackend_listBackendJobsCmd.MarkFlagRequired("backend-environment-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_listBackendJobsCmd)
}
