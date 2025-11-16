package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createPartnerAppCmd = &cobra.Command{
	Use:   "create-partner-app",
	Short: "Creates an Amazon SageMaker Partner AI App.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createPartnerAppCmd).Standalone()

	sagemaker_createPartnerAppCmd.Flags().String("application-config", "", "Configuration settings for the SageMaker Partner AI App.")
	sagemaker_createPartnerAppCmd.Flags().String("auth-type", "", "The authorization type that users use to access the SageMaker Partner AI App.")
	sagemaker_createPartnerAppCmd.Flags().String("client-token", "", "A unique token that guarantees that the call to this API is idempotent.")
	sagemaker_createPartnerAppCmd.Flags().Bool("enable-auto-minor-version-upgrade", false, "When set to `TRUE`, the SageMaker Partner AI App is automatically upgraded to the latest minor version during the next scheduled maintenance window, if one is available.")
	sagemaker_createPartnerAppCmd.Flags().Bool("enable-iam-session-based-identity", false, "When set to `TRUE`, the SageMaker Partner AI App sets the Amazon Web Services IAM session name or the authenticated IAM user as the identity of the SageMaker Partner AI App user.")
	sagemaker_createPartnerAppCmd.Flags().String("execution-role-arn", "", "The ARN of the IAM role that the partner application uses.")
	sagemaker_createPartnerAppCmd.Flags().String("kms-key-id", "", "SageMaker Partner AI Apps uses Amazon Web Services KMS to encrypt data at rest using an Amazon Web Services managed key by default.")
	sagemaker_createPartnerAppCmd.Flags().String("maintenance-config", "", "Maintenance configuration settings for the SageMaker Partner AI App.")
	sagemaker_createPartnerAppCmd.Flags().String("name", "", "The name to give the SageMaker Partner AI App.")
	sagemaker_createPartnerAppCmd.Flags().Bool("no-enable-auto-minor-version-upgrade", false, "When set to `TRUE`, the SageMaker Partner AI App is automatically upgraded to the latest minor version during the next scheduled maintenance window, if one is available.")
	sagemaker_createPartnerAppCmd.Flags().Bool("no-enable-iam-session-based-identity", false, "When set to `TRUE`, the SageMaker Partner AI App sets the Amazon Web Services IAM session name or the authenticated IAM user as the identity of the SageMaker Partner AI App user.")
	sagemaker_createPartnerAppCmd.Flags().String("tags", "", "Each tag consists of a key and an optional value.")
	sagemaker_createPartnerAppCmd.Flags().String("tier", "", "Indicates the instance type and size of the cluster attached to the SageMaker Partner AI App.")
	sagemaker_createPartnerAppCmd.Flags().String("type", "", "The type of SageMaker Partner AI App to create.")
	sagemaker_createPartnerAppCmd.MarkFlagRequired("auth-type")
	sagemaker_createPartnerAppCmd.MarkFlagRequired("execution-role-arn")
	sagemaker_createPartnerAppCmd.MarkFlagRequired("name")
	sagemaker_createPartnerAppCmd.Flag("no-enable-auto-minor-version-upgrade").Hidden = true
	sagemaker_createPartnerAppCmd.Flag("no-enable-iam-session-based-identity").Hidden = true
	sagemaker_createPartnerAppCmd.MarkFlagRequired("tier")
	sagemaker_createPartnerAppCmd.MarkFlagRequired("type")
	sagemakerCmd.AddCommand(sagemaker_createPartnerAppCmd)
}
