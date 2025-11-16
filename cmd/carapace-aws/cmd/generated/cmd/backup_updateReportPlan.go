package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_updateReportPlanCmd = &cobra.Command{
	Use:   "update-report-plan",
	Short: "Updates the specified report plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_updateReportPlanCmd).Standalone()

	backup_updateReportPlanCmd.Flags().String("idempotency-token", "", "A customer-chosen string that you can use to distinguish between otherwise identical calls to `UpdateReportPlanInput`.")
	backup_updateReportPlanCmd.Flags().String("report-delivery-channel", "", "The information about where to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.")
	backup_updateReportPlanCmd.Flags().String("report-plan-description", "", "An optional description of the report plan with a maximum 1,024 characters.")
	backup_updateReportPlanCmd.Flags().String("report-plan-name", "", "The unique name of the report plan.")
	backup_updateReportPlanCmd.Flags().String("report-setting", "", "The report template for the report.")
	backup_updateReportPlanCmd.MarkFlagRequired("report-plan-name")
	backupCmd.AddCommand(backup_updateReportPlanCmd)
}
