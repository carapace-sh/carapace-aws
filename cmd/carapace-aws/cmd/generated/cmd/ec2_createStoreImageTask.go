package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createStoreImageTaskCmd = &cobra.Command{
	Use:   "create-store-image-task",
	Short: "Stores an AMI as a single object in an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createStoreImageTaskCmd).Standalone()

	ec2_createStoreImageTaskCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket in which the AMI object will be stored.")
	ec2_createStoreImageTaskCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createStoreImageTaskCmd.Flags().String("image-id", "", "The ID of the AMI.")
	ec2_createStoreImageTaskCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createStoreImageTaskCmd.Flags().String("s3-object-tags", "", "The tags to apply to the AMI object that will be stored in the Amazon S3 bucket.")
	ec2_createStoreImageTaskCmd.MarkFlagRequired("bucket")
	ec2_createStoreImageTaskCmd.MarkFlagRequired("image-id")
	ec2_createStoreImageTaskCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createStoreImageTaskCmd)
}
