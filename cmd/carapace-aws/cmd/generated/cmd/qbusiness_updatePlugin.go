package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updatePluginCmd = &cobra.Command{
	Use:   "update-plugin",
	Short: "Updates an Amazon Q Business plugin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updatePluginCmd).Standalone()

	qbusiness_updatePluginCmd.Flags().String("application-id", "", "The identifier of the application the plugin is attached to.")
	qbusiness_updatePluginCmd.Flags().String("auth-configuration", "", "The authentication configuration the plugin is using.")
	qbusiness_updatePluginCmd.Flags().String("custom-plugin-configuration", "", "The configuration for a custom plugin.")
	qbusiness_updatePluginCmd.Flags().String("display-name", "", "The name of the plugin.")
	qbusiness_updatePluginCmd.Flags().String("plugin-id", "", "The identifier of the plugin.")
	qbusiness_updatePluginCmd.Flags().String("server-url", "", "The source URL used for plugin configuration.")
	qbusiness_updatePluginCmd.Flags().String("state", "", "The status of the plugin.")
	qbusiness_updatePluginCmd.MarkFlagRequired("application-id")
	qbusiness_updatePluginCmd.MarkFlagRequired("plugin-id")
	qbusinessCmd.AddCommand(qbusiness_updatePluginCmd)
}
