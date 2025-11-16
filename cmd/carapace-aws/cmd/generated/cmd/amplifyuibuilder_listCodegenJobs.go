package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_listCodegenJobsCmd = &cobra.Command{
	Use:   "list-codegen-jobs",
	Short: "Retrieves a list of code generation jobs for a specified Amplify app and backend environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_listCodegenJobsCmd).Standalone()

	amplifyuibuilder_listCodegenJobsCmd.Flags().String("app-id", "", "The unique ID for the Amplify app.")
	amplifyuibuilder_listCodegenJobsCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
	amplifyuibuilder_listCodegenJobsCmd.Flags().String("max-results", "", "The maximum number of jobs to retrieve.")
	amplifyuibuilder_listCodegenJobsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	amplifyuibuilder_listCodegenJobsCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_listCodegenJobsCmd.MarkFlagRequired("environment-name")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_listCodegenJobsCmd)
}
