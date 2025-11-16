package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listPluginTypeActionsCmd = &cobra.Command{
	Use:   "list-plugin-type-actions",
	Short: "Lists configured Amazon Q Business actions for any plugin type—both built-in and custom.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listPluginTypeActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listPluginTypeActionsCmd).Standalone()

		qbusiness_listPluginTypeActionsCmd.Flags().String("max-results", "", "The maximum number of plugins to return.")
		qbusiness_listPluginTypeActionsCmd.Flags().String("next-token", "", "If the number of plugins returned exceeds `maxResults`, Amazon Q Business returns a next token as a pagination token to retrieve the next set of plugins.")
		qbusiness_listPluginTypeActionsCmd.Flags().String("plugin-type", "", "The type of the plugin.")
		qbusiness_listPluginTypeActionsCmd.MarkFlagRequired("plugin-type")
	})
	qbusinessCmd.AddCommand(qbusiness_listPluginTypeActionsCmd)
}
