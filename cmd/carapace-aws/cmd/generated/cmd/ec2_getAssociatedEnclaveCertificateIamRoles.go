package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getAssociatedEnclaveCertificateIamRolesCmd = &cobra.Command{
	Use:   "get-associated-enclave-certificate-iam-roles",
	Short: "Returns the IAM roles that are associated with the specified ACM (ACM) certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getAssociatedEnclaveCertificateIamRolesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getAssociatedEnclaveCertificateIamRolesCmd).Standalone()

		ec2_getAssociatedEnclaveCertificateIamRolesCmd.Flags().String("certificate-arn", "", "The ARN of the ACM certificate for which to view the associated IAM roles, encryption keys, and Amazon S3 object information.")
		ec2_getAssociatedEnclaveCertificateIamRolesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getAssociatedEnclaveCertificateIamRolesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getAssociatedEnclaveCertificateIamRolesCmd.MarkFlagRequired("certificate-arn")
		ec2_getAssociatedEnclaveCertificateIamRolesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getAssociatedEnclaveCertificateIamRolesCmd)
}
