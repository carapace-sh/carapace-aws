package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_restoreFromRecoveryPointCmd = &cobra.Command{
	Use:   "restore-from-recovery-point",
	Short: "Restore the data from a recovery point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_restoreFromRecoveryPointCmd).Standalone()

	redshiftServerless_restoreFromRecoveryPointCmd.Flags().String("namespace-name", "", "The name of the namespace to restore data into.")
	redshiftServerless_restoreFromRecoveryPointCmd.Flags().String("recovery-point-id", "", "The unique identifier of the recovery point to restore from.")
	redshiftServerless_restoreFromRecoveryPointCmd.Flags().String("workgroup-name", "", "The name of the workgroup used to restore data.")
	redshiftServerless_restoreFromRecoveryPointCmd.MarkFlagRequired("namespace-name")
	redshiftServerless_restoreFromRecoveryPointCmd.MarkFlagRequired("recovery-point-id")
	redshiftServerless_restoreFromRecoveryPointCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_restoreFromRecoveryPointCmd)
}
