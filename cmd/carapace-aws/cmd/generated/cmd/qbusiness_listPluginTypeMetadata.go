package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listPluginTypeMetadataCmd = &cobra.Command{
	Use:   "list-plugin-type-metadata",
	Short: "Lists metadata for all Amazon Q Business plugin types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listPluginTypeMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listPluginTypeMetadataCmd).Standalone()

		qbusiness_listPluginTypeMetadataCmd.Flags().String("max-results", "", "The maximum number of plugin metadata items to return.")
		qbusiness_listPluginTypeMetadataCmd.Flags().String("next-token", "", "If the metadata returned exceeds `maxResults`, Amazon Q Business returns a next token as a pagination token to retrieve the next set of metadata.")
	})
	qbusinessCmd.AddCommand(qbusiness_listPluginTypeMetadataCmd)
}
