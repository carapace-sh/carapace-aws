package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_createCustomPluginCmd = &cobra.Command{
	Use:   "create-custom-plugin",
	Short: "Creates a custom plugin using the specified properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_createCustomPluginCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaconnect_createCustomPluginCmd).Standalone()

		kafkaconnect_createCustomPluginCmd.Flags().String("content-type", "", "The type of the plugin file.")
		kafkaconnect_createCustomPluginCmd.Flags().String("description", "", "A summary description of the custom plugin.")
		kafkaconnect_createCustomPluginCmd.Flags().String("location", "", "Information about the location of a custom plugin.")
		kafkaconnect_createCustomPluginCmd.Flags().String("name", "", "The name of the custom plugin.")
		kafkaconnect_createCustomPluginCmd.Flags().String("tags", "", "The tags you want to attach to the custom plugin.")
		kafkaconnect_createCustomPluginCmd.MarkFlagRequired("content-type")
		kafkaconnect_createCustomPluginCmd.MarkFlagRequired("location")
		kafkaconnect_createCustomPluginCmd.MarkFlagRequired("name")
	})
	kafkaconnectCmd.AddCommand(kafkaconnect_createCustomPluginCmd)
}
