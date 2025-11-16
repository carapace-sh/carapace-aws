package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_decreaseReplicationFactorCmd = &cobra.Command{
	Use:   "decrease-replication-factor",
	Short: "Removes one or more nodes from a DAX cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_decreaseReplicationFactorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_decreaseReplicationFactorCmd).Standalone()

		dax_decreaseReplicationFactorCmd.Flags().String("availability-zones", "", "The Availability Zone(s) from which to remove nodes.")
		dax_decreaseReplicationFactorCmd.Flags().String("cluster-name", "", "The name of the DAX cluster from which you want to remove nodes.")
		dax_decreaseReplicationFactorCmd.Flags().String("new-replication-factor", "", "The new number of nodes for the DAX cluster.")
		dax_decreaseReplicationFactorCmd.Flags().String("node-ids-to-remove", "", "The unique identifiers of the nodes to be removed from the cluster.")
		dax_decreaseReplicationFactorCmd.MarkFlagRequired("cluster-name")
		dax_decreaseReplicationFactorCmd.MarkFlagRequired("new-replication-factor")
	})
	daxCmd.AddCommand(dax_decreaseReplicationFactorCmd)
}
