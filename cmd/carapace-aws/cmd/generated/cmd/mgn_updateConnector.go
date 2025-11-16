package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_updateConnectorCmd = &cobra.Command{
	Use:   "update-connector",
	Short: "Update Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_updateConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_updateConnectorCmd).Standalone()

		mgn_updateConnectorCmd.Flags().String("connector-id", "", "Update Connector request connector ID.")
		mgn_updateConnectorCmd.Flags().String("name", "", "Update Connector request name.")
		mgn_updateConnectorCmd.Flags().String("ssm-command-config", "", "Update Connector request SSM command config.")
		mgn_updateConnectorCmd.MarkFlagRequired("connector-id")
	})
	mgnCmd.AddCommand(mgn_updateConnectorCmd)
}
