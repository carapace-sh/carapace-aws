package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateIamInstanceProfileCmd = &cobra.Command{
	Use:   "disassociate-iam-instance-profile",
	Short: "Disassociates an IAM instance profile from a running or stopped instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateIamInstanceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateIamInstanceProfileCmd).Standalone()

		ec2_disassociateIamInstanceProfileCmd.Flags().String("association-id", "", "The ID of the IAM instance profile association.")
		ec2_disassociateIamInstanceProfileCmd.MarkFlagRequired("association-id")
	})
	ec2Cmd.AddCommand(ec2_disassociateIamInstanceProfileCmd)
}
