package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_deleteJobCmd = &cobra.Command{
	Use:   "delete-job",
	Short: "Deletes a job for a branch of an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_deleteJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_deleteJobCmd).Standalone()

		amplify_deleteJobCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_deleteJobCmd.Flags().String("branch-name", "", "The name of the branch to use for the job.")
		amplify_deleteJobCmd.Flags().String("job-id", "", "The unique ID for the job.")
		amplify_deleteJobCmd.MarkFlagRequired("app-id")
		amplify_deleteJobCmd.MarkFlagRequired("branch-name")
		amplify_deleteJobCmd.MarkFlagRequired("job-id")
	})
	amplifyCmd.AddCommand(amplify_deleteJobCmd)
}
