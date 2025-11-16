package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_resetEbsDefaultKmsKeyIdCmd = &cobra.Command{
	Use:   "reset-ebs-default-kms-key-id",
	Short: "Resets the default KMS key for EBS encryption for your account in this Region to the Amazon Web Services managed KMS key for EBS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_resetEbsDefaultKmsKeyIdCmd).Standalone()

	ec2_resetEbsDefaultKmsKeyIdCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_resetEbsDefaultKmsKeyIdCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_resetEbsDefaultKmsKeyIdCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_resetEbsDefaultKmsKeyIdCmd)
}
