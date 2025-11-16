package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_listGraphsCmd = &cobra.Command{
	Use:   "list-graphs",
	Short: "Returns the list of behavior graphs that the calling account is an administrator account of.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_listGraphsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_listGraphsCmd).Standalone()

		detective_listGraphsCmd.Flags().String("max-results", "", "The maximum number of graphs to return at a time.")
		detective_listGraphsCmd.Flags().String("next-token", "", "For requests to get the next page of results, the pagination token that was returned with the previous set of results.")
	})
	detectiveCmd.AddCommand(detective_listGraphsCmd)
}
