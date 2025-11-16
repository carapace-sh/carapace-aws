package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteImageUsageReportCmd = &cobra.Command{
	Use:   "delete-image-usage-report",
	Short: "Deletes the specified image usage report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteImageUsageReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteImageUsageReportCmd).Standalone()

		ec2_deleteImageUsageReportCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteImageUsageReportCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteImageUsageReportCmd.Flags().String("report-id", "", "The ID of the report to delete.")
		ec2_deleteImageUsageReportCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteImageUsageReportCmd.MarkFlagRequired("report-id")
	})
	ec2Cmd.AddCommand(ec2_deleteImageUsageReportCmd)
}
