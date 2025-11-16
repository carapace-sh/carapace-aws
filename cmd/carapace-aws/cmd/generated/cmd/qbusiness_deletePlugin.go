package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deletePluginCmd = &cobra.Command{
	Use:   "delete-plugin",
	Short: "Deletes an Amazon Q Business plugin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deletePluginCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_deletePluginCmd).Standalone()

		qbusiness_deletePluginCmd.Flags().String("application-id", "", "The identifier the application attached to the Amazon Q Business plugin.")
		qbusiness_deletePluginCmd.Flags().String("plugin-id", "", "The identifier of the plugin being deleted.")
		qbusiness_deletePluginCmd.MarkFlagRequired("application-id")
		qbusiness_deletePluginCmd.MarkFlagRequired("plugin-id")
	})
	qbusinessCmd.AddCommand(qbusiness_deletePluginCmd)
}
