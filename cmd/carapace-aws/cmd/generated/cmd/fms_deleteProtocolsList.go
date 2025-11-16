package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_deleteProtocolsListCmd = &cobra.Command{
	Use:   "delete-protocols-list",
	Short: "Permanently deletes an Firewall Manager protocols list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_deleteProtocolsListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_deleteProtocolsListCmd).Standalone()

		fms_deleteProtocolsListCmd.Flags().String("list-id", "", "The ID of the protocols list that you want to delete.")
		fms_deleteProtocolsListCmd.MarkFlagRequired("list-id")
	})
	fmsCmd.AddCommand(fms_deleteProtocolsListCmd)
}
