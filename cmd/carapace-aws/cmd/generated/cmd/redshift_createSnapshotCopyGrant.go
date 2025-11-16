package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createSnapshotCopyGrantCmd = &cobra.Command{
	Use:   "create-snapshot-copy-grant",
	Short: "Creates a snapshot copy grant that permits Amazon Redshift to use an encrypted symmetric key from Key Management Service (KMS) to encrypt copied snapshots in a destination region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createSnapshotCopyGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_createSnapshotCopyGrantCmd).Standalone()

		redshift_createSnapshotCopyGrantCmd.Flags().String("kms-key-id", "", "The unique identifier of the encrypted symmetric key to which to grant Amazon Redshift permission.")
		redshift_createSnapshotCopyGrantCmd.Flags().String("snapshot-copy-grant-name", "", "The name of the snapshot copy grant.")
		redshift_createSnapshotCopyGrantCmd.Flags().String("tags", "", "A list of tag instances.")
		redshift_createSnapshotCopyGrantCmd.MarkFlagRequired("snapshot-copy-grant-name")
	})
	redshiftCmd.AddCommand(redshift_createSnapshotCopyGrantCmd)
}
