package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createInterconnectCmd = &cobra.Command{
	Use:   "create-interconnect",
	Short: "Creates an interconnect between an Direct Connect Partner's network and a specific Direct Connect location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createInterconnectCmd).Standalone()

	directconnect_createInterconnectCmd.Flags().String("bandwidth", "", "The port bandwidth, in Gbps.")
	directconnect_createInterconnectCmd.Flags().String("interconnect-name", "", "The name of the interconnect.")
	directconnect_createInterconnectCmd.Flags().String("lag-id", "", "The ID of the LAG.")
	directconnect_createInterconnectCmd.Flags().String("location", "", "The location of the interconnect.")
	directconnect_createInterconnectCmd.Flags().String("provider-name", "", "The name of the service provider associated with the interconnect.")
	directconnect_createInterconnectCmd.Flags().String("request-macsec", "", "Indicates whether you want the interconnect to support MAC Security (MACsec).")
	directconnect_createInterconnectCmd.Flags().String("tags", "", "The tags to associate with the interconnect.")
	directconnect_createInterconnectCmd.MarkFlagRequired("bandwidth")
	directconnect_createInterconnectCmd.MarkFlagRequired("interconnect-name")
	directconnect_createInterconnectCmd.MarkFlagRequired("location")
	directconnectCmd.AddCommand(directconnect_createInterconnectCmd)
}
