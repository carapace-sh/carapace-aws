package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_disconnectRecoveryInstanceCmd = &cobra.Command{
	Use:   "disconnect-recovery-instance",
	Short: "Disconnect a Recovery Instance from Elastic Disaster Recovery.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_disconnectRecoveryInstanceCmd).Standalone()

	drs_disconnectRecoveryInstanceCmd.Flags().String("recovery-instance-id", "", "The ID of the Recovery Instance to disconnect.")
	drs_disconnectRecoveryInstanceCmd.MarkFlagRequired("recovery-instance-id")
	drsCmd.AddCommand(drs_disconnectRecoveryInstanceCmd)
}
