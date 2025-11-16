package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_replaceIamInstanceProfileAssociationCmd = &cobra.Command{
	Use:   "replace-iam-instance-profile-association",
	Short: "Replaces an IAM instance profile for the specified running instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_replaceIamInstanceProfileAssociationCmd).Standalone()

	ec2_replaceIamInstanceProfileAssociationCmd.Flags().String("association-id", "", "The ID of the existing IAM instance profile association.")
	ec2_replaceIamInstanceProfileAssociationCmd.Flags().String("iam-instance-profile", "", "The IAM instance profile.")
	ec2_replaceIamInstanceProfileAssociationCmd.MarkFlagRequired("association-id")
	ec2_replaceIamInstanceProfileAssociationCmd.MarkFlagRequired("iam-instance-profile")
	ec2Cmd.AddCommand(ec2_replaceIamInstanceProfileAssociationCmd)
}
