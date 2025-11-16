package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourcegroupstaggingapi_describeReportCreationCmd = &cobra.Command{
	Use:   "describe-report-creation",
	Short: "Describes the status of the `StartReportCreation` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourcegroupstaggingapi_describeReportCreationCmd).Standalone()

	resourcegroupstaggingapiCmd.AddCommand(resourcegroupstaggingapi_describeReportCreationCmd)
}
