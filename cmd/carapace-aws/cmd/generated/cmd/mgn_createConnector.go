package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_createConnectorCmd = &cobra.Command{
	Use:   "create-connector",
	Short: "Create Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_createConnectorCmd).Standalone()

	mgn_createConnectorCmd.Flags().String("name", "", "Create Connector request name.")
	mgn_createConnectorCmd.Flags().String("ssm-command-config", "", "Create Connector request SSM command config.")
	mgn_createConnectorCmd.Flags().String("ssm-instance-id", "", "Create Connector request SSM instance ID.")
	mgn_createConnectorCmd.Flags().String("tags", "", "Create Connector request tags.")
	mgn_createConnectorCmd.MarkFlagRequired("name")
	mgn_createConnectorCmd.MarkFlagRequired("ssm-instance-id")
	mgnCmd.AddCommand(mgn_createConnectorCmd)
}
