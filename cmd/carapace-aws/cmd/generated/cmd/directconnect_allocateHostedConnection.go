package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_allocateHostedConnectionCmd = &cobra.Command{
	Use:   "allocate-hosted-connection",
	Short: "Creates a hosted connection on the specified interconnect or a link aggregation group (LAG) of interconnects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_allocateHostedConnectionCmd).Standalone()

	directconnect_allocateHostedConnectionCmd.Flags().String("bandwidth", "", "The bandwidth of the connection.")
	directconnect_allocateHostedConnectionCmd.Flags().String("connection-id", "", "The ID of the interconnect or LAG.")
	directconnect_allocateHostedConnectionCmd.Flags().String("connection-name", "", "The name of the hosted connection.")
	directconnect_allocateHostedConnectionCmd.Flags().String("owner-account", "", "The ID of the Amazon Web Services account ID of the customer for the connection.")
	directconnect_allocateHostedConnectionCmd.Flags().String("tags", "", "The tags associated with the connection.")
	directconnect_allocateHostedConnectionCmd.Flags().String("vlan", "", "The dedicated VLAN provisioned to the hosted connection.")
	directconnect_allocateHostedConnectionCmd.MarkFlagRequired("bandwidth")
	directconnect_allocateHostedConnectionCmd.MarkFlagRequired("connection-id")
	directconnect_allocateHostedConnectionCmd.MarkFlagRequired("connection-name")
	directconnect_allocateHostedConnectionCmd.MarkFlagRequired("owner-account")
	directconnect_allocateHostedConnectionCmd.MarkFlagRequired("vlan")
	directconnectCmd.AddCommand(directconnect_allocateHostedConnectionCmd)
}
