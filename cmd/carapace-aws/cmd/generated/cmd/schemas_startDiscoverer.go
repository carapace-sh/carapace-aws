package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_startDiscovererCmd = &cobra.Command{
	Use:   "start-discoverer",
	Short: "Starts the discoverer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_startDiscovererCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_startDiscovererCmd).Standalone()

		schemas_startDiscovererCmd.Flags().String("discoverer-id", "", "The ID of the discoverer.")
		schemas_startDiscovererCmd.MarkFlagRequired("discoverer-id")
	})
	schemasCmd.AddCommand(schemas_startDiscovererCmd)
}
