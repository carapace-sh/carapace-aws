package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_deleteConnectorCmd = &cobra.Command{
	Use:   "delete-connector",
	Short: "Delete Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_deleteConnectorCmd).Standalone()

	mgn_deleteConnectorCmd.Flags().String("connector-id", "", "Delete Connector request connector ID.")
	mgn_deleteConnectorCmd.MarkFlagRequired("connector-id")
	mgnCmd.AddCommand(mgn_deleteConnectorCmd)
}
