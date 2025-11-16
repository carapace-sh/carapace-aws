package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getEbsEncryptionByDefaultCmd = &cobra.Command{
	Use:   "get-ebs-encryption-by-default",
	Short: "Describes whether EBS encryption by default is enabled for your account in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getEbsEncryptionByDefaultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getEbsEncryptionByDefaultCmd).Standalone()

		ec2_getEbsEncryptionByDefaultCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getEbsEncryptionByDefaultCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getEbsEncryptionByDefaultCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getEbsEncryptionByDefaultCmd)
}
