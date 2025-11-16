package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_updateStorageLensGroupCmd = &cobra.Command{
	Use:   "update-storage-lens-group",
	Short: "Updates the existing Storage Lens group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_updateStorageLensGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_updateStorageLensGroupCmd).Standalone()

		s3control_updateStorageLensGroupCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Storage Lens group owner.")
		s3control_updateStorageLensGroupCmd.Flags().String("name", "", "The name of the Storage Lens group that you want to update.")
		s3control_updateStorageLensGroupCmd.Flags().String("storage-lens-group", "", "The JSON file that contains the Storage Lens group configuration.")
		s3control_updateStorageLensGroupCmd.MarkFlagRequired("account-id")
		s3control_updateStorageLensGroupCmd.MarkFlagRequired("name")
		s3control_updateStorageLensGroupCmd.MarkFlagRequired("storage-lens-group")
	})
	s3controlCmd.AddCommand(s3control_updateStorageLensGroupCmd)
}
