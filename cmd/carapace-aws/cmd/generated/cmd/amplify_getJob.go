package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "Returns a job for a branch of an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_getJobCmd).Standalone()

	amplify_getJobCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_getJobCmd.Flags().String("branch-name", "", "The name of the branch to use for the job.")
	amplify_getJobCmd.Flags().String("job-id", "", "The unique ID for the job.")
	amplify_getJobCmd.MarkFlagRequired("app-id")
	amplify_getJobCmd.MarkFlagRequired("branch-name")
	amplify_getJobCmd.MarkFlagRequired("job-id")
	amplifyCmd.AddCommand(amplify_getJobCmd)
}
