package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getPluginCmd = &cobra.Command{
	Use:   "get-plugin",
	Short: "Gets information about an existing Amazon Q Business plugin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getPluginCmd).Standalone()

	qbusiness_getPluginCmd.Flags().String("application-id", "", "The identifier of the application which contains the plugin.")
	qbusiness_getPluginCmd.Flags().String("plugin-id", "", "The identifier of the plugin.")
	qbusiness_getPluginCmd.MarkFlagRequired("application-id")
	qbusiness_getPluginCmd.MarkFlagRequired("plugin-id")
	qbusinessCmd.AddCommand(qbusiness_getPluginCmd)
}
