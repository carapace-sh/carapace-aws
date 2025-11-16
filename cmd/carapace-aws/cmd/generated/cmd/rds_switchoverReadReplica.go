package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_switchoverReadReplicaCmd = &cobra.Command{
	Use:   "switchover-read-replica",
	Short: "Switches over an Oracle standby database in an Oracle Data Guard environment, making it the new primary database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_switchoverReadReplicaCmd).Standalone()

	rds_switchoverReadReplicaCmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier of the current standby database.")
	rds_switchoverReadReplicaCmd.MarkFlagRequired("dbinstance-identifier")
	rdsCmd.AddCommand(rds_switchoverReadReplicaCmd)
}
