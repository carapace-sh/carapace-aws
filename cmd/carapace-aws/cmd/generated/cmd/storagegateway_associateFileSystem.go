package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_associateFileSystemCmd = &cobra.Command{
	Use:   "associate-file-system",
	Short: "Associate an Amazon FSx file system with the FSx File Gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_associateFileSystemCmd).Standalone()

	storagegateway_associateFileSystemCmd.Flags().String("audit-destination-arn", "", "The Amazon Resource Name (ARN) of the storage used for the audit logs.")
	storagegateway_associateFileSystemCmd.Flags().String("cache-attributes", "", "")
	storagegateway_associateFileSystemCmd.Flags().String("client-token", "", "A unique string value that you supply that is used by the FSx File Gateway to ensure idempotent file system association creation.")
	storagegateway_associateFileSystemCmd.Flags().String("endpoint-network-configuration", "", "Specifies the network configuration information for the gateway associated with the Amazon FSx file system.")
	storagegateway_associateFileSystemCmd.Flags().String("gateway-arn", "", "")
	storagegateway_associateFileSystemCmd.Flags().String("location-arn", "", "The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with the FSx File Gateway.")
	storagegateway_associateFileSystemCmd.Flags().String("password", "", "The password of the user credential.")
	storagegateway_associateFileSystemCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to the file system association.")
	storagegateway_associateFileSystemCmd.Flags().String("user-name", "", "The user name of the user credential that has permission to access the root share D$ of the Amazon FSx file system.")
	storagegateway_associateFileSystemCmd.MarkFlagRequired("client-token")
	storagegateway_associateFileSystemCmd.MarkFlagRequired("gateway-arn")
	storagegateway_associateFileSystemCmd.MarkFlagRequired("location-arn")
	storagegateway_associateFileSystemCmd.MarkFlagRequired("password")
	storagegateway_associateFileSystemCmd.MarkFlagRequired("user-name")
	storagegatewayCmd.AddCommand(storagegateway_associateFileSystemCmd)
}
