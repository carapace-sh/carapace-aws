package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_listLaunchesCmd = &cobra.Command{
	Use:   "list-launches",
	Short: "Returns configuration details about all the launches in the specified project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_listLaunchesCmd).Standalone()

	evidently_listLaunchesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	evidently_listLaunchesCmd.Flags().String("next-token", "", "The token to use when requesting the next set of results.")
	evidently_listLaunchesCmd.Flags().String("project", "", "The name or ARN of the project to return the launch list from.")
	evidently_listLaunchesCmd.Flags().String("status", "", "Use this optional parameter to limit the returned results to only the launches with the status that you specify here.")
	evidently_listLaunchesCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_listLaunchesCmd)
}
