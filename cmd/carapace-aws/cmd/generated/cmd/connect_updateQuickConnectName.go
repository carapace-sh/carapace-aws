package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateQuickConnectNameCmd = &cobra.Command{
	Use:   "update-quick-connect-name",
	Short: "Updates the name and description of a quick connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateQuickConnectNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateQuickConnectNameCmd).Standalone()

		connect_updateQuickConnectNameCmd.Flags().String("description", "", "The description of the quick connect.")
		connect_updateQuickConnectNameCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateQuickConnectNameCmd.Flags().String("name", "", "The name of the quick connect.")
		connect_updateQuickConnectNameCmd.Flags().String("quick-connect-id", "", "The identifier for the quick connect.")
		connect_updateQuickConnectNameCmd.MarkFlagRequired("instance-id")
		connect_updateQuickConnectNameCmd.MarkFlagRequired("quick-connect-id")
	})
	connectCmd.AddCommand(connect_updateQuickConnectNameCmd)
}
