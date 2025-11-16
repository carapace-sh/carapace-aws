package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listRecoveryPointsByResourceCmd = &cobra.Command{
	Use:   "list-recovery-points-by-resource",
	Short: "The information about the recovery points of the type specified by a resource Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listRecoveryPointsByResourceCmd).Standalone()

	backup_listRecoveryPointsByResourceCmd.Flags().String("managed-by-awsbackup-only", "", "This attribute filters recovery points based on ownership.")
	backup_listRecoveryPointsByResourceCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
	backup_listRecoveryPointsByResourceCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backup_listRecoveryPointsByResourceCmd.Flags().String("resource-arn", "", "An ARN that uniquely identifies a resource.")
	backup_listRecoveryPointsByResourceCmd.MarkFlagRequired("resource-arn")
	backupCmd.AddCommand(backup_listRecoveryPointsByResourceCmd)
}
