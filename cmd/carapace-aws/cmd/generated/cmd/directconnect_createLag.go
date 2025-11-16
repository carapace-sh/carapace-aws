package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createLagCmd = &cobra.Command{
	Use:   "create-lag",
	Short: "Creates a link aggregation group (LAG) with the specified number of bundled physical dedicated connections between the customer network and a specific Direct Connect location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createLagCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_createLagCmd).Standalone()

		directconnect_createLagCmd.Flags().String("child-connection-tags", "", "The tags to associate with the automtically created LAGs.")
		directconnect_createLagCmd.Flags().String("connection-id", "", "The ID of an existing dedicated connection to migrate to the LAG.")
		directconnect_createLagCmd.Flags().String("connections-bandwidth", "", "The bandwidth of the individual physical dedicated connections bundled by the LAG.")
		directconnect_createLagCmd.Flags().String("lag-name", "", "The name of the LAG.")
		directconnect_createLagCmd.Flags().String("location", "", "The location for the LAG.")
		directconnect_createLagCmd.Flags().String("number-of-connections", "", "The number of physical dedicated connections initially provisioned and bundled by the LAG.")
		directconnect_createLagCmd.Flags().String("provider-name", "", "The name of the service provider associated with the LAG.")
		directconnect_createLagCmd.Flags().String("request-macsec", "", "Indicates whether the connection will support MAC Security (MACsec).")
		directconnect_createLagCmd.Flags().String("tags", "", "The tags to associate with the LAG.")
		directconnect_createLagCmd.MarkFlagRequired("connections-bandwidth")
		directconnect_createLagCmd.MarkFlagRequired("lag-name")
		directconnect_createLagCmd.MarkFlagRequired("location")
		directconnect_createLagCmd.MarkFlagRequired("number-of-connections")
	})
	directconnectCmd.AddCommand(directconnect_createLagCmd)
}
