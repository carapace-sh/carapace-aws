package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateEnclaveCertificateIamRoleCmd = &cobra.Command{
	Use:   "disassociate-enclave-certificate-iam-role",
	Short: "Disassociates an IAM role from an Certificate Manager (ACM) certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateEnclaveCertificateIamRoleCmd).Standalone()

	ec2_disassociateEnclaveCertificateIamRoleCmd.Flags().String("certificate-arn", "", "The ARN of the ACM certificate from which to disassociate the IAM role.")
	ec2_disassociateEnclaveCertificateIamRoleCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateEnclaveCertificateIamRoleCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateEnclaveCertificateIamRoleCmd.Flags().String("role-arn", "", "The ARN of the IAM role to disassociate.")
	ec2_disassociateEnclaveCertificateIamRoleCmd.MarkFlagRequired("certificate-arn")
	ec2_disassociateEnclaveCertificateIamRoleCmd.Flag("no-dry-run").Hidden = true
	ec2_disassociateEnclaveCertificateIamRoleCmd.MarkFlagRequired("role-arn")
	ec2Cmd.AddCommand(ec2_disassociateEnclaveCertificateIamRoleCmd)
}
