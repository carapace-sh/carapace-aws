package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_allocateConnectionOnInterconnectCmd = &cobra.Command{
	Use:   "allocate-connection-on-interconnect",
	Short: "Deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_allocateConnectionOnInterconnectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_allocateConnectionOnInterconnectCmd).Standalone()

		directconnect_allocateConnectionOnInterconnectCmd.Flags().String("bandwidth", "", "The bandwidth of the connection.")
		directconnect_allocateConnectionOnInterconnectCmd.Flags().String("connection-name", "", "The name of the provisioned connection.")
		directconnect_allocateConnectionOnInterconnectCmd.Flags().String("interconnect-id", "", "The ID of the interconnect on which the connection will be provisioned.")
		directconnect_allocateConnectionOnInterconnectCmd.Flags().String("owner-account", "", "The ID of the Amazon Web Services account of the customer for whom the connection will be provisioned.")
		directconnect_allocateConnectionOnInterconnectCmd.Flags().String("vlan", "", "The dedicated VLAN provisioned to the connection.")
		directconnect_allocateConnectionOnInterconnectCmd.MarkFlagRequired("bandwidth")
		directconnect_allocateConnectionOnInterconnectCmd.MarkFlagRequired("connection-name")
		directconnect_allocateConnectionOnInterconnectCmd.MarkFlagRequired("interconnect-id")
		directconnect_allocateConnectionOnInterconnectCmd.MarkFlagRequired("owner-account")
		directconnect_allocateConnectionOnInterconnectCmd.MarkFlagRequired("vlan")
	})
	directconnectCmd.AddCommand(directconnect_allocateConnectionOnInterconnectCmd)
}
