package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createDelegateMacVolumeOwnershipTaskCmd = &cobra.Command{
	Use:   "create-delegate-mac-volume-ownership-task",
	Short: "Delegates ownership of the Amazon EBS root volume for an Apple silicon Mac instance to an administrative user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createDelegateMacVolumeOwnershipTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createDelegateMacVolumeOwnershipTaskCmd).Standalone()

		ec2_createDelegateMacVolumeOwnershipTaskCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createDelegateMacVolumeOwnershipTaskCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createDelegateMacVolumeOwnershipTaskCmd.Flags().String("instance-id", "", "The ID of the Amazon EC2 Mac instance.")
		ec2_createDelegateMacVolumeOwnershipTaskCmd.Flags().String("mac-credentials", "", "Specifies the following credentials:")
		ec2_createDelegateMacVolumeOwnershipTaskCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createDelegateMacVolumeOwnershipTaskCmd.Flags().String("tag-specifications", "", "The tags to assign to the volume ownership delegation task.")
		ec2_createDelegateMacVolumeOwnershipTaskCmd.MarkFlagRequired("instance-id")
		ec2_createDelegateMacVolumeOwnershipTaskCmd.MarkFlagRequired("mac-credentials")
		ec2_createDelegateMacVolumeOwnershipTaskCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createDelegateMacVolumeOwnershipTaskCmd)
}
