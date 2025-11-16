package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createReplaceRootVolumeTaskCmd = &cobra.Command{
	Use:   "create-replace-root-volume-task",
	Short: "Replaces the EBS-backed root volume for a `running` instance with a new volume that is restored to the original root volume's launch state, that is restored to a specific snapshot taken from the original root volume, or that is restored from an AMI that has the same key characteristics as that of the instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createReplaceRootVolumeTaskCmd).Standalone()

	ec2_createReplaceRootVolumeTaskCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
	ec2_createReplaceRootVolumeTaskCmd.Flags().Bool("delete-replaced-root-volume", false, "Indicates whether to automatically delete the original root volume after the root volume replacement task completes.")
	ec2_createReplaceRootVolumeTaskCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createReplaceRootVolumeTaskCmd.Flags().String("image-id", "", "The ID of the AMI to use to restore the root volume.")
	ec2_createReplaceRootVolumeTaskCmd.Flags().String("instance-id", "", "The ID of the instance for which to replace the root volume.")
	ec2_createReplaceRootVolumeTaskCmd.Flags().Bool("no-delete-replaced-root-volume", false, "Indicates whether to automatically delete the original root volume after the root volume replacement task completes.")
	ec2_createReplaceRootVolumeTaskCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createReplaceRootVolumeTaskCmd.Flags().String("snapshot-id", "", "The ID of the snapshot from which to restore the replacement root volume.")
	ec2_createReplaceRootVolumeTaskCmd.Flags().String("tag-specifications", "", "The tags to apply to the root volume replacement task.")
	ec2_createReplaceRootVolumeTaskCmd.Flags().String("volume-initialization-rate", "", "Specifies the Amazon EBS Provisioned Rate for Volume Initialization (volume initialization rate), in MiB/s, at which to download the snapshot blocks from Amazon S3 to the replacement root volume.")
	ec2_createReplaceRootVolumeTaskCmd.MarkFlagRequired("instance-id")
	ec2_createReplaceRootVolumeTaskCmd.Flag("no-delete-replaced-root-volume").Hidden = true
	ec2_createReplaceRootVolumeTaskCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createReplaceRootVolumeTaskCmd)
}
