package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createQuickConnectCmd = &cobra.Command{
	Use:   "create-quick-connect",
	Short: "Creates a quick connect for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createQuickConnectCmd).Standalone()

	connect_createQuickConnectCmd.Flags().String("description", "", "The description of the quick connect.")
	connect_createQuickConnectCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createQuickConnectCmd.Flags().String("name", "", "A unique name of the quick connect.")
	connect_createQuickConnectCmd.Flags().String("quick-connect-config", "", "Configuration settings for the quick connect.")
	connect_createQuickConnectCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createQuickConnectCmd.MarkFlagRequired("instance-id")
	connect_createQuickConnectCmd.MarkFlagRequired("name")
	connect_createQuickConnectCmd.MarkFlagRequired("quick-connect-config")
	connectCmd.AddCommand(connect_createQuickConnectCmd)
}
