package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_listArtifactsCmd = &cobra.Command{
	Use:   "list-artifacts",
	Short: "Returns a list of end-to-end testing artifacts for a specified app, branch, and job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_listArtifactsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_listArtifactsCmd).Standalone()

		amplify_listArtifactsCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_listArtifactsCmd.Flags().String("branch-name", "", "The name of a branch that is part of an Amplify app.")
		amplify_listArtifactsCmd.Flags().String("job-id", "", "The unique ID for a job.")
		amplify_listArtifactsCmd.Flags().String("max-results", "", "The maximum number of records to list in a single response.")
		amplify_listArtifactsCmd.Flags().String("next-token", "", "A pagination token.")
		amplify_listArtifactsCmd.MarkFlagRequired("app-id")
		amplify_listArtifactsCmd.MarkFlagRequired("branch-name")
		amplify_listArtifactsCmd.MarkFlagRequired("job-id")
	})
	amplifyCmd.AddCommand(amplify_listArtifactsCmd)
}
