package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeEngineDefaultParametersCmd = &cobra.Command{
	Use:   "describe-engine-default-parameters",
	Short: "Returns the default engine and system parameter information for the specified cache engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeEngineDefaultParametersCmd).Standalone()

	elasticache_describeEngineDefaultParametersCmd.Flags().String("cache-parameter-group-family", "", "The name of the cache parameter group family.")
	elasticache_describeEngineDefaultParametersCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
	elasticache_describeEngineDefaultParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	elasticache_describeEngineDefaultParametersCmd.MarkFlagRequired("cache-parameter-group-family")
	elasticacheCmd.AddCommand(elasticache_describeEngineDefaultParametersCmd)
}
