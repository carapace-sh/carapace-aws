package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_createStorageLensGroupCmd = &cobra.Command{
	Use:   "create-storage-lens-group",
	Short: "Creates a new S3 Storage Lens group and associates it with the specified Amazon Web Services account ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_createStorageLensGroupCmd).Standalone()

	s3control_createStorageLensGroupCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that the Storage Lens group is created from and associated with.")
	s3control_createStorageLensGroupCmd.Flags().String("storage-lens-group", "", "The Storage Lens group configuration.")
	s3control_createStorageLensGroupCmd.Flags().String("tags", "", "The Amazon Web Services resource tags that you're adding to your Storage Lens group.")
	s3control_createStorageLensGroupCmd.MarkFlagRequired("account-id")
	s3control_createStorageLensGroupCmd.MarkFlagRequired("storage-lens-group")
	s3controlCmd.AddCommand(s3control_createStorageLensGroupCmd)
}
