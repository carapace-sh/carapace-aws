package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_promoteReadReplicaCmd = &cobra.Command{
	Use:   "promote-read-replica",
	Short: "Promotes a read replica DB instance to a standalone DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_promoteReadReplicaCmd).Standalone()

	rds_promoteReadReplicaCmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups are retained.")
	rds_promoteReadReplicaCmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier.")
	rds_promoteReadReplicaCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled, using the `BackupRetentionPeriod` parameter.")
	rds_promoteReadReplicaCmd.MarkFlagRequired("dbinstance-identifier")
	rdsCmd.AddCommand(rds_promoteReadReplicaCmd)
}
