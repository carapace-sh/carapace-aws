package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_associateConnectionWithLagCmd = &cobra.Command{
	Use:   "associate-connection-with-lag",
	Short: "Associates an existing connection with a link aggregation group (LAG).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_associateConnectionWithLagCmd).Standalone()

	directconnect_associateConnectionWithLagCmd.Flags().String("connection-id", "", "The ID of the connection.")
	directconnect_associateConnectionWithLagCmd.Flags().String("lag-id", "", "The ID of the LAG with which to associate the connection.")
	directconnect_associateConnectionWithLagCmd.MarkFlagRequired("connection-id")
	directconnect_associateConnectionWithLagCmd.MarkFlagRequired("lag-id")
	directconnectCmd.AddCommand(directconnect_associateConnectionWithLagCmd)
}
