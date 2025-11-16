package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateConnectorCmd = &cobra.Command{
	Use:   "update-connector",
	Short: "Updates some of the parameters for an existing connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateConnectorCmd).Standalone()

	transfer_updateConnectorCmd.Flags().String("access-role", "", "Connectors are used to send files using either the AS2 or SFTP protocol.")
	transfer_updateConnectorCmd.Flags().String("as2-config", "", "A structure that contains the parameters for an AS2 connector object.")
	transfer_updateConnectorCmd.Flags().String("connector-id", "", "The unique identifier for the connector.")
	transfer_updateConnectorCmd.Flags().String("egress-config", "", "Updates the egress configuration for the connector, allowing you to modify how traffic is routed from the connector to the SFTP server.")
	transfer_updateConnectorCmd.Flags().String("logging-role", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that allows a connector to turn on CloudWatch logging for Amazon S3 events.")
	transfer_updateConnectorCmd.Flags().String("security-policy-name", "", "Specifies the name of the security policy for the connector.")
	transfer_updateConnectorCmd.Flags().String("sftp-config", "", "A structure that contains the parameters for an SFTP connector object.")
	transfer_updateConnectorCmd.Flags().String("url", "", "The URL of the partner's AS2 or SFTP endpoint.")
	transfer_updateConnectorCmd.MarkFlagRequired("connector-id")
	transferCmd.AddCommand(transfer_updateConnectorCmd)
}
