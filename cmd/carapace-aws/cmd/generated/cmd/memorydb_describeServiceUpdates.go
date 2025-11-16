package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeServiceUpdatesCmd = &cobra.Command{
	Use:   "describe-service-updates",
	Short: "Returns details of the service updates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeServiceUpdatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_describeServiceUpdatesCmd).Standalone()

		memorydb_describeServiceUpdatesCmd.Flags().String("cluster-names", "", "The list of cluster names to identify service updates to apply.")
		memorydb_describeServiceUpdatesCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
		memorydb_describeServiceUpdatesCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
		memorydb_describeServiceUpdatesCmd.Flags().String("service-update-name", "", "The unique ID of the service update to describe.")
		memorydb_describeServiceUpdatesCmd.Flags().String("status", "", "The status(es) of the service updates to filter on.")
	})
	memorydbCmd.AddCommand(memorydb_describeServiceUpdatesCmd)
}
