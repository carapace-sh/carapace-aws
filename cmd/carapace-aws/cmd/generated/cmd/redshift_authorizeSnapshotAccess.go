package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_authorizeSnapshotAccessCmd = &cobra.Command{
	Use:   "authorize-snapshot-access",
	Short: "Authorizes the specified Amazon Web Services account to restore the specified snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_authorizeSnapshotAccessCmd).Standalone()

	redshift_authorizeSnapshotAccessCmd.Flags().String("account-with-restore-access", "", "The identifier of the Amazon Web Services account authorized to restore the specified snapshot.")
	redshift_authorizeSnapshotAccessCmd.Flags().String("snapshot-arn", "", "The Amazon Resource Name (ARN) of the snapshot to authorize access to.")
	redshift_authorizeSnapshotAccessCmd.Flags().String("snapshot-cluster-identifier", "", "The identifier of the cluster the snapshot was created from.")
	redshift_authorizeSnapshotAccessCmd.Flags().String("snapshot-identifier", "", "The identifier of the snapshot the account is authorized to restore.")
	redshift_authorizeSnapshotAccessCmd.MarkFlagRequired("account-with-restore-access")
	redshiftCmd.AddCommand(redshift_authorizeSnapshotAccessCmd)
}
