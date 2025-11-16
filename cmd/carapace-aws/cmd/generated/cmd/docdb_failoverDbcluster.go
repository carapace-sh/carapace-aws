package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_failoverDbclusterCmd = &cobra.Command{
	Use:   "failover-dbcluster",
	Short: "Forces a failover for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_failoverDbclusterCmd).Standalone()

	docdb_failoverDbclusterCmd.Flags().String("dbcluster-identifier", "", "A cluster identifier to force a failover for.")
	docdb_failoverDbclusterCmd.Flags().String("target-dbinstance-identifier", "", "The name of the instance to promote to the primary instance.")
	docdbCmd.AddCommand(docdb_failoverDbclusterCmd)
}
