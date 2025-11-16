package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeInstanceTypeLimitsCmd = &cobra.Command{
	Use:   "describe-instance-type-limits",
	Short: "Describes the instance count, storage, and master node limits for a given OpenSearch or Elasticsearch version and instance type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeInstanceTypeLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_describeInstanceTypeLimitsCmd).Standalone()

		opensearch_describeInstanceTypeLimitsCmd.Flags().String("domain-name", "", "The name of the domain.")
		opensearch_describeInstanceTypeLimitsCmd.Flags().String("engine-version", "", "Version of OpenSearch or Elasticsearch, in the format Elasticsearch\\_X.Y or OpenSearch\\_X.Y. Defaults to the latest version of OpenSearch.")
		opensearch_describeInstanceTypeLimitsCmd.Flags().String("instance-type", "", "The OpenSearch Service instance type for which you need limit information.")
		opensearch_describeInstanceTypeLimitsCmd.MarkFlagRequired("engine-version")
		opensearch_describeInstanceTypeLimitsCmd.MarkFlagRequired("instance-type")
	})
	opensearchCmd.AddCommand(opensearch_describeInstanceTypeLimitsCmd)
}
