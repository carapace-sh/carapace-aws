package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getRecoveryPointCmd = &cobra.Command{
	Use:   "get-recovery-point",
	Short: "Returns information about a recovery point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getRecoveryPointCmd).Standalone()

	redshiftServerless_getRecoveryPointCmd.Flags().String("recovery-point-id", "", "The unique identifier of the recovery point to return information for.")
	redshiftServerless_getRecoveryPointCmd.MarkFlagRequired("recovery-point-id")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getRecoveryPointCmd)
}
