package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateIamInstanceProfileCmd = &cobra.Command{
	Use:   "associate-iam-instance-profile",
	Short: "Associates an IAM instance profile with a running or stopped instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateIamInstanceProfileCmd).Standalone()

	ec2_associateIamInstanceProfileCmd.Flags().String("iam-instance-profile", "", "The IAM instance profile.")
	ec2_associateIamInstanceProfileCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_associateIamInstanceProfileCmd.MarkFlagRequired("iam-instance-profile")
	ec2_associateIamInstanceProfileCmd.MarkFlagRequired("instance-id")
	ec2Cmd.AddCommand(ec2_associateIamInstanceProfileCmd)
}
