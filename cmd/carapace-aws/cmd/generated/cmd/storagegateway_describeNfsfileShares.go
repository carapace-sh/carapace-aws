package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeNfsfileSharesCmd = &cobra.Command{
	Use:   "describe-nfsfile-shares",
	Short: "Gets a description for one or more Network File System (NFS) file shares from an S3 File Gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeNfsfileSharesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeNfsfileSharesCmd).Standalone()

		storagegateway_describeNfsfileSharesCmd.Flags().String("file-share-arnlist", "", "An array containing the Amazon Resource Name (ARN) of each file share to be described.")
		storagegateway_describeNfsfileSharesCmd.MarkFlagRequired("file-share-arnlist")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeNfsfileSharesCmd)
}
