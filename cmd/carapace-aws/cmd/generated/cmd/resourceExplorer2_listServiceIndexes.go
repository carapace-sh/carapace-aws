package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listServiceIndexesCmd = &cobra.Command{
	Use:   "list-service-indexes",
	Short: "Lists all Resource Explorer indexes across the specified Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listServiceIndexesCmd).Standalone()

	resourceExplorer2_listServiceIndexesCmd.Flags().String("max-results", "", "The maximum number of index results to return in a single response.")
	resourceExplorer2_listServiceIndexesCmd.Flags().String("next-token", "", "The pagination token from a previous `ListServiceIndexes` response.")
	resourceExplorer2_listServiceIndexesCmd.Flags().String("regions", "", "A list of Amazon Web Services Regions to include in the search for indexes.")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listServiceIndexesCmd)
}
