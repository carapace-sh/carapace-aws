package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_describeImportTasksCmd = &cobra.Command{
	Use:   "describe-import-tasks",
	Short: "Returns an array of import tasks for your account, including status information, times, IDs, the Amazon S3 Object URL for the import file, and more.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_describeImportTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_describeImportTasksCmd).Standalone()

		discovery_describeImportTasksCmd.Flags().String("filters", "", "An array of name-value pairs that you provide to filter the results for the `DescribeImportTask` request to a specific subset of results.")
		discovery_describeImportTasksCmd.Flags().String("max-results", "", "The maximum number of results that you want this request to return, up to 100.")
		discovery_describeImportTasksCmd.Flags().String("next-token", "", "The token to request a specific page of results.")
	})
	discoveryCmd.AddCommand(discovery_describeImportTasksCmd)
}
