package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_startExportTaskCmd = &cobra.Command{
	Use:   "start-export-task",
	Short: "Begins the export of a discovered data report to an Amazon S3 bucket managed by Amazon Web Services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_startExportTaskCmd).Standalone()

	discovery_startExportTaskCmd.Flags().String("end-time", "", "The end timestamp for exported data from the single Application Discovery Agent selected in the filters.")
	discovery_startExportTaskCmd.Flags().String("export-data-format", "", "The file format for the returned export data.")
	discovery_startExportTaskCmd.Flags().String("filters", "", "If a filter is present, it selects the single `agentId` of the Application Discovery Agent for which data is exported.")
	discovery_startExportTaskCmd.Flags().String("preferences", "", "Indicates the type of data that needs to be exported.")
	discovery_startExportTaskCmd.Flags().String("start-time", "", "The start timestamp for exported data from the single Application Discovery Agent selected in the filters.")
	discoveryCmd.AddCommand(discovery_startExportTaskCmd)
}
