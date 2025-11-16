package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeImageUsageReportsCmd = &cobra.Command{
	Use:   "describe-image-usage-reports",
	Short: "Describes the configuration and status of image usage reports, filtered by report IDs or image IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeImageUsageReportsCmd).Standalone()

	ec2_describeImageUsageReportsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImageUsageReportsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeImageUsageReportsCmd.Flags().String("image-ids", "", "The IDs of the images for filtering the reports.")
	ec2_describeImageUsageReportsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeImageUsageReportsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeImageUsageReportsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImageUsageReportsCmd.Flags().String("report-ids", "", "The IDs of the image usage reports.")
	ec2_describeImageUsageReportsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeImageUsageReportsCmd)
}
