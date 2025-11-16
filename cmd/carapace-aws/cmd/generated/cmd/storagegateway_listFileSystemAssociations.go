package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listFileSystemAssociationsCmd = &cobra.Command{
	Use:   "list-file-system-associations",
	Short: "Gets a list of `FileSystemAssociationSummary` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listFileSystemAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_listFileSystemAssociationsCmd).Standalone()

		storagegateway_listFileSystemAssociationsCmd.Flags().String("gateway-arn", "", "")
		storagegateway_listFileSystemAssociationsCmd.Flags().String("limit", "", "The maximum number of file system associations to return in the response.")
		storagegateway_listFileSystemAssociationsCmd.Flags().String("marker", "", "Opaque pagination token returned from a previous `ListFileSystemAssociations` operation.")
	})
	storagegatewayCmd.AddCommand(storagegateway_listFileSystemAssociationsCmd)
}
