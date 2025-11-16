package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeConnectionsOnInterconnectCmd = &cobra.Command{
	Use:   "describe-connections-on-interconnect",
	Short: "Deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeConnectionsOnInterconnectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_describeConnectionsOnInterconnectCmd).Standalone()

		directconnect_describeConnectionsOnInterconnectCmd.Flags().String("interconnect-id", "", "The ID of the interconnect.")
		directconnect_describeConnectionsOnInterconnectCmd.MarkFlagRequired("interconnect-id")
	})
	directconnectCmd.AddCommand(directconnect_describeConnectionsOnInterconnectCmd)
}
