package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_createConnectorProfileCmd = &cobra.Command{
	Use:   "create-connector-profile",
	Short: "Creates a new connector profile associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_createConnectorProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_createConnectorProfileCmd).Standalone()

		appflow_createConnectorProfileCmd.Flags().String("client-token", "", "The `clientToken` parameter is an idempotency token.")
		appflow_createConnectorProfileCmd.Flags().String("connection-mode", "", "Indicates the connection mode and specifies whether it is public or private.")
		appflow_createConnectorProfileCmd.Flags().String("connector-label", "", "The label of the connector.")
		appflow_createConnectorProfileCmd.Flags().String("connector-profile-config", "", "Defines the connector-specific configuration and credentials.")
		appflow_createConnectorProfileCmd.Flags().String("connector-profile-name", "", "The name of the connector profile.")
		appflow_createConnectorProfileCmd.Flags().String("connector-type", "", "The type of connector, such as Salesforce, Amplitude, and so on.")
		appflow_createConnectorProfileCmd.Flags().String("kms-arn", "", "The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.")
		appflow_createConnectorProfileCmd.MarkFlagRequired("connection-mode")
		appflow_createConnectorProfileCmd.MarkFlagRequired("connector-profile-config")
		appflow_createConnectorProfileCmd.MarkFlagRequired("connector-profile-name")
		appflow_createConnectorProfileCmd.MarkFlagRequired("connector-type")
	})
	appflowCmd.AddCommand(appflow_createConnectorProfileCmd)
}
