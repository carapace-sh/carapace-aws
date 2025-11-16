package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_failoverDbclusterCmd = &cobra.Command{
	Use:   "failover-dbcluster",
	Short: "Forces a failover for a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_failoverDbclusterCmd).Standalone()

	neptune_failoverDbclusterCmd.Flags().String("dbcluster-identifier", "", "A DB cluster identifier to force a failover for.")
	neptune_failoverDbclusterCmd.Flags().String("target-dbinstance-identifier", "", "The name of the instance to promote to the primary instance.")
	neptuneCmd.AddCommand(neptune_failoverDbclusterCmd)
}
