package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_unregisterConnectorCmd = &cobra.Command{
	Use:   "unregister-connector",
	Short: "Unregisters the custom connector registered in your account that matches the connector label provided in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_unregisterConnectorCmd).Standalone()

	appflow_unregisterConnectorCmd.Flags().String("connector-label", "", "The label of the connector.")
	appflow_unregisterConnectorCmd.Flags().Bool("force-delete", false, "Indicates whether Amazon AppFlow should unregister the connector, even if it is currently in use in one or more connector profiles.")
	appflow_unregisterConnectorCmd.Flags().Bool("no-force-delete", false, "Indicates whether Amazon AppFlow should unregister the connector, even if it is currently in use in one or more connector profiles.")
	appflow_unregisterConnectorCmd.MarkFlagRequired("connector-label")
	appflow_unregisterConnectorCmd.Flag("no-force-delete").Hidden = true
	appflowCmd.AddCommand(appflow_unregisterConnectorCmd)
}
