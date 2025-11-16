package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_deleteCustomPluginCmd = &cobra.Command{
	Use:   "delete-custom-plugin",
	Short: "Deletes a custom plugin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_deleteCustomPluginCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaconnect_deleteCustomPluginCmd).Standalone()

		kafkaconnect_deleteCustomPluginCmd.Flags().String("custom-plugin-arn", "", "The Amazon Resource Name (ARN) of the custom plugin that you want to delete.")
		kafkaconnect_deleteCustomPluginCmd.MarkFlagRequired("custom-plugin-arn")
	})
	kafkaconnectCmd.AddCommand(kafkaconnect_deleteCustomPluginCmd)
}
