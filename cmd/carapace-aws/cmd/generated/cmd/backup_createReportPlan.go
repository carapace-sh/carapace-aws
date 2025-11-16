package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createReportPlanCmd = &cobra.Command{
	Use:   "create-report-plan",
	Short: "Creates a report plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createReportPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_createReportPlanCmd).Standalone()

		backup_createReportPlanCmd.Flags().String("idempotency-token", "", "A customer-chosen string that you can use to distinguish between otherwise identical calls to `CreateReportPlanInput`.")
		backup_createReportPlanCmd.Flags().String("report-delivery-channel", "", "A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.")
		backup_createReportPlanCmd.Flags().String("report-plan-description", "", "An optional description of the report plan with a maximum of 1,024 characters.")
		backup_createReportPlanCmd.Flags().String("report-plan-name", "", "The unique name of the report plan.")
		backup_createReportPlanCmd.Flags().String("report-plan-tags", "", "The tags to assign to the report plan.")
		backup_createReportPlanCmd.Flags().String("report-setting", "", "Identifies the report template for the report.")
		backup_createReportPlanCmd.MarkFlagRequired("report-delivery-channel")
		backup_createReportPlanCmd.MarkFlagRequired("report-plan-name")
		backup_createReportPlanCmd.MarkFlagRequired("report-setting")
	})
	backupCmd.AddCommand(backup_createReportPlanCmd)
}
