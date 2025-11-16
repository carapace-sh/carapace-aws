package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyEbsDefaultKmsKeyIdCmd = &cobra.Command{
	Use:   "modify-ebs-default-kms-key-id",
	Short: "Changes the default KMS key for EBS encryption by default for your account in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyEbsDefaultKmsKeyIdCmd).Standalone()

	ec2_modifyEbsDefaultKmsKeyIdCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyEbsDefaultKmsKeyIdCmd.Flags().String("kms-key-id", "", "The identifier of the KMS key to use for Amazon EBS encryption.")
	ec2_modifyEbsDefaultKmsKeyIdCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyEbsDefaultKmsKeyIdCmd.MarkFlagRequired("kms-key-id")
	ec2_modifyEbsDefaultKmsKeyIdCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyEbsDefaultKmsKeyIdCmd)
}
