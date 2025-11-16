package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_createServerCmd = &cobra.Command{
	Use:   "create-server",
	Short: "Instantiates an auto-scaling virtual server based on the selected file transfer protocol in Amazon Web Services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_createServerCmd).Standalone()

	transfer_createServerCmd.Flags().String("certificate", "", "The Amazon Resource Name (ARN) of the Certificate Manager (ACM) certificate.")
	transfer_createServerCmd.Flags().String("domain", "", "The domain of the storage system that is used for file transfers.")
	transfer_createServerCmd.Flags().String("endpoint-details", "", "The virtual private cloud (VPC) endpoint settings that are configured for your server.")
	transfer_createServerCmd.Flags().String("endpoint-type", "", "The type of endpoint that you want your server to use.")
	transfer_createServerCmd.Flags().String("host-key", "", "The RSA, ECDSA, or ED25519 private key to use for your SFTP-enabled server.")
	transfer_createServerCmd.Flags().String("identity-provider-details", "", "Required when `IdentityProviderType` is set to `AWS_DIRECTORY_SERVICE`, `Amazon Web Services_LAMBDA` or `API_GATEWAY`.")
	transfer_createServerCmd.Flags().String("identity-provider-type", "", "The mode of authentication for a server.")
	transfer_createServerCmd.Flags().String("ip-address-type", "", "Specifies whether to use IPv4 only, or to use dual-stack (IPv4 and IPv6) for your Transfer Family endpoint.")
	transfer_createServerCmd.Flags().String("logging-role", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFS events.")
	transfer_createServerCmd.Flags().String("post-authentication-login-banner", "", "Specifies a string to display when users connect to a server.")
	transfer_createServerCmd.Flags().String("pre-authentication-login-banner", "", "Specifies a string to display when users connect to a server.")
	transfer_createServerCmd.Flags().String("protocol-details", "", "The protocol settings that are configured for your server.")
	transfer_createServerCmd.Flags().String("protocols", "", "Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint.")
	transfer_createServerCmd.Flags().String("s3-storage-options", "", "Specifies whether or not performance for your Amazon S3 directories is optimized.")
	transfer_createServerCmd.Flags().String("security-policy-name", "", "Specifies the name of the security policy for the server.")
	transfer_createServerCmd.Flags().String("structured-log-destinations", "", "Specifies the log groups to which your server logs are sent.")
	transfer_createServerCmd.Flags().String("tags", "", "Key-value pairs that can be used to group and search for servers.")
	transfer_createServerCmd.Flags().String("workflow-details", "", "Specifies the workflow ID for the workflow to assign and the execution role that's used for executing the workflow.")
	transferCmd.AddCommand(transfer_createServerCmd)
}
