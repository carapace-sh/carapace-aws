package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateFileSystemAssociationCmd = &cobra.Command{
	Use:   "update-file-system-association",
	Short: "Updates a file system association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateFileSystemAssociationCmd).Standalone()

	storagegateway_updateFileSystemAssociationCmd.Flags().String("audit-destination-arn", "", "The Amazon Resource Name (ARN) of the storage used for the audit logs.")
	storagegateway_updateFileSystemAssociationCmd.Flags().String("cache-attributes", "", "")
	storagegateway_updateFileSystemAssociationCmd.Flags().String("file-system-association-arn", "", "The Amazon Resource Name (ARN) of the file system association that you want to update.")
	storagegateway_updateFileSystemAssociationCmd.Flags().String("password", "", "The password of the user credential.")
	storagegateway_updateFileSystemAssociationCmd.Flags().String("user-name", "", "The user name of the user credential that has permission to access the root share D$ of the Amazon FSx file system.")
	storagegateway_updateFileSystemAssociationCmd.MarkFlagRequired("file-system-association-arn")
	storagegatewayCmd.AddCommand(storagegateway_updateFileSystemAssociationCmd)
}
