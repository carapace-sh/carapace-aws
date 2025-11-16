package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_deleteConnectorV2Cmd = &cobra.Command{
	Use:   "delete-connector-v2",
	Short: "Grants permission to delete a connectorV2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_deleteConnectorV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_deleteConnectorV2Cmd).Standalone()

		securityhub_deleteConnectorV2Cmd.Flags().String("connector-id", "", "The UUID of the connectorV2 to identify connectorV2 resource.")
		securityhub_deleteConnectorV2Cmd.MarkFlagRequired("connector-id")
	})
	securityhubCmd.AddCommand(securityhub_deleteConnectorV2Cmd)
}
