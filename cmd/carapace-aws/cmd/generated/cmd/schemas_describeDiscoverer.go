package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_describeDiscovererCmd = &cobra.Command{
	Use:   "describe-discoverer",
	Short: "Describes the discoverer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_describeDiscovererCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_describeDiscovererCmd).Standalone()

		schemas_describeDiscovererCmd.Flags().String("discoverer-id", "", "The ID of the discoverer.")
		schemas_describeDiscovererCmd.MarkFlagRequired("discoverer-id")
	})
	schemasCmd.AddCommand(schemas_describeDiscovererCmd)
}
