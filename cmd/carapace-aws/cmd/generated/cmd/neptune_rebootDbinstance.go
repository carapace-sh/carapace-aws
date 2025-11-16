package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_rebootDbinstanceCmd = &cobra.Command{
	Use:   "reboot-dbinstance",
	Short: "You might need to reboot your DB instance, usually for maintenance reasons.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_rebootDbinstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_rebootDbinstanceCmd).Standalone()

		neptune_rebootDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier.")
		neptune_rebootDbinstanceCmd.Flags().String("force-failover", "", "When `true`, the reboot is conducted through a MultiAZ failover.")
		neptune_rebootDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	})
	neptuneCmd.AddCommand(neptune_rebootDbinstanceCmd)
}
