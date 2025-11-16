package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getStorageLensGroupCmd = &cobra.Command{
	Use:   "get-storage-lens-group",
	Short: "Retrieves the Storage Lens group configuration details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getStorageLensGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getStorageLensGroupCmd).Standalone()

		s3control_getStorageLensGroupCmd.Flags().String("account-id", "", "The Amazon Web Services account ID associated with the Storage Lens group that you're trying to retrieve the details for.")
		s3control_getStorageLensGroupCmd.Flags().String("name", "", "The name of the Storage Lens group that you're trying to retrieve the configuration details for.")
		s3control_getStorageLensGroupCmd.MarkFlagRequired("account-id")
		s3control_getStorageLensGroupCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_getStorageLensGroupCmd)
}
