package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_stopDbinstanceCmd = &cobra.Command{
	Use:   "stop-dbinstance",
	Short: "Stops an Amazon RDS DB instance temporarily.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_stopDbinstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_stopDbinstanceCmd).Standalone()

		rds_stopDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The user-supplied instance identifier.")
		rds_stopDbinstanceCmd.Flags().String("dbsnapshot-identifier", "", "The user-supplied instance identifier of the DB Snapshot created immediately before the DB instance is stopped.")
		rds_stopDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	})
	rdsCmd.AddCommand(rds_stopDbinstanceCmd)
}
