package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateQuickConnectConfigCmd = &cobra.Command{
	Use:   "update-quick-connect-config",
	Short: "Updates the configuration settings for the specified quick connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateQuickConnectConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateQuickConnectConfigCmd).Standalone()

		connect_updateQuickConnectConfigCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateQuickConnectConfigCmd.Flags().String("quick-connect-config", "", "Information about the configuration settings for the quick connect.")
		connect_updateQuickConnectConfigCmd.Flags().String("quick-connect-id", "", "The identifier for the quick connect.")
		connect_updateQuickConnectConfigCmd.MarkFlagRequired("instance-id")
		connect_updateQuickConnectConfigCmd.MarkFlagRequired("quick-connect-config")
		connect_updateQuickConnectConfigCmd.MarkFlagRequired("quick-connect-id")
	})
	connectCmd.AddCommand(connect_updateQuickConnectConfigCmd)
}
