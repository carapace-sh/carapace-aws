package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_stopSourceNetworkReplicationCmd = &cobra.Command{
	Use:   "stop-source-network-replication",
	Short: "Stops replication for a Source Network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_stopSourceNetworkReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_stopSourceNetworkReplicationCmd).Standalone()

		drs_stopSourceNetworkReplicationCmd.Flags().String("source-network-id", "", "ID of the Source Network to stop replication.")
		drs_stopSourceNetworkReplicationCmd.MarkFlagRequired("source-network-id")
	})
	drsCmd.AddCommand(drs_stopSourceNetworkReplicationCmd)
}
