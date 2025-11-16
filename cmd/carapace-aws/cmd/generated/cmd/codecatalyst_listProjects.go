package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "Retrieves a list of projects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listProjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_listProjectsCmd).Standalone()

		codecatalyst_listProjectsCmd.Flags().String("filters", "", "Information about filters to apply to narrow the results returned in the list.")
		codecatalyst_listProjectsCmd.Flags().String("max-results", "", "The maximum number of results to show in a single call to this API.")
		codecatalyst_listProjectsCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
		codecatalyst_listProjectsCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_listProjectsCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_listProjectsCmd)
}
