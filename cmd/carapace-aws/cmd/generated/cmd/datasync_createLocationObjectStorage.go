package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationObjectStorageCmd = &cobra.Command{
	Use:   "create-location-object-storage",
	Short: "Creates a transfer *location* for an object storage system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationObjectStorageCmd).Standalone()

	datasync_createLocationObjectStorageCmd.Flags().String("access-key", "", "Specifies the access key (for example, a user name) if credentials are required to authenticate with the object storage server.")
	datasync_createLocationObjectStorageCmd.Flags().String("agent-arns", "", "(Optional) Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can connect with your object storage system.")
	datasync_createLocationObjectStorageCmd.Flags().String("bucket-name", "", "Specifies the name of the object storage bucket involved in the transfer.")
	datasync_createLocationObjectStorageCmd.Flags().String("cmk-secret-config", "", "Specifies configuration information for a DataSync-managed secret, which includes the `SecretKey` that DataSync uses to access a specific object storage location, with a customer-managed KMS key.")
	datasync_createLocationObjectStorageCmd.Flags().String("custom-secret-config", "", "Specifies configuration information for a customer-managed Secrets Manager secret where the secret key for a specific object storage location is stored in plain text.")
	datasync_createLocationObjectStorageCmd.Flags().String("secret-key", "", "Specifies the secret key (for example, a password) if credentials are required to authenticate with the object storage server.")
	datasync_createLocationObjectStorageCmd.Flags().String("server-certificate", "", "Specifies a certificate chain for DataSync to authenticate with your object storage system if the system uses a private or self-signed certificate authority (CA).")
	datasync_createLocationObjectStorageCmd.Flags().String("server-hostname", "", "Specifies the domain name or IP address (IPv4 or IPv6) of the object storage server that your DataSync agent connects to.")
	datasync_createLocationObjectStorageCmd.Flags().String("server-port", "", "Specifies the port that your object storage server accepts inbound network traffic on (for example, port 443).")
	datasync_createLocationObjectStorageCmd.Flags().String("server-protocol", "", "Specifies the protocol that your object storage server uses to communicate.")
	datasync_createLocationObjectStorageCmd.Flags().String("subdirectory", "", "Specifies the object prefix for your object storage server.")
	datasync_createLocationObjectStorageCmd.Flags().String("tags", "", "Specifies the key-value pair that represents a tag that you want to add to the resource.")
	datasync_createLocationObjectStorageCmd.MarkFlagRequired("bucket-name")
	datasync_createLocationObjectStorageCmd.MarkFlagRequired("server-hostname")
	datasyncCmd.AddCommand(datasync_createLocationObjectStorageCmd)
}
