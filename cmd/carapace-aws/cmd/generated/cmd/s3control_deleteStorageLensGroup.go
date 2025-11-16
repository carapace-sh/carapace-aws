package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteStorageLensGroupCmd = &cobra.Command{
	Use:   "delete-storage-lens-group",
	Short: "Deletes an existing S3 Storage Lens group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteStorageLensGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deleteStorageLensGroupCmd).Standalone()

		s3control_deleteStorageLensGroupCmd.Flags().String("account-id", "", "The Amazon Web Services account ID used to create the Storage Lens group that you're trying to delete.")
		s3control_deleteStorageLensGroupCmd.Flags().String("name", "", "The name of the Storage Lens group that you're trying to delete.")
		s3control_deleteStorageLensGroupCmd.MarkFlagRequired("account-id")
		s3control_deleteStorageLensGroupCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_deleteStorageLensGroupCmd)
}
