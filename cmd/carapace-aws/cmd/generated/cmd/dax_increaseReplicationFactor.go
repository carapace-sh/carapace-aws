package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_increaseReplicationFactorCmd = &cobra.Command{
	Use:   "increase-replication-factor",
	Short: "Adds one or more nodes to a DAX cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_increaseReplicationFactorCmd).Standalone()

	dax_increaseReplicationFactorCmd.Flags().String("availability-zones", "", "The Availability Zones (AZs) in which the cluster nodes will be created.")
	dax_increaseReplicationFactorCmd.Flags().String("cluster-name", "", "The name of the DAX cluster that will receive additional nodes.")
	dax_increaseReplicationFactorCmd.Flags().String("new-replication-factor", "", "The new number of nodes for the DAX cluster.")
	dax_increaseReplicationFactorCmd.MarkFlagRequired("cluster-name")
	dax_increaseReplicationFactorCmd.MarkFlagRequired("new-replication-factor")
	daxCmd.AddCommand(dax_increaseReplicationFactorCmd)
}
