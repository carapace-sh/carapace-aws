package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_updateConnectionCmd = &cobra.Command{
	Use:   "update-connection",
	Short: "Updates the Direct Connect connection configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_updateConnectionCmd).Standalone()

	directconnect_updateConnectionCmd.Flags().String("connection-id", "", "The ID of the connection.")
	directconnect_updateConnectionCmd.Flags().String("connection-name", "", "The name of the connection.")
	directconnect_updateConnectionCmd.Flags().String("encryption-mode", "", "The connection MAC Security (MACsec) encryption mode.")
	directconnect_updateConnectionCmd.MarkFlagRequired("connection-id")
	directconnectCmd.AddCommand(directconnect_updateConnectionCmd)
}
