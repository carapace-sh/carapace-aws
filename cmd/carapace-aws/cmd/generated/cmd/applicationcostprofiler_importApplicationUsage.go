package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationcostprofiler_importApplicationUsageCmd = &cobra.Command{
	Use:   "import-application-usage",
	Short: "Ingests application usage data from Amazon Simple Storage Service (Amazon S3).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationcostprofiler_importApplicationUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationcostprofiler_importApplicationUsageCmd).Standalone()

		applicationcostprofiler_importApplicationUsageCmd.Flags().String("source-s3-location", "", "Amazon S3 location to import application usage data from.")
		applicationcostprofiler_importApplicationUsageCmd.MarkFlagRequired("source-s3-location")
	})
	applicationcostprofilerCmd.AddCommand(applicationcostprofiler_importApplicationUsageCmd)
}
