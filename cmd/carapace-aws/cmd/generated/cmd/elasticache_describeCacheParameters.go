package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeCacheParametersCmd = &cobra.Command{
	Use:   "describe-cache-parameters",
	Short: "Returns the detailed parameter list for a particular cache parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeCacheParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeCacheParametersCmd).Standalone()

		elasticache_describeCacheParametersCmd.Flags().String("cache-parameter-group-name", "", "The name of a specific cache parameter group to return details for.")
		elasticache_describeCacheParametersCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeCacheParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		elasticache_describeCacheParametersCmd.Flags().String("source", "", "The parameter types to return.")
		elasticache_describeCacheParametersCmd.MarkFlagRequired("cache-parameter-group-name")
	})
	elasticacheCmd.AddCommand(elasticache_describeCacheParametersCmd)
}
