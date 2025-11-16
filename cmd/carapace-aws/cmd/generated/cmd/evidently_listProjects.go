package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_listProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "Returns configuration details about all the projects in the current Region in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_listProjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_listProjectsCmd).Standalone()

		evidently_listProjectsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		evidently_listProjectsCmd.Flags().String("next-token", "", "The token to use when requesting the next set of results.")
	})
	evidentlyCmd.AddCommand(evidently_listProjectsCmd)
}
