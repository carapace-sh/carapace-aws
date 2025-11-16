package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_modifyCacheParameterGroupCmd = &cobra.Command{
	Use:   "modify-cache-parameter-group",
	Short: "Modifies the parameters of a cache parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_modifyCacheParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_modifyCacheParameterGroupCmd).Standalone()

		elasticache_modifyCacheParameterGroupCmd.Flags().String("cache-parameter-group-name", "", "The name of the cache parameter group to modify.")
		elasticache_modifyCacheParameterGroupCmd.Flags().String("parameter-name-values", "", "An array of parameter names and values for the parameter update.")
		elasticache_modifyCacheParameterGroupCmd.MarkFlagRequired("cache-parameter-group-name")
		elasticache_modifyCacheParameterGroupCmd.MarkFlagRequired("parameter-name-values")
	})
	elasticacheCmd.AddCommand(elasticache_modifyCacheParameterGroupCmd)
}
