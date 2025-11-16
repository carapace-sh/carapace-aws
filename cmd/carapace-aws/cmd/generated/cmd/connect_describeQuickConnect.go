package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeQuickConnectCmd = &cobra.Command{
	Use:   "describe-quick-connect",
	Short: "Describes the quick connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeQuickConnectCmd).Standalone()

	connect_describeQuickConnectCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeQuickConnectCmd.Flags().String("quick-connect-id", "", "The identifier for the quick connect.")
	connect_describeQuickConnectCmd.MarkFlagRequired("instance-id")
	connect_describeQuickConnectCmd.MarkFlagRequired("quick-connect-id")
	connectCmd.AddCommand(connect_describeQuickConnectCmd)
}
