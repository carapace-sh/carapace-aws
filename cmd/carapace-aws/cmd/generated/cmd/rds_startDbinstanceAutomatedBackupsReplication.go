package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_startDbinstanceAutomatedBackupsReplicationCmd = &cobra.Command{
	Use:   "start-dbinstance-automated-backups-replication",
	Short: "Enables replication of automated backups to a different Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_startDbinstanceAutomatedBackupsReplicationCmd).Standalone()

	rds_startDbinstanceAutomatedBackupsReplicationCmd.Flags().String("backup-retention-period", "", "The retention period for the replicated automated backups.")
	rds_startDbinstanceAutomatedBackupsReplicationCmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier for encryption of the replicated automated backups.")
	rds_startDbinstanceAutomatedBackupsReplicationCmd.Flags().String("pre-signed-url", "", "In an Amazon Web Services GovCloud (US) Region, an URL that contains a Signature Version 4 signed request for the `StartDBInstanceAutomatedBackupsReplication` operation to call in the Amazon Web Services Region of the source DB instance.")
	rds_startDbinstanceAutomatedBackupsReplicationCmd.Flags().String("source-dbinstance-arn", "", "The Amazon Resource Name (ARN) of the source DB instance for the replicated automated backups, for example, `arn:aws:rds:us-west-2:123456789012:db:mydatabase`.")
	rds_startDbinstanceAutomatedBackupsReplicationCmd.MarkFlagRequired("source-dbinstance-arn")
	rdsCmd.AddCommand(rds_startDbinstanceAutomatedBackupsReplicationCmd)
}
