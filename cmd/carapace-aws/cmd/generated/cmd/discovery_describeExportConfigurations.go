package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_describeExportConfigurationsCmd = &cobra.Command{
	Use:   "describe-export-configurations",
	Short: "`DescribeExportConfigurations` is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_describeExportConfigurationsCmd).Standalone()

	discovery_describeExportConfigurationsCmd.Flags().String("export-ids", "", "A list of continuous export IDs to search for.")
	discovery_describeExportConfigurationsCmd.Flags().String("max-results", "", "A number between 1 and 100 specifying the maximum number of continuous export descriptions returned.")
	discovery_describeExportConfigurationsCmd.Flags().String("next-token", "", "The token from the previous call to describe-export-tasks.")
	discoveryCmd.AddCommand(discovery_describeExportConfigurationsCmd)
}
