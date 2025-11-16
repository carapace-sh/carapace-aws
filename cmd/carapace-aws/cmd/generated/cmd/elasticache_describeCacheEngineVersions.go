package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeCacheEngineVersionsCmd = &cobra.Command{
	Use:   "describe-cache-engine-versions",
	Short: "Returns a list of the available cache engines and their versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeCacheEngineVersionsCmd).Standalone()

	elasticache_describeCacheEngineVersionsCmd.Flags().String("cache-parameter-group-family", "", "The name of a specific cache parameter group family to return details for.")
	elasticache_describeCacheEngineVersionsCmd.Flags().Bool("default-only", false, "If `true`, specifies that only the default version of the specified engine or engine and major version combination is to be returned.")
	elasticache_describeCacheEngineVersionsCmd.Flags().String("engine", "", "The cache engine to return.")
	elasticache_describeCacheEngineVersionsCmd.Flags().String("engine-version", "", "The cache engine version to return.")
	elasticache_describeCacheEngineVersionsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
	elasticache_describeCacheEngineVersionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	elasticache_describeCacheEngineVersionsCmd.Flags().Bool("no-default-only", false, "If `true`, specifies that only the default version of the specified engine or engine and major version combination is to be returned.")
	elasticache_describeCacheEngineVersionsCmd.Flag("no-default-only").Hidden = true
	elasticacheCmd.AddCommand(elasticache_describeCacheEngineVersionsCmd)
}
