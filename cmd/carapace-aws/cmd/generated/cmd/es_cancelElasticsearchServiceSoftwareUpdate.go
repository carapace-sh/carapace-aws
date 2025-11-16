package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_cancelElasticsearchServiceSoftwareUpdateCmd = &cobra.Command{
	Use:   "cancel-elasticsearch-service-software-update",
	Short: "Cancels a scheduled service software update for an Amazon ES domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_cancelElasticsearchServiceSoftwareUpdateCmd).Standalone()

	es_cancelElasticsearchServiceSoftwareUpdateCmd.Flags().String("domain-name", "", "The name of the domain that you want to stop the latest service software update on.")
	es_cancelElasticsearchServiceSoftwareUpdateCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_cancelElasticsearchServiceSoftwareUpdateCmd)
}
