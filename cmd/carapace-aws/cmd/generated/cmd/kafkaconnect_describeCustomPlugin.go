package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_describeCustomPluginCmd = &cobra.Command{
	Use:   "describe-custom-plugin",
	Short: "A summary description of the custom plugin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_describeCustomPluginCmd).Standalone()

	kafkaconnect_describeCustomPluginCmd.Flags().String("custom-plugin-arn", "", "Returns information about a custom plugin.")
	kafkaconnect_describeCustomPluginCmd.MarkFlagRequired("custom-plugin-arn")
	kafkaconnectCmd.AddCommand(kafkaconnect_describeCustomPluginCmd)
}
