package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteQuickConnectCmd = &cobra.Command{
	Use:   "delete-quick-connect",
	Short: "Deletes a quick connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteQuickConnectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteQuickConnectCmd).Standalone()

		connect_deleteQuickConnectCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteQuickConnectCmd.Flags().String("quick-connect-id", "", "The identifier for the quick connect.")
		connect_deleteQuickConnectCmd.MarkFlagRequired("instance-id")
		connect_deleteQuickConnectCmd.MarkFlagRequired("quick-connect-id")
	})
	connectCmd.AddCommand(connect_deleteQuickConnectCmd)
}
