package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listConnectorsCmd = &cobra.Command{
	Use:   "list-connectors",
	Short: "List Connectors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listConnectorsCmd).Standalone()

	mgn_listConnectorsCmd.Flags().String("filters", "", "List Connectors Request filters.")
	mgn_listConnectorsCmd.Flags().String("max-results", "", "List Connectors Request max results.")
	mgn_listConnectorsCmd.Flags().String("next-token", "", "List Connectors Request next token.")
	mgnCmd.AddCommand(mgn_listConnectorsCmd)
}
