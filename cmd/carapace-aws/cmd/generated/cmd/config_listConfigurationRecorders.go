package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listConfigurationRecordersCmd = &cobra.Command{
	Use:   "list-configuration-recorders",
	Short: "Returns a list of configuration recorders depending on the filters you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listConfigurationRecordersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_listConfigurationRecordersCmd).Standalone()

		config_listConfigurationRecordersCmd.Flags().String("filters", "", "Filters the results based on a list of `ConfigurationRecorderFilter` objects that you specify.")
		config_listConfigurationRecordersCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		config_listConfigurationRecordersCmd.Flags().String("next-token", "", "The `NextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	})
	configCmd.AddCommand(config_listConfigurationRecordersCmd)
}
