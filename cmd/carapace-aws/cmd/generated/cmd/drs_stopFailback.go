package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_stopFailbackCmd = &cobra.Command{
	Use:   "stop-failback",
	Short: "Stops the failback process for a specified Recovery Instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_stopFailbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_stopFailbackCmd).Standalone()

		drs_stopFailbackCmd.Flags().String("recovery-instance-id", "", "The ID of the Recovery Instance we want to stop failback for.")
		drs_stopFailbackCmd.MarkFlagRequired("recovery-instance-id")
	})
	drsCmd.AddCommand(drs_stopFailbackCmd)
}
