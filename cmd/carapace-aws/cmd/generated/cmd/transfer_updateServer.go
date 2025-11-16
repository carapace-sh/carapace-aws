package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateServerCmd = &cobra.Command{
	Use:   "update-server",
	Short: "Updates the file transfer protocol-enabled server's properties after that server has been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateServerCmd).Standalone()

	transfer_updateServerCmd.Flags().String("certificate", "", "The Amazon Resource Name (ARN) of the Amazon Web ServicesCertificate Manager (ACM) certificate.")
	transfer_updateServerCmd.Flags().String("endpoint-details", "", "The virtual private cloud (VPC) endpoint settings that are configured for your server.")
	transfer_updateServerCmd.Flags().String("endpoint-type", "", "The type of endpoint that you want your server to use.")
	transfer_updateServerCmd.Flags().String("host-key", "", "The RSA, ECDSA, or ED25519 private key to use for your SFTP-enabled server.")
	transfer_updateServerCmd.Flags().String("identity-provider-details", "", "An array containing all of the information required to call a customer's authentication API method.")
	transfer_updateServerCmd.Flags().String("identity-provider-type", "", "The mode of authentication for a server.")
	transfer_updateServerCmd.Flags().String("ip-address-type", "", "Specifies whether to use IPv4 only, or to use dual-stack (IPv4 and IPv6) for your Transfer Family endpoint.")
	transfer_updateServerCmd.Flags().String("logging-role", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFS events.")
	transfer_updateServerCmd.Flags().String("post-authentication-login-banner", "", "Specifies a string to display when users connect to a server.")
	transfer_updateServerCmd.Flags().String("pre-authentication-login-banner", "", "Specifies a string to display when users connect to a server.")
	transfer_updateServerCmd.Flags().String("protocol-details", "", "The protocol settings that are configured for your server.")
	transfer_updateServerCmd.Flags().String("protocols", "", "Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint.")
	transfer_updateServerCmd.Flags().String("s3-storage-options", "", "Specifies whether or not performance for your Amazon S3 directories is optimized.")
	transfer_updateServerCmd.Flags().String("security-policy-name", "", "Specifies the name of the security policy for the server.")
	transfer_updateServerCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server instance that the Transfer Family user is assigned to.")
	transfer_updateServerCmd.Flags().String("structured-log-destinations", "", "Specifies the log groups to which your server logs are sent.")
	transfer_updateServerCmd.Flags().String("workflow-details", "", "Specifies the workflow ID for the workflow to assign and the execution role that's used for executing the workflow.")
	transfer_updateServerCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_updateServerCmd)
}
