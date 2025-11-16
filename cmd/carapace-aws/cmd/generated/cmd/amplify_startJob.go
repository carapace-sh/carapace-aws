package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_startJobCmd = &cobra.Command{
	Use:   "start-job",
	Short: "Starts a new job for a branch of an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_startJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_startJobCmd).Standalone()

		amplify_startJobCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_startJobCmd.Flags().String("branch-name", "", "The name of the branch to use for the job.")
		amplify_startJobCmd.Flags().String("commit-id", "", "The commit ID from a third-party repository provider for the job.")
		amplify_startJobCmd.Flags().String("commit-message", "", "The commit message from a third-party repository provider for the job.")
		amplify_startJobCmd.Flags().String("commit-time", "", "The commit date and time for the job.")
		amplify_startJobCmd.Flags().String("job-id", "", "The unique ID for an existing job.")
		amplify_startJobCmd.Flags().String("job-reason", "", "A descriptive reason for starting the job.")
		amplify_startJobCmd.Flags().String("job-type", "", "Describes the type for the job.")
		amplify_startJobCmd.MarkFlagRequired("app-id")
		amplify_startJobCmd.MarkFlagRequired("branch-name")
		amplify_startJobCmd.MarkFlagRequired("job-type")
	})
	amplifyCmd.AddCommand(amplify_startJobCmd)
}
