package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_describeContinuousExportsCmd = &cobra.Command{
	Use:   "describe-continuous-exports",
	Short: "Lists exports as specified by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_describeContinuousExportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_describeContinuousExportsCmd).Standalone()

		discovery_describeContinuousExportsCmd.Flags().String("export-ids", "", "The unique IDs assigned to the exports.")
		discovery_describeContinuousExportsCmd.Flags().String("max-results", "", "A number between 1 and 100 specifying the maximum number of continuous export descriptions returned.")
		discovery_describeContinuousExportsCmd.Flags().String("next-token", "", "The token from the previous call to `DescribeExportTasks`.")
	})
	discoveryCmd.AddCommand(discovery_describeContinuousExportsCmd)
}
