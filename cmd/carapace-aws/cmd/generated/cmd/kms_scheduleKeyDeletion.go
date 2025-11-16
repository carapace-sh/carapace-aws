package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_scheduleKeyDeletionCmd = &cobra.Command{
	Use:   "schedule-key-deletion",
	Short: "Schedules the deletion of a KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_scheduleKeyDeletionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_scheduleKeyDeletionCmd).Standalone()

		kms_scheduleKeyDeletionCmd.Flags().String("key-id", "", "The unique identifier of the KMS key to delete.")
		kms_scheduleKeyDeletionCmd.Flags().String("pending-window-in-days", "", "The waiting period, specified in number of days.")
		kms_scheduleKeyDeletionCmd.MarkFlagRequired("key-id")
	})
	kmsCmd.AddCommand(kms_scheduleKeyDeletionCmd)
}
