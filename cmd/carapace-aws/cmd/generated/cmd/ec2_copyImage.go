package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_copyImageCmd = &cobra.Command{
	Use:   "copy-image",
	Short: "Initiates an AMI copy operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_copyImageCmd).Standalone()

	ec2_copyImageCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
	ec2_copyImageCmd.Flags().Bool("copy-image-tags", false, "Specifies whether to copy your user-defined AMI tags to the new AMI.")
	ec2_copyImageCmd.Flags().String("description", "", "A description for the new AMI.")
	ec2_copyImageCmd.Flags().String("destination-availability-zone", "", "The Local Zone for the new AMI (for example, `cn-north-1-pkx-1a`).")
	ec2_copyImageCmd.Flags().String("destination-availability-zone-id", "", "The ID of the Local Zone for the new AMI (for example, `cnn1-pkx1-az1`).")
	ec2_copyImageCmd.Flags().String("destination-outpost-arn", "", "The Amazon Resource Name (ARN) of the Outpost for the new AMI.")
	ec2_copyImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_copyImageCmd.Flags().Bool("encrypted", false, "Specifies whether to encrypt the snapshots of the copied image.")
	ec2_copyImageCmd.Flags().String("kms-key-id", "", "The identifier of the symmetric Key Management Service (KMS) KMS key to use when creating encrypted volumes.")
	ec2_copyImageCmd.Flags().String("name", "", "The name of the new AMI.")
	ec2_copyImageCmd.Flags().Bool("no-copy-image-tags", false, "Specifies whether to copy your user-defined AMI tags to the new AMI.")
	ec2_copyImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_copyImageCmd.Flags().Bool("no-encrypted", false, "Specifies whether to encrypt the snapshots of the copied image.")
	ec2_copyImageCmd.Flags().String("snapshot-copy-completion-duration-minutes", "", "Specify a completion duration, in 15 minute increments, to initiate a time-based AMI copy.")
	ec2_copyImageCmd.Flags().String("source-image-id", "", "The ID of the AMI to copy.")
	ec2_copyImageCmd.Flags().String("source-region", "", "The name of the Region that contains the AMI to copy.")
	ec2_copyImageCmd.Flags().String("tag-specifications", "", "The tags to apply to the new AMI and new snapshots.")
	ec2_copyImageCmd.MarkFlagRequired("name")
	ec2_copyImageCmd.Flag("no-copy-image-tags").Hidden = true
	ec2_copyImageCmd.Flag("no-dry-run").Hidden = true
	ec2_copyImageCmd.Flag("no-encrypted").Hidden = true
	ec2_copyImageCmd.MarkFlagRequired("source-image-id")
	ec2_copyImageCmd.MarkFlagRequired("source-region")
	ec2Cmd.AddCommand(ec2_copyImageCmd)
}
