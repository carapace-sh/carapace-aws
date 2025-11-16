package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_updateConnectorRegistrationCmd = &cobra.Command{
	Use:   "update-connector-registration",
	Short: "Updates a custom connector that you've previously registered.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_updateConnectorRegistrationCmd).Standalone()

	appflow_updateConnectorRegistrationCmd.Flags().String("client-token", "", "The `clientToken` parameter is an idempotency token.")
	appflow_updateConnectorRegistrationCmd.Flags().String("connector-label", "", "The name of the connector.")
	appflow_updateConnectorRegistrationCmd.Flags().String("connector-provisioning-config", "", "")
	appflow_updateConnectorRegistrationCmd.Flags().String("description", "", "A description about the update that you're applying to the connector.")
	appflow_updateConnectorRegistrationCmd.MarkFlagRequired("connector-label")
	appflowCmd.AddCommand(appflow_updateConnectorRegistrationCmd)
}
