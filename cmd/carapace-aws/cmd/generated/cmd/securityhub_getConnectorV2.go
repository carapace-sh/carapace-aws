package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getConnectorV2Cmd = &cobra.Command{
	Use:   "get-connector-v2",
	Short: "Grants permission to retrieve details for a connectorV2 based on connector id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getConnectorV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_getConnectorV2Cmd).Standalone()

		securityhub_getConnectorV2Cmd.Flags().String("connector-id", "", "The UUID of the connectorV2 to identify connectorV2 resource.")
		securityhub_getConnectorV2Cmd.MarkFlagRequired("connector-id")
	})
	securityhubCmd.AddCommand(securityhub_getConnectorV2Cmd)
}
