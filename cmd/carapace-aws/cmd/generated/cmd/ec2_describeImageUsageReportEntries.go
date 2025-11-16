package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeImageUsageReportEntriesCmd = &cobra.Command{
	Use:   "describe-image-usage-report-entries",
	Short: "Describes the entries in image usage reports, showing how your images are used across other Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeImageUsageReportEntriesCmd).Standalone()

	ec2_describeImageUsageReportEntriesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImageUsageReportEntriesCmd.Flags().String("filters", "", "The filters.")
	ec2_describeImageUsageReportEntriesCmd.Flags().String("image-ids", "", "The IDs of the images for filtering the report entries.")
	ec2_describeImageUsageReportEntriesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeImageUsageReportEntriesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeImageUsageReportEntriesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImageUsageReportEntriesCmd.Flags().String("report-ids", "", "The IDs of the usage reports.")
	ec2_describeImageUsageReportEntriesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeImageUsageReportEntriesCmd)
}
