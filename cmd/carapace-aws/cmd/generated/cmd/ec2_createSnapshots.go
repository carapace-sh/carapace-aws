package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createSnapshotsCmd = &cobra.Command{
	Use:   "create-snapshots",
	Short: "Creates crash-consistent snapshots of multiple EBS volumes attached to an Amazon EC2 instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createSnapshotsCmd).Standalone()

		ec2_createSnapshotsCmd.Flags().String("copy-tags-from-source", "", "Copies the tags from the specified volume to corresponding snapshot.")
		ec2_createSnapshotsCmd.Flags().String("description", "", "A description propagated to every snapshot specified by the instance.")
		ec2_createSnapshotsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createSnapshotsCmd.Flags().String("instance-specification", "", "The instance to specify which volumes should be included in the snapshots.")
		ec2_createSnapshotsCmd.Flags().String("location", "", "Only supported for instances in Local Zones.")
		ec2_createSnapshotsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createSnapshotsCmd.Flags().String("outpost-arn", "", "Only supported for instances on Outposts.")
		ec2_createSnapshotsCmd.Flags().String("tag-specifications", "", "Tags to apply to every snapshot specified by the instance.")
		ec2_createSnapshotsCmd.MarkFlagRequired("instance-specification")
		ec2_createSnapshotsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createSnapshotsCmd)
}
