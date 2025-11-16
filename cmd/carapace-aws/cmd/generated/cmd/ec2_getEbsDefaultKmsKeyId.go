package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getEbsDefaultKmsKeyIdCmd = &cobra.Command{
	Use:   "get-ebs-default-kms-key-id",
	Short: "Describes the default KMS key for EBS encryption by default for your account in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getEbsDefaultKmsKeyIdCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getEbsDefaultKmsKeyIdCmd).Standalone()

		ec2_getEbsDefaultKmsKeyIdCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getEbsDefaultKmsKeyIdCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getEbsDefaultKmsKeyIdCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getEbsDefaultKmsKeyIdCmd)
}
