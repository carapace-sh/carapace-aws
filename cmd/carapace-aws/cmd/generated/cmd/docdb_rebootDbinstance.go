package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_rebootDbinstanceCmd = &cobra.Command{
	Use:   "reboot-dbinstance",
	Short: "You might need to reboot your instance, usually for maintenance reasons.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_rebootDbinstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_rebootDbinstanceCmd).Standalone()

		docdb_rebootDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The instance identifier.")
		docdb_rebootDbinstanceCmd.Flags().String("force-failover", "", "When `true`, the reboot is conducted through a Multi-AZ failover.")
		docdb_rebootDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	})
	docdbCmd.AddCommand(docdb_rebootDbinstanceCmd)
}
