package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationObjectStorageCmd = &cobra.Command{
	Use:   "update-location-object-storage",
	Short: "Modifies the following configuration parameters of the object storage transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationObjectStorageCmd).Standalone()

	datasync_updateLocationObjectStorageCmd.Flags().String("access-key", "", "Specifies the access key (for example, a user name) if credentials are required to authenticate with the object storage server.")
	datasync_updateLocationObjectStorageCmd.Flags().String("agent-arns", "", "(Optional) Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can connect with your object storage system.")
	datasync_updateLocationObjectStorageCmd.Flags().String("cmk-secret-config", "", "Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed KMS key.")
	datasync_updateLocationObjectStorageCmd.Flags().String("custom-secret-config", "", "Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed KMS key.")
	datasync_updateLocationObjectStorageCmd.Flags().String("location-arn", "", "Specifies the ARN of the object storage system location that you're updating.")
	datasync_updateLocationObjectStorageCmd.Flags().String("secret-key", "", "Specifies the secret key (for example, a password) if credentials are required to authenticate with the object storage server.")
	datasync_updateLocationObjectStorageCmd.Flags().String("server-certificate", "", "Specifies a certificate chain for DataSync to authenticate with your object storage system if the system uses a private or self-signed certificate authority (CA).")
	datasync_updateLocationObjectStorageCmd.Flags().String("server-hostname", "", "Specifies the domain name or IP address (IPv4 or IPv6) of the object storage server that your DataSync agent connects to.")
	datasync_updateLocationObjectStorageCmd.Flags().String("server-port", "", "Specifies the port that your object storage server accepts inbound network traffic on (for example, port 443).")
	datasync_updateLocationObjectStorageCmd.Flags().String("server-protocol", "", "Specifies the protocol that your object storage server uses to communicate.")
	datasync_updateLocationObjectStorageCmd.Flags().String("subdirectory", "", "Specifies the object prefix for your object storage server.")
	datasync_updateLocationObjectStorageCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_updateLocationObjectStorageCmd)
}
