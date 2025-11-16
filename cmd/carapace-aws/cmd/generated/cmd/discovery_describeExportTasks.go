package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_describeExportTasksCmd = &cobra.Command{
	Use:   "describe-export-tasks",
	Short: "Retrieve status of one or more export tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_describeExportTasksCmd).Standalone()

	discovery_describeExportTasksCmd.Flags().String("export-ids", "", "One or more unique identifiers used to query the status of an export request.")
	discovery_describeExportTasksCmd.Flags().String("filters", "", "One or more filters.")
	discovery_describeExportTasksCmd.Flags().String("max-results", "", "The maximum number of volume results returned by `DescribeExportTasks` in paginated output.")
	discovery_describeExportTasksCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeExportTasks` request where `maxResults` was used and the results exceeded the value of that parameter.")
	discoveryCmd.AddCommand(discovery_describeExportTasksCmd)
}
