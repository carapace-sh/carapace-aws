package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Creates a connection between a customer network and a specific Direct Connect location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createConnectionCmd).Standalone()

	directconnect_createConnectionCmd.Flags().String("bandwidth", "", "The bandwidth of the connection.")
	directconnect_createConnectionCmd.Flags().String("connection-name", "", "The name of the connection.")
	directconnect_createConnectionCmd.Flags().String("lag-id", "", "The ID of the LAG.")
	directconnect_createConnectionCmd.Flags().String("location", "", "The location of the connection.")
	directconnect_createConnectionCmd.Flags().String("provider-name", "", "The name of the service provider associated with the requested connection.")
	directconnect_createConnectionCmd.Flags().String("request-macsec", "", "Indicates whether you want the connection to support MAC Security (MACsec).")
	directconnect_createConnectionCmd.Flags().String("tags", "", "The tags to associate with the lag.")
	directconnect_createConnectionCmd.MarkFlagRequired("bandwidth")
	directconnect_createConnectionCmd.MarkFlagRequired("connection-name")
	directconnect_createConnectionCmd.MarkFlagRequired("location")
	directconnectCmd.AddCommand(directconnect_createConnectionCmd)
}
