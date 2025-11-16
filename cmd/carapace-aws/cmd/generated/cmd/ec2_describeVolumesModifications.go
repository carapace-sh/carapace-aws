package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVolumesModificationsCmd = &cobra.Command{
	Use:   "describe-volumes-modifications",
	Short: "Describes the most recent volume modification request for the specified EBS volumes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVolumesModificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVolumesModificationsCmd).Standalone()

		ec2_describeVolumesModificationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVolumesModificationsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeVolumesModificationsCmd.Flags().String("max-results", "", "The maximum number of results (up to a limit of 500) to be returned in a paginated request.")
		ec2_describeVolumesModificationsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeVolumesModificationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVolumesModificationsCmd.Flags().String("volume-ids", "", "The IDs of the volumes.")
		ec2_describeVolumesModificationsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVolumesModificationsCmd)
}
