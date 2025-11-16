package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_startSourceNetworkReplicationCmd = &cobra.Command{
	Use:   "start-source-network-replication",
	Short: "Starts replication for a Source Network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_startSourceNetworkReplicationCmd).Standalone()

	drs_startSourceNetworkReplicationCmd.Flags().String("source-network-id", "", "ID of the Source Network to replicate.")
	drs_startSourceNetworkReplicationCmd.MarkFlagRequired("source-network-id")
	drsCmd.AddCommand(drs_startSourceNetworkReplicationCmd)
}
