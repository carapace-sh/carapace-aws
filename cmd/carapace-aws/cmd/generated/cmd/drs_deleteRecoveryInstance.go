package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_deleteRecoveryInstanceCmd = &cobra.Command{
	Use:   "delete-recovery-instance",
	Short: "Deletes a single Recovery Instance by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_deleteRecoveryInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_deleteRecoveryInstanceCmd).Standalone()

		drs_deleteRecoveryInstanceCmd.Flags().String("recovery-instance-id", "", "The ID of the Recovery Instance to be deleted.")
		drs_deleteRecoveryInstanceCmd.MarkFlagRequired("recovery-instance-id")
	})
	drsCmd.AddCommand(drs_deleteRecoveryInstanceCmd)
}
