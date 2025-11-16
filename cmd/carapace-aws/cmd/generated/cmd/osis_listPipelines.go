package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_listPipelinesCmd = &cobra.Command{
	Use:   "list-pipelines",
	Short: "Lists all OpenSearch Ingestion pipelines in the current Amazon Web Services account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_listPipelinesCmd).Standalone()

	osis_listPipelinesCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	osis_listPipelinesCmd.Flags().String("next-token", "", "If your initial `ListPipelines` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListPipelines` operations, which returns results in the next page.")
	osisCmd.AddCommand(osis_listPipelinesCmd)
}
