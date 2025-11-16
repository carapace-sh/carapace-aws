package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVolumesCmd = &cobra.Command{
	Use:   "describe-volumes",
	Short: "Describes the specified EBS volumes or all of your EBS volumes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVolumesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVolumesCmd).Standalone()

		ec2_describeVolumesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVolumesCmd.Flags().String("filters", "", "The filters.")
		ec2_describeVolumesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeVolumesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeVolumesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVolumesCmd.Flags().String("volume-ids", "", "The volume IDs.")
		ec2_describeVolumesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVolumesCmd)
}
