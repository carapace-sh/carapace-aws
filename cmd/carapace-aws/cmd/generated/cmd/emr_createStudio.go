package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_createStudioCmd = &cobra.Command{
	Use:   "create-studio",
	Short: "Creates a new Amazon EMR Studio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_createStudioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_createStudioCmd).Standalone()

		emr_createStudioCmd.Flags().String("auth-mode", "", "Specifies whether the Studio authenticates users using IAM or IAM Identity Center.")
		emr_createStudioCmd.Flags().String("default-s3-location", "", "The Amazon S3 location to back up Amazon EMR Studio Workspaces and notebook files.")
		emr_createStudioCmd.Flags().String("description", "", "A detailed description of the Amazon EMR Studio.")
		emr_createStudioCmd.Flags().String("encryption-key-arn", "", "The KMS key identifier (ARN) used to encrypt Amazon EMR Studio workspace and notebook files when backed up to Amazon S3.")
		emr_createStudioCmd.Flags().String("engine-security-group-id", "", "The ID of the Amazon EMR Studio Engine security group.")
		emr_createStudioCmd.Flags().String("idc-instance-arn", "", "The ARN of the IAM Identity Center instance to create the Studio application.")
		emr_createStudioCmd.Flags().String("idc-user-assignment", "", "Specifies whether IAM Identity Center user assignment is `REQUIRED` or `OPTIONAL`.")
		emr_createStudioCmd.Flags().String("idp-auth-url", "", "The authentication endpoint of your identity provider (IdP).")
		emr_createStudioCmd.Flags().String("idp-relay-state-parameter-name", "", "The name that your identity provider (IdP) uses for its `RelayState` parameter.")
		emr_createStudioCmd.Flags().String("name", "", "A descriptive name for the Amazon EMR Studio.")
		emr_createStudioCmd.Flags().String("service-role", "", "The IAM role that the Amazon EMR Studio assumes.")
		emr_createStudioCmd.Flags().String("subnet-ids", "", "A list of subnet IDs to associate with the Amazon EMR Studio.")
		emr_createStudioCmd.Flags().String("tags", "", "A list of tags to associate with the Amazon EMR Studio.")
		emr_createStudioCmd.Flags().String("trusted-identity-propagation-enabled", "", "A Boolean indicating whether to enable Trusted identity propagation for the Studio.")
		emr_createStudioCmd.Flags().String("user-role", "", "The IAM user role that users and groups assume when logged in to an Amazon EMR Studio.")
		emr_createStudioCmd.Flags().String("vpc-id", "", "The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.")
		emr_createStudioCmd.Flags().String("workspace-security-group-id", "", "The ID of the Amazon EMR Studio Workspace security group.")
		emr_createStudioCmd.MarkFlagRequired("auth-mode")
		emr_createStudioCmd.MarkFlagRequired("default-s3-location")
		emr_createStudioCmd.MarkFlagRequired("engine-security-group-id")
		emr_createStudioCmd.MarkFlagRequired("name")
		emr_createStudioCmd.MarkFlagRequired("service-role")
		emr_createStudioCmd.MarkFlagRequired("subnet-ids")
		emr_createStudioCmd.MarkFlagRequired("vpc-id")
		emr_createStudioCmd.MarkFlagRequired("workspace-security-group-id")
	})
	emrCmd.AddCommand(emr_createStudioCmd)
}
