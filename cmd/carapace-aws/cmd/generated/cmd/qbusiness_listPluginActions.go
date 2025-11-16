package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listPluginActionsCmd = &cobra.Command{
	Use:   "list-plugin-actions",
	Short: "Lists configured Amazon Q Business actions for a specific plugin in an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listPluginActionsCmd).Standalone()

	qbusiness_listPluginActionsCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application the plugin is attached to.")
	qbusiness_listPluginActionsCmd.Flags().String("max-results", "", "The maximum number of plugin actions to return.")
	qbusiness_listPluginActionsCmd.Flags().String("next-token", "", "If the number of plugin actions returned exceeds `maxResults`, Amazon Q Business returns a next token as a pagination token to retrieve the next set of plugin actions.")
	qbusiness_listPluginActionsCmd.Flags().String("plugin-id", "", "The identifier of the Amazon Q Business plugin.")
	qbusiness_listPluginActionsCmd.MarkFlagRequired("application-id")
	qbusiness_listPluginActionsCmd.MarkFlagRequired("plugin-id")
	qbusinessCmd.AddCommand(qbusiness_listPluginActionsCmd)
}
