package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listIndexesCmd = &cobra.Command{
	Use:   "list-indexes",
	Short: "Retrieves a list of all of the indexes in Amazon Web Services Regions that are currently collecting resource information for Amazon Web Services Resource Explorer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listIndexesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_listIndexesCmd).Standalone()

		resourceExplorer2_listIndexesCmd.Flags().String("max-results", "", "The maximum number of results that you want included on each page of the response.")
		resourceExplorer2_listIndexesCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		resourceExplorer2_listIndexesCmd.Flags().String("regions", "", "If specified, limits the response to only information about the index in the specified list of Amazon Web Services Regions.")
		resourceExplorer2_listIndexesCmd.Flags().String("type", "", "If specified, limits the output to only indexes of the specified Type, either `LOCAL` or `AGGREGATOR`.")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listIndexesCmd)
}
