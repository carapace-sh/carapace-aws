package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createRestoreImageTaskCmd = &cobra.Command{
	Use:   "create-restore-image-task",
	Short: "Starts a task that restores an AMI from an Amazon S3 object that was previously created by using [CreateStoreImageTask](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateStoreImageTask.html).\n\nTo use this API, you must have the required permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createRestoreImageTaskCmd).Standalone()

	ec2_createRestoreImageTaskCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket that contains the stored AMI object.")
	ec2_createRestoreImageTaskCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createRestoreImageTaskCmd.Flags().String("name", "", "The name for the restored AMI.")
	ec2_createRestoreImageTaskCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createRestoreImageTaskCmd.Flags().String("object-key", "", "The name of the stored AMI object in the bucket.")
	ec2_createRestoreImageTaskCmd.Flags().String("tag-specifications", "", "The tags to apply to the AMI and snapshots on restoration.")
	ec2_createRestoreImageTaskCmd.MarkFlagRequired("bucket")
	ec2_createRestoreImageTaskCmd.Flag("no-dry-run").Hidden = true
	ec2_createRestoreImageTaskCmd.MarkFlagRequired("object-key")
	ec2Cmd.AddCommand(ec2_createRestoreImageTaskCmd)
}
