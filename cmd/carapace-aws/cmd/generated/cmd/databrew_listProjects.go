package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_listProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "Lists all of the DataBrew projects that are defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_listProjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_listProjectsCmd).Standalone()

		databrew_listProjectsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
		databrew_listProjectsCmd.Flags().String("next-token", "", "The token returned by a previous call to retrieve the next set of results.")
	})
	databrewCmd.AddCommand(databrew_listProjectsCmd)
}
