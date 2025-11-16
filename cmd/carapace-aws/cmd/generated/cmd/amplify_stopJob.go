package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_stopJobCmd = &cobra.Command{
	Use:   "stop-job",
	Short: "Stops a job that is in progress for a branch of an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_stopJobCmd).Standalone()

	amplify_stopJobCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_stopJobCmd.Flags().String("branch-name", "", "The name of the branch to use for the stop job request.")
	amplify_stopJobCmd.Flags().String("job-id", "", "The unique id for the job.")
	amplify_stopJobCmd.MarkFlagRequired("app-id")
	amplify_stopJobCmd.MarkFlagRequired("branch-name")
	amplify_stopJobCmd.MarkFlagRequired("job-id")
	amplifyCmd.AddCommand(amplify_stopJobCmd)
}
