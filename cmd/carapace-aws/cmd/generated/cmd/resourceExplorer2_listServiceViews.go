package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listServiceViewsCmd = &cobra.Command{
	Use:   "list-service-views",
	Short: "Lists all Resource Explorer service views available in the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listServiceViewsCmd).Standalone()

	resourceExplorer2_listServiceViewsCmd.Flags().String("max-results", "", "The maximum number of service view results to return in a single response.")
	resourceExplorer2_listServiceViewsCmd.Flags().String("next-token", "", "The pagination token from a previous `ListServiceViews` response.")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listServiceViewsCmd)
}
