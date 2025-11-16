package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createPluginCmd = &cobra.Command{
	Use:   "create-plugin",
	Short: "Creates an Amazon Q Business plugin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createPluginCmd).Standalone()

	qbusiness_createPluginCmd.Flags().String("application-id", "", "The identifier of the application that will contain the plugin.")
	qbusiness_createPluginCmd.Flags().String("auth-configuration", "", "")
	qbusiness_createPluginCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create your Amazon Q Business plugin.")
	qbusiness_createPluginCmd.Flags().String("custom-plugin-configuration", "", "Contains configuration for a custom plugin.")
	qbusiness_createPluginCmd.Flags().String("display-name", "", "A the name for your plugin.")
	qbusiness_createPluginCmd.Flags().String("server-url", "", "The source URL used for plugin configuration.")
	qbusiness_createPluginCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize the data source connector.")
	qbusiness_createPluginCmd.Flags().String("type", "", "The type of plugin you want to create.")
	qbusiness_createPluginCmd.MarkFlagRequired("application-id")
	qbusiness_createPluginCmd.MarkFlagRequired("auth-configuration")
	qbusiness_createPluginCmd.MarkFlagRequired("display-name")
	qbusiness_createPluginCmd.MarkFlagRequired("type")
	qbusinessCmd.AddCommand(qbusiness_createPluginCmd)
}
