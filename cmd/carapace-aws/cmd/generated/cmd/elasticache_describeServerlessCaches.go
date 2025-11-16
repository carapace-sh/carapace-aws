package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeServerlessCachesCmd = &cobra.Command{
	Use:   "describe-serverless-caches",
	Short: "Returns information about a specific serverless cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeServerlessCachesCmd).Standalone()

	elasticache_describeServerlessCachesCmd.Flags().String("max-results", "", "The maximum number of records in the response.")
	elasticache_describeServerlessCachesCmd.Flags().String("next-token", "", "An optional marker returned from a prior request to support pagination of results from this operation.")
	elasticache_describeServerlessCachesCmd.Flags().String("serverless-cache-name", "", "The identifier for the serverless cache.")
	elasticacheCmd.AddCommand(elasticache_describeServerlessCachesCmd)
}
