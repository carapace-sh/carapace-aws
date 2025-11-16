package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_resetCacheParameterGroupCmd = &cobra.Command{
	Use:   "reset-cache-parameter-group",
	Short: "Modifies the parameters of a cache parameter group to the engine or system default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_resetCacheParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_resetCacheParameterGroupCmd).Standalone()

		elasticache_resetCacheParameterGroupCmd.Flags().String("cache-parameter-group-name", "", "The name of the cache parameter group to reset.")
		elasticache_resetCacheParameterGroupCmd.Flags().Bool("no-reset-all-parameters", false, "If `true`, all parameters in the cache parameter group are reset to their default values.")
		elasticache_resetCacheParameterGroupCmd.Flags().String("parameter-name-values", "", "An array of parameter names to reset to their default values.")
		elasticache_resetCacheParameterGroupCmd.Flags().Bool("reset-all-parameters", false, "If `true`, all parameters in the cache parameter group are reset to their default values.")
		elasticache_resetCacheParameterGroupCmd.MarkFlagRequired("cache-parameter-group-name")
		elasticache_resetCacheParameterGroupCmd.Flag("no-reset-all-parameters").Hidden = true
	})
	elasticacheCmd.AddCommand(elasticache_resetCacheParameterGroupCmd)
}
