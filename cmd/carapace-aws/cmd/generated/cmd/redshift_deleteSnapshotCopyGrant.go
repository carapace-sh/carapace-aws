package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteSnapshotCopyGrantCmd = &cobra.Command{
	Use:   "delete-snapshot-copy-grant",
	Short: "Deletes the specified snapshot copy grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteSnapshotCopyGrantCmd).Standalone()

	redshift_deleteSnapshotCopyGrantCmd.Flags().String("snapshot-copy-grant-name", "", "The name of the snapshot copy grant to delete.")
	redshift_deleteSnapshotCopyGrantCmd.MarkFlagRequired("snapshot-copy-grant-name")
	redshiftCmd.AddCommand(redshift_deleteSnapshotCopyGrantCmd)
}
