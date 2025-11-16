package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createImageUsageReportCmd = &cobra.Command{
	Use:   "create-image-usage-report",
	Short: "Creates a report that shows how your image is used across other Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createImageUsageReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createImageUsageReportCmd).Standalone()

		ec2_createImageUsageReportCmd.Flags().String("account-ids", "", "The Amazon Web Services account IDs to include in the report.")
		ec2_createImageUsageReportCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure idempotency of the request.")
		ec2_createImageUsageReportCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createImageUsageReportCmd.Flags().String("image-id", "", "The ID of the image to report on.")
		ec2_createImageUsageReportCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createImageUsageReportCmd.Flags().String("resource-types", "", "The resource types to include in the report.")
		ec2_createImageUsageReportCmd.Flags().String("tag-specifications", "", "The tags to apply to the report on creation.")
		ec2_createImageUsageReportCmd.MarkFlagRequired("image-id")
		ec2_createImageUsageReportCmd.Flag("no-dry-run").Hidden = true
		ec2_createImageUsageReportCmd.MarkFlagRequired("resource-types")
	})
	ec2Cmd.AddCommand(ec2_createImageUsageReportCmd)
}
