package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_deleteDiscovererCmd = &cobra.Command{
	Use:   "delete-discoverer",
	Short: "Deletes a discoverer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_deleteDiscovererCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_deleteDiscovererCmd).Standalone()

		schemas_deleteDiscovererCmd.Flags().String("discoverer-id", "", "The ID of the discoverer.")
		schemas_deleteDiscovererCmd.MarkFlagRequired("discoverer-id")
	})
	schemasCmd.AddCommand(schemas_deleteDiscovererCmd)
}
