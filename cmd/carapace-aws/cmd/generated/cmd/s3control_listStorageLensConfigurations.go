package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listStorageLensConfigurationsCmd = &cobra.Command{
	Use:   "list-storage-lens-configurations",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listStorageLensConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_listStorageLensConfigurationsCmd).Standalone()

		s3control_listStorageLensConfigurationsCmd.Flags().String("account-id", "", "The account ID of the requester.")
		s3control_listStorageLensConfigurationsCmd.Flags().String("next-token", "", "A pagination token to request the next page of results.")
		s3control_listStorageLensConfigurationsCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_listStorageLensConfigurationsCmd)
}
