package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_disassociateConnectionFromLagCmd = &cobra.Command{
	Use:   "disassociate-connection-from-lag",
	Short: "Disassociates a connection from a link aggregation group (LAG).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_disassociateConnectionFromLagCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_disassociateConnectionFromLagCmd).Standalone()

		directconnect_disassociateConnectionFromLagCmd.Flags().String("connection-id", "", "The ID of the connection.")
		directconnect_disassociateConnectionFromLagCmd.Flags().String("lag-id", "", "The ID of the LAG.")
		directconnect_disassociateConnectionFromLagCmd.MarkFlagRequired("connection-id")
		directconnect_disassociateConnectionFromLagCmd.MarkFlagRequired("lag-id")
	})
	directconnectCmd.AddCommand(directconnect_disassociateConnectionFromLagCmd)
}
