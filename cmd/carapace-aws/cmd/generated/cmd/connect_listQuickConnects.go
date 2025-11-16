package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listQuickConnectsCmd = &cobra.Command{
	Use:   "list-quick-connects",
	Short: "Provides information about the quick connects for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listQuickConnectsCmd).Standalone()

	connect_listQuickConnectsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listQuickConnectsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listQuickConnectsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listQuickConnectsCmd.Flags().String("quick-connect-types", "", "The type of quick connect.")
	connect_listQuickConnectsCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listQuickConnectsCmd)
}
