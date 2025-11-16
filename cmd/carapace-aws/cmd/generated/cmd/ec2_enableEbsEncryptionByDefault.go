package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableEbsEncryptionByDefaultCmd = &cobra.Command{
	Use:   "enable-ebs-encryption-by-default",
	Short: "Enables EBS encryption by default for your account in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableEbsEncryptionByDefaultCmd).Standalone()

	ec2_enableEbsEncryptionByDefaultCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableEbsEncryptionByDefaultCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableEbsEncryptionByDefaultCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_enableEbsEncryptionByDefaultCmd)
}
