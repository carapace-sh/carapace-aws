package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_startElasticsearchServiceSoftwareUpdateCmd = &cobra.Command{
	Use:   "start-elasticsearch-service-software-update",
	Short: "Schedules a service software update for an Amazon ES domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_startElasticsearchServiceSoftwareUpdateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_startElasticsearchServiceSoftwareUpdateCmd).Standalone()

		es_startElasticsearchServiceSoftwareUpdateCmd.Flags().String("domain-name", "", "The name of the domain that you want to update to the latest service software.")
		es_startElasticsearchServiceSoftwareUpdateCmd.MarkFlagRequired("domain-name")
	})
	esCmd.AddCommand(es_startElasticsearchServiceSoftwareUpdateCmd)
}
