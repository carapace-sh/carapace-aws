package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_stopDiscovererCmd = &cobra.Command{
	Use:   "stop-discoverer",
	Short: "Stops the discoverer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_stopDiscovererCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_stopDiscovererCmd).Standalone()

		schemas_stopDiscovererCmd.Flags().String("discoverer-id", "", "The ID of the discoverer.")
		schemas_stopDiscovererCmd.MarkFlagRequired("discoverer-id")
	})
	schemasCmd.AddCommand(schemas_stopDiscovererCmd)
}
