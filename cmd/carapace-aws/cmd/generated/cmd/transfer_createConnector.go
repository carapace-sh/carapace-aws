package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_createConnectorCmd = &cobra.Command{
	Use:   "create-connector",
	Short: "Creates the connector, which captures the parameters for a connection for the AS2 or SFTP protocol.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_createConnectorCmd).Standalone()

	transfer_createConnectorCmd.Flags().String("access-role", "", "Connectors are used to send files using either the AS2 or SFTP protocol.")
	transfer_createConnectorCmd.Flags().String("as2-config", "", "A structure that contains the parameters for an AS2 connector object.")
	transfer_createConnectorCmd.Flags().String("egress-config", "", "Specifies the egress configuration for the connector, which determines how traffic is routed from the connector to the SFTP server.")
	transfer_createConnectorCmd.Flags().String("logging-role", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that allows a connector to turn on CloudWatch logging for Amazon S3 events.")
	transfer_createConnectorCmd.Flags().String("security-policy-name", "", "Specifies the name of the security policy for the connector.")
	transfer_createConnectorCmd.Flags().String("sftp-config", "", "A structure that contains the parameters for an SFTP connector object.")
	transfer_createConnectorCmd.Flags().String("tags", "", "Key-value pairs that can be used to group and search for connectors.")
	transfer_createConnectorCmd.Flags().String("url", "", "The URL of the partner's AS2 or SFTP endpoint.")
	transfer_createConnectorCmd.MarkFlagRequired("access-role")
	transferCmd.AddCommand(transfer_createConnectorCmd)
}
