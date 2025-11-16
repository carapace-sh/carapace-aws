package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Lists the jobs for a branch of an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_listJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_listJobsCmd).Standalone()

		amplify_listJobsCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_listJobsCmd.Flags().String("branch-name", "", "The name of the branch to use for the request.")
		amplify_listJobsCmd.Flags().String("max-results", "", "The maximum number of records to list in a single response.")
		amplify_listJobsCmd.Flags().String("next-token", "", "A pagination token.")
		amplify_listJobsCmd.MarkFlagRequired("app-id")
		amplify_listJobsCmd.MarkFlagRequired("branch-name")
	})
	amplifyCmd.AddCommand(amplify_listJobsCmd)
}
