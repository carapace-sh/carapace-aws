package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_disconnectFromServiceCmd = &cobra.Command{
	Use:   "disconnect-from-service",
	Short: "Disconnects specific Source Servers from Application Migration Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_disconnectFromServiceCmd).Standalone()

	mgn_disconnectFromServiceCmd.Flags().String("account-id", "", "Request to disconnect Source Server from service by Account ID.")
	mgn_disconnectFromServiceCmd.Flags().String("source-server-id", "", "Request to disconnect Source Server from service by Server ID.")
	mgn_disconnectFromServiceCmd.MarkFlagRequired("source-server-id")
	mgnCmd.AddCommand(mgn_disconnectFromServiceCmd)
}
