package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Returns information about provisioned Amazon DocumentDB elastic clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_listClustersCmd).Standalone()

	docdbElastic_listClustersCmd.Flags().String("max-results", "", "The maximum number of elastic cluster snapshot results to receive in the response.")
	docdbElastic_listClustersCmd.Flags().String("next-token", "", "A pagination token provided by a previous request.")
	docdbElasticCmd.AddCommand(docdbElastic_listClustersCmd)
}
