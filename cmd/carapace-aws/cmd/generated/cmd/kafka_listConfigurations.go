package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listConfigurationsCmd = &cobra.Command{
	Use:   "list-configurations",
	Short: "Returns a list of all the MSK configurations in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_listConfigurationsCmd).Standalone()

		kafka_listConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		kafka_listConfigurationsCmd.Flags().String("next-token", "", "The paginated results marker.")
	})
	kafkaCmd.AddCommand(kafka_listConfigurationsCmd)
}
