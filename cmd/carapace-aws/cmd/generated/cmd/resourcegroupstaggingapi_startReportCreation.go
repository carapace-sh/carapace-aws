package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourcegroupstaggingapi_startReportCreationCmd = &cobra.Command{
	Use:   "start-report-creation",
	Short: "Generates a report that lists all tagged resources in the accounts across your organization and tells whether each resource is compliant with the effective tag policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourcegroupstaggingapi_startReportCreationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourcegroupstaggingapi_startReportCreationCmd).Standalone()

		resourcegroupstaggingapi_startReportCreationCmd.Flags().String("s3-bucket", "", "The name of the Amazon S3 bucket where the report will be stored; for example:")
		resourcegroupstaggingapi_startReportCreationCmd.MarkFlagRequired("s3-bucket")
	})
	resourcegroupstaggingapiCmd.AddCommand(resourcegroupstaggingapi_startReportCreationCmd)
}
