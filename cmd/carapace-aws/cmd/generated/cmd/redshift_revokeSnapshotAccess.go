package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_revokeSnapshotAccessCmd = &cobra.Command{
	Use:   "revoke-snapshot-access",
	Short: "Removes the ability of the specified Amazon Web Services account to restore the specified snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_revokeSnapshotAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_revokeSnapshotAccessCmd).Standalone()

		redshift_revokeSnapshotAccessCmd.Flags().String("account-with-restore-access", "", "The identifier of the Amazon Web Services account that can no longer restore the specified snapshot.")
		redshift_revokeSnapshotAccessCmd.Flags().String("snapshot-arn", "", "The Amazon Resource Name (ARN) of the snapshot associated with the message to revoke access.")
		redshift_revokeSnapshotAccessCmd.Flags().String("snapshot-cluster-identifier", "", "The identifier of the cluster the snapshot was created from.")
		redshift_revokeSnapshotAccessCmd.Flags().String("snapshot-identifier", "", "The identifier of the snapshot that the account can no longer access.")
		redshift_revokeSnapshotAccessCmd.MarkFlagRequired("account-with-restore-access")
	})
	redshiftCmd.AddCommand(redshift_revokeSnapshotAccessCmd)
}
