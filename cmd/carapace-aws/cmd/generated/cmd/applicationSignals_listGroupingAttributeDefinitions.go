package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listGroupingAttributeDefinitionsCmd = &cobra.Command{
	Use:   "list-grouping-attribute-definitions",
	Short: "Retrieves the available grouping attribute definitions that can be used to create grouping configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listGroupingAttributeDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_listGroupingAttributeDefinitionsCmd).Standalone()

		applicationSignals_listGroupingAttributeDefinitionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_listGroupingAttributeDefinitionsCmd)
}
