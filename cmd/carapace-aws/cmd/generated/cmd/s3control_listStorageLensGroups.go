package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listStorageLensGroupsCmd = &cobra.Command{
	Use:   "list-storage-lens-groups",
	Short: "Lists all the Storage Lens groups in the specified home Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listStorageLensGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_listStorageLensGroupsCmd).Standalone()

		s3control_listStorageLensGroupsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the Storage Lens groups.")
		s3control_listStorageLensGroupsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` if there are no more results.")
		s3control_listStorageLensGroupsCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_listStorageLensGroupsCmd)
}
