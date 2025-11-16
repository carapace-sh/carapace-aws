package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_stopDbinstanceAutomatedBackupsReplicationCmd = &cobra.Command{
	Use:   "stop-dbinstance-automated-backups-replication",
	Short: "Stops automated backup replication for a DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_stopDbinstanceAutomatedBackupsReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_stopDbinstanceAutomatedBackupsReplicationCmd).Standalone()

		rds_stopDbinstanceAutomatedBackupsReplicationCmd.Flags().String("source-dbinstance-arn", "", "The Amazon Resource Name (ARN) of the source DB instance for which to stop replicating automate backups, for example, `arn:aws:rds:us-west-2:123456789012:db:mydatabase`.")
		rds_stopDbinstanceAutomatedBackupsReplicationCmd.MarkFlagRequired("source-dbinstance-arn")
	})
	rdsCmd.AddCommand(rds_stopDbinstanceAutomatedBackupsReplicationCmd)
}
