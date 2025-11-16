package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_batchUnsuspendUserCmd = &cobra.Command{
	Use:   "batch-unsuspend-user",
	Short: "Removes the suspension from up to 50 previously suspended users for the specified Amazon Chime `EnterpriseLWA` account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_batchUnsuspendUserCmd).Standalone()

	chime_batchUnsuspendUserCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_batchUnsuspendUserCmd.Flags().String("user-id-list", "", "The request containing the user IDs to unsuspend.")
	chime_batchUnsuspendUserCmd.MarkFlagRequired("account-id")
	chime_batchUnsuspendUserCmd.MarkFlagRequired("user-id-list")
	chimeCmd.AddCommand(chime_batchUnsuspendUserCmd)
}
