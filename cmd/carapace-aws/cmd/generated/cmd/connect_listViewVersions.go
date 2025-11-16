package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listViewVersionsCmd = &cobra.Command{
	Use:   "list-view-versions",
	Short: "Returns all the available versions for the specified Amazon Connect instance and view identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listViewVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listViewVersionsCmd).Standalone()

		connect_listViewVersionsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listViewVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listViewVersionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listViewVersionsCmd.Flags().String("view-id", "", "The identifier of the view.")
		connect_listViewVersionsCmd.MarkFlagRequired("instance-id")
		connect_listViewVersionsCmd.MarkFlagRequired("view-id")
	})
	connectCmd.AddCommand(connect_listViewVersionsCmd)
}
