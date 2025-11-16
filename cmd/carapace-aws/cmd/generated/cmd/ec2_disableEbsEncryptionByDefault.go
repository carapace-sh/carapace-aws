package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableEbsEncryptionByDefaultCmd = &cobra.Command{
	Use:   "disable-ebs-encryption-by-default",
	Short: "Disables EBS encryption by default for your account in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableEbsEncryptionByDefaultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disableEbsEncryptionByDefaultCmd).Standalone()

		ec2_disableEbsEncryptionByDefaultCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableEbsEncryptionByDefaultCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableEbsEncryptionByDefaultCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disableEbsEncryptionByDefaultCmd)
}
