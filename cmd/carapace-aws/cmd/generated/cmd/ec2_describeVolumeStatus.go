package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVolumeStatusCmd = &cobra.Command{
	Use:   "describe-volume-status",
	Short: "Describes the status of the specified volumes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVolumeStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVolumeStatusCmd).Standalone()

		ec2_describeVolumeStatusCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVolumeStatusCmd.Flags().String("filters", "", "The filters.")
		ec2_describeVolumeStatusCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeVolumeStatusCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeVolumeStatusCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVolumeStatusCmd.Flags().String("volume-ids", "", "The IDs of the volumes.")
		ec2_describeVolumeStatusCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVolumeStatusCmd)
}
