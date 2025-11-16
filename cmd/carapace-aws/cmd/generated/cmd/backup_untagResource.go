package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a set of key-value pairs from a recovery point, backup plan, or backup vault identified by an Amazon Resource Name (ARN)\n\nThis API is not supported for recovery points for resource types including Aurora, Amazon DocumentDB.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_untagResourceCmd).Standalone()

		backup_untagResourceCmd.Flags().String("resource-arn", "", "An ARN that uniquely identifies a resource.")
		backup_untagResourceCmd.Flags().String("tag-key-list", "", "The keys to identify which key-value tags to remove from a resource.")
		backup_untagResourceCmd.MarkFlagRequired("resource-arn")
		backup_untagResourceCmd.MarkFlagRequired("tag-key-list")
	})
	backupCmd.AddCommand(backup_untagResourceCmd)
}
