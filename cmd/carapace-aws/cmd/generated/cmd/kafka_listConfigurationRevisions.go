package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listConfigurationRevisionsCmd = &cobra.Command{
	Use:   "list-configuration-revisions",
	Short: "Returns a list of all the MSK configurations in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listConfigurationRevisionsCmd).Standalone()

	kafka_listConfigurationRevisionsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that uniquely identifies an MSK configuration and all of its revisions.")
	kafka_listConfigurationRevisionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	kafka_listConfigurationRevisionsCmd.Flags().String("next-token", "", "The paginated results marker.")
	kafka_listConfigurationRevisionsCmd.MarkFlagRequired("arn")
	kafkaCmd.AddCommand(kafka_listConfigurationRevisionsCmd)
}
