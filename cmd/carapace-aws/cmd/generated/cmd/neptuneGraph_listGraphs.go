package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_listGraphsCmd = &cobra.Command{
	Use:   "list-graphs",
	Short: "Lists available Neptune Analytics graphs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_listGraphsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_listGraphsCmd).Standalone()

		neptuneGraph_listGraphsCmd.Flags().String("max-results", "", "The total number of records to return in the command's output.")
		neptuneGraph_listGraphsCmd.Flags().String("next-token", "", "Pagination token used to paginate output.")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_listGraphsCmd)
}
