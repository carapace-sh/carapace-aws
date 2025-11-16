package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeFileSystemAssociationsCmd = &cobra.Command{
	Use:   "describe-file-system-associations",
	Short: "Gets the file system association information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeFileSystemAssociationsCmd).Standalone()

	storagegateway_describeFileSystemAssociationsCmd.Flags().String("file-system-association-arnlist", "", "An array containing the Amazon Resource Name (ARN) of each file system association to be described.")
	storagegateway_describeFileSystemAssociationsCmd.MarkFlagRequired("file-system-association-arnlist")
	storagegatewayCmd.AddCommand(storagegateway_describeFileSystemAssociationsCmd)
}
