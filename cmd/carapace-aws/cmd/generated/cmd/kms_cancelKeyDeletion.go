package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_cancelKeyDeletionCmd = &cobra.Command{
	Use:   "cancel-key-deletion",
	Short: "Cancels the deletion of a KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_cancelKeyDeletionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_cancelKeyDeletionCmd).Standalone()

		kms_cancelKeyDeletionCmd.Flags().String("key-id", "", "Identifies the KMS key whose deletion is being canceled.")
		kms_cancelKeyDeletionCmd.MarkFlagRequired("key-id")
	})
	kmsCmd.AddCommand(kms_cancelKeyDeletionCmd)
}
