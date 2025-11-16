package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_updateConnectorProfileCmd = &cobra.Command{
	Use:   "update-connector-profile",
	Short: "Updates a given connector profile associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_updateConnectorProfileCmd).Standalone()

	appflow_updateConnectorProfileCmd.Flags().String("client-token", "", "The `clientToken` parameter is an idempotency token.")
	appflow_updateConnectorProfileCmd.Flags().String("connection-mode", "", "Indicates the connection mode and if it is public or private.")
	appflow_updateConnectorProfileCmd.Flags().String("connector-profile-config", "", "Defines the connector-specific profile configuration and credentials.")
	appflow_updateConnectorProfileCmd.Flags().String("connector-profile-name", "", "The name of the connector profile and is unique for each `ConnectorProfile` in the Amazon Web Services account.")
	appflow_updateConnectorProfileCmd.MarkFlagRequired("connection-mode")
	appflow_updateConnectorProfileCmd.MarkFlagRequired("connector-profile-config")
	appflow_updateConnectorProfileCmd.MarkFlagRequired("connector-profile-name")
	appflowCmd.AddCommand(appflow_updateConnectorProfileCmd)
}
