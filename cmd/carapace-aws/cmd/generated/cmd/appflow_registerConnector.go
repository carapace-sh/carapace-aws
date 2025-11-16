package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_registerConnectorCmd = &cobra.Command{
	Use:   "register-connector",
	Short: "Registers a new custom connector with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_registerConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_registerConnectorCmd).Standalone()

		appflow_registerConnectorCmd.Flags().String("client-token", "", "The `clientToken` parameter is an idempotency token.")
		appflow_registerConnectorCmd.Flags().String("connector-label", "", "The name of the connector.")
		appflow_registerConnectorCmd.Flags().String("connector-provisioning-config", "", "The provisioning type of the connector.")
		appflow_registerConnectorCmd.Flags().String("connector-provisioning-type", "", "The provisioning type of the connector.")
		appflow_registerConnectorCmd.Flags().String("description", "", "A description about the connector that's being registered.")
	})
	appflowCmd.AddCommand(appflow_registerConnectorCmd)
}
