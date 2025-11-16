package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_updateLagCmd = &cobra.Command{
	Use:   "update-lag",
	Short: "Updates the attributes of the specified link aggregation group (LAG).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_updateLagCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_updateLagCmd).Standalone()

		directconnect_updateLagCmd.Flags().String("encryption-mode", "", "The LAG MAC Security (MACsec) encryption mode.")
		directconnect_updateLagCmd.Flags().String("lag-id", "", "The ID of the LAG.")
		directconnect_updateLagCmd.Flags().String("lag-name", "", "The name of the LAG.")
		directconnect_updateLagCmd.Flags().String("minimum-links", "", "The minimum number of physical connections that must be operational for the LAG itself to be operational.")
		directconnect_updateLagCmd.MarkFlagRequired("lag-id")
	})
	directconnectCmd.AddCommand(directconnect_updateLagCmd)
}
