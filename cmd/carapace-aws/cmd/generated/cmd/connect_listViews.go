package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listViewsCmd = &cobra.Command{
	Use:   "list-views",
	Short: "Returns views in the given instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listViewsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listViewsCmd).Standalone()

		connect_listViewsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listViewsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listViewsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listViewsCmd.Flags().String("type", "", "The type of the view.")
		connect_listViewsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listViewsCmd)
}
