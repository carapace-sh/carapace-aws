package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_createConnectorV2Cmd = &cobra.Command{
	Use:   "create-connector-v2",
	Short: "Grants permission to create a connectorV2 based on input parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_createConnectorV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_createConnectorV2Cmd).Standalone()

		securityhub_createConnectorV2Cmd.Flags().String("client-token", "", "A unique identifier used to ensure idempotency.")
		securityhub_createConnectorV2Cmd.Flags().String("description", "", "The description of the connectorV2.")
		securityhub_createConnectorV2Cmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of KMS key used to encrypt secrets for the connectorV2.")
		securityhub_createConnectorV2Cmd.Flags().String("name", "", "The unique name of the connectorV2.")
		securityhub_createConnectorV2Cmd.Flags().String("provider", "", "The third-party provider’s service configuration.")
		securityhub_createConnectorV2Cmd.Flags().String("tags", "", "The tags to add to the connectorV2 when you create.")
		securityhub_createConnectorV2Cmd.MarkFlagRequired("name")
		securityhub_createConnectorV2Cmd.MarkFlagRequired("provider")
	})
	securityhubCmd.AddCommand(securityhub_createConnectorV2Cmd)
}
