package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_updateDiscovererCmd = &cobra.Command{
	Use:   "update-discoverer",
	Short: "Updates the discoverer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_updateDiscovererCmd).Standalone()

	schemas_updateDiscovererCmd.Flags().String("cross-account", "", "Support discovery of schemas in events sent to the bus from another account.")
	schemas_updateDiscovererCmd.Flags().String("description", "", "The description of the discoverer to update.")
	schemas_updateDiscovererCmd.Flags().String("discoverer-id", "", "The ID of the discoverer.")
	schemas_updateDiscovererCmd.MarkFlagRequired("discoverer-id")
	schemasCmd.AddCommand(schemas_updateDiscovererCmd)
}
