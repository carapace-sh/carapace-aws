package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_startDbinstanceCmd = &cobra.Command{
	Use:   "start-dbinstance",
	Short: "Starts an Amazon RDS DB instance that was stopped using the Amazon Web Services console, the stop-db-instance CLI command, or the `StopDBInstance` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_startDbinstanceCmd).Standalone()

	rds_startDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The user-supplied instance identifier.")
	rds_startDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	rdsCmd.AddCommand(rds_startDbinstanceCmd)
}
