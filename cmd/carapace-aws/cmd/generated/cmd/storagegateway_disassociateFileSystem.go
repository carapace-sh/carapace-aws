package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_disassociateFileSystemCmd = &cobra.Command{
	Use:   "disassociate-file-system",
	Short: "Disassociates an Amazon FSx file system from the specified gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_disassociateFileSystemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_disassociateFileSystemCmd).Standalone()

		storagegateway_disassociateFileSystemCmd.Flags().String("file-system-association-arn", "", "The Amazon Resource Name (ARN) of the file system association to be deleted.")
		storagegateway_disassociateFileSystemCmd.Flags().String("force-delete", "", "If this value is set to true, the operation disassociates an Amazon FSx file system immediately.")
		storagegateway_disassociateFileSystemCmd.MarkFlagRequired("file-system-association-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_disassociateFileSystemCmd)
}
