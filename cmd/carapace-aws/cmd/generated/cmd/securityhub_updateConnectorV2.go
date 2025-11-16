package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateConnectorV2Cmd = &cobra.Command{
	Use:   "update-connector-v2",
	Short: "Grants permission to update a connectorV2 based on its id and input parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateConnectorV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_updateConnectorV2Cmd).Standalone()

		securityhub_updateConnectorV2Cmd.Flags().String("client-secret", "", "The clientSecret of ServiceNow.")
		securityhub_updateConnectorV2Cmd.Flags().String("connector-id", "", "The UUID of the connectorV2 to identify connectorV2 resource.")
		securityhub_updateConnectorV2Cmd.Flags().String("description", "", "The description of the connectorV2.")
		securityhub_updateConnectorV2Cmd.Flags().String("provider", "", "The third-party provider’s service configuration.")
		securityhub_updateConnectorV2Cmd.MarkFlagRequired("connector-id")
	})
	securityhubCmd.AddCommand(securityhub_updateConnectorV2Cmd)
}
