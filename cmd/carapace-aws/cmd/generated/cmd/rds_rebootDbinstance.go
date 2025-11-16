package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_rebootDbinstanceCmd = &cobra.Command{
	Use:   "reboot-dbinstance",
	Short: "You might need to reboot your DB instance, usually for maintenance reasons.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_rebootDbinstanceCmd).Standalone()

	rds_rebootDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier.")
	rds_rebootDbinstanceCmd.Flags().String("force-failover", "", "Specifies whether the reboot is conducted through a Multi-AZ failover.")
	rds_rebootDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	rdsCmd.AddCommand(rds_rebootDbinstanceCmd)
}
