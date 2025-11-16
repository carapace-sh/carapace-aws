package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateEnclaveCertificateIamRoleCmd = &cobra.Command{
	Use:   "associate-enclave-certificate-iam-role",
	Short: "Associates an Identity and Access Management (IAM) role with an Certificate Manager (ACM) certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateEnclaveCertificateIamRoleCmd).Standalone()

	ec2_associateEnclaveCertificateIamRoleCmd.Flags().String("certificate-arn", "", "The ARN of the ACM certificate with which to associate the IAM role.")
	ec2_associateEnclaveCertificateIamRoleCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateEnclaveCertificateIamRoleCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateEnclaveCertificateIamRoleCmd.Flags().String("role-arn", "", "The ARN of the IAM role to associate with the ACM certificate.")
	ec2_associateEnclaveCertificateIamRoleCmd.MarkFlagRequired("certificate-arn")
	ec2_associateEnclaveCertificateIamRoleCmd.Flag("no-dry-run").Hidden = true
	ec2_associateEnclaveCertificateIamRoleCmd.MarkFlagRequired("role-arn")
	ec2Cmd.AddCommand(ec2_associateEnclaveCertificateIamRoleCmd)
}
