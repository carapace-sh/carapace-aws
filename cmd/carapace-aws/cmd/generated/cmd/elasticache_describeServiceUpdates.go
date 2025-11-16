package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeServiceUpdatesCmd = &cobra.Command{
	Use:   "describe-service-updates",
	Short: "Returns details of the service updates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeServiceUpdatesCmd).Standalone()

	elasticache_describeServiceUpdatesCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
	elasticache_describeServiceUpdatesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response")
	elasticache_describeServiceUpdatesCmd.Flags().String("service-update-name", "", "The unique ID of the service update")
	elasticache_describeServiceUpdatesCmd.Flags().String("service-update-status", "", "The status of the service update")
	elasticacheCmd.AddCommand(elasticache_describeServiceUpdatesCmd)
}
