package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_listWorkerConfigurationsCmd = &cobra.Command{
	Use:   "list-worker-configurations",
	Short: "Returns a list of all of the worker configurations in this account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_listWorkerConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaconnect_listWorkerConfigurationsCmd).Standalone()

		kafkaconnect_listWorkerConfigurationsCmd.Flags().String("max-results", "", "The maximum number of worker configurations to list in one response.")
		kafkaconnect_listWorkerConfigurationsCmd.Flags().String("name-prefix", "", "Lists worker configuration names that start with the specified text string.")
		kafkaconnect_listWorkerConfigurationsCmd.Flags().String("next-token", "", "If the response of a ListWorkerConfigurations operation is truncated, it will include a NextToken.")
	})
	kafkaconnectCmd.AddCommand(kafkaconnect_listWorkerConfigurationsCmd)
}
