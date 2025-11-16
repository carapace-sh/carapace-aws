package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listInstanceTypeDetailsCmd = &cobra.Command{
	Use:   "list-instance-type-details",
	Short: "Lists all instance types and available features for a given OpenSearch or Elasticsearch version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listInstanceTypeDetailsCmd).Standalone()

	opensearch_listInstanceTypeDetailsCmd.Flags().String("domain-name", "", "The name of the domain.")
	opensearch_listInstanceTypeDetailsCmd.Flags().String("engine-version", "", "The version of OpenSearch or Elasticsearch, in the format Elasticsearch\\_X.Y or OpenSearch\\_X.Y. Defaults to the latest version of OpenSearch.")
	opensearch_listInstanceTypeDetailsCmd.Flags().String("instance-type", "", "An optional parameter that lists information for a given instance type.")
	opensearch_listInstanceTypeDetailsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearch_listInstanceTypeDetailsCmd.Flags().String("next-token", "", "If your initial `ListInstanceTypeDetails` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListInstanceTypeDetails` operations, which returns results in the next page.")
	opensearch_listInstanceTypeDetailsCmd.Flags().Bool("no-retrieve-azs", false, "An optional parameter that specifies the Availability Zones for the domain.")
	opensearch_listInstanceTypeDetailsCmd.Flags().Bool("retrieve-azs", false, "An optional parameter that specifies the Availability Zones for the domain.")
	opensearch_listInstanceTypeDetailsCmd.MarkFlagRequired("engine-version")
	opensearch_listInstanceTypeDetailsCmd.Flag("no-retrieve-azs").Hidden = true
	opensearchCmd.AddCommand(opensearch_listInstanceTypeDetailsCmd)
}
