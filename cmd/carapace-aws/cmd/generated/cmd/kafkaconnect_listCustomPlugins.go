package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_listCustomPluginsCmd = &cobra.Command{
	Use:   "list-custom-plugins",
	Short: "Returns a list of all of the custom plugins in this account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_listCustomPluginsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaconnect_listCustomPluginsCmd).Standalone()

		kafkaconnect_listCustomPluginsCmd.Flags().String("max-results", "", "The maximum number of custom plugins to list in one response.")
		kafkaconnect_listCustomPluginsCmd.Flags().String("name-prefix", "", "Lists custom plugin names that start with the specified text string.")
		kafkaconnect_listCustomPluginsCmd.Flags().String("next-token", "", "If the response of a ListCustomPlugins operation is truncated, it will include a NextToken.")
	})
	kafkaconnectCmd.AddCommand(kafkaconnect_listCustomPluginsCmd)
}
