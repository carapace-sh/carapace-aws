package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_failoverDbclusterCmd = &cobra.Command{
	Use:   "failover-dbcluster",
	Short: "Forces a failover for a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_failoverDbclusterCmd).Standalone()

	rds_failoverDbclusterCmd.Flags().String("dbcluster-identifier", "", "The identifier of the DB cluster to force a failover for.")
	rds_failoverDbclusterCmd.Flags().String("target-dbinstance-identifier", "", "The name of the DB instance to promote to the primary DB instance.")
	rds_failoverDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	rdsCmd.AddCommand(rds_failoverDbclusterCmd)
}
