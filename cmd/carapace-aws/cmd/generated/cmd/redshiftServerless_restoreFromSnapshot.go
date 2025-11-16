package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_restoreFromSnapshotCmd = &cobra.Command{
	Use:   "restore-from-snapshot",
	Short: "Restores a namespace from a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_restoreFromSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_restoreFromSnapshotCmd).Standalone()

		redshiftServerless_restoreFromSnapshotCmd.Flags().String("admin-password-secret-kms-key-id", "", "The ID of the Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret.")
		redshiftServerless_restoreFromSnapshotCmd.Flags().Bool("manage-admin-password", false, "If `true`, Amazon Redshift uses Secrets Manager to manage the restored snapshot's admin credentials.")
		redshiftServerless_restoreFromSnapshotCmd.Flags().String("namespace-name", "", "The name of the namespace to restore the snapshot to.")
		redshiftServerless_restoreFromSnapshotCmd.Flags().Bool("no-manage-admin-password", false, "If `true`, Amazon Redshift uses Secrets Manager to manage the restored snapshot's admin credentials.")
		redshiftServerless_restoreFromSnapshotCmd.Flags().String("owner-account", "", "The Amazon Web Services account that owns the snapshot.")
		redshiftServerless_restoreFromSnapshotCmd.Flags().String("snapshot-arn", "", "The Amazon Resource Name (ARN) of the snapshot to restore from.")
		redshiftServerless_restoreFromSnapshotCmd.Flags().String("snapshot-name", "", "The name of the snapshot to restore from.")
		redshiftServerless_restoreFromSnapshotCmd.Flags().String("workgroup-name", "", "The name of the workgroup used to restore the snapshot.")
		redshiftServerless_restoreFromSnapshotCmd.MarkFlagRequired("namespace-name")
		redshiftServerless_restoreFromSnapshotCmd.Flag("no-manage-admin-password").Hidden = true
		redshiftServerless_restoreFromSnapshotCmd.MarkFlagRequired("workgroup-name")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_restoreFromSnapshotCmd)
}
