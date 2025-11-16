package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updatePartnerAppCmd = &cobra.Command{
	Use:   "update-partner-app",
	Short: "Updates all of the SageMaker Partner AI Apps in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updatePartnerAppCmd).Standalone()

	sagemaker_updatePartnerAppCmd.Flags().String("app-version", "", "The semantic version to upgrade the SageMaker Partner AI App to.")
	sagemaker_updatePartnerAppCmd.Flags().String("application-config", "", "Configuration settings for the SageMaker Partner AI App.")
	sagemaker_updatePartnerAppCmd.Flags().String("arn", "", "The ARN of the SageMaker Partner AI App to update.")
	sagemaker_updatePartnerAppCmd.Flags().String("client-token", "", "A unique token that guarantees that the call to this API is idempotent.")
	sagemaker_updatePartnerAppCmd.Flags().Bool("enable-auto-minor-version-upgrade", false, "When set to `TRUE`, the SageMaker Partner AI App is automatically upgraded to the latest minor version during the next scheduled maintenance window, if one is available.")
	sagemaker_updatePartnerAppCmd.Flags().Bool("enable-iam-session-based-identity", false, "When set to `TRUE`, the SageMaker Partner AI App sets the Amazon Web Services IAM session name or the authenticated IAM user as the identity of the SageMaker Partner AI App user.")
	sagemaker_updatePartnerAppCmd.Flags().String("maintenance-config", "", "Maintenance configuration settings for the SageMaker Partner AI App.")
	sagemaker_updatePartnerAppCmd.Flags().Bool("no-enable-auto-minor-version-upgrade", false, "When set to `TRUE`, the SageMaker Partner AI App is automatically upgraded to the latest minor version during the next scheduled maintenance window, if one is available.")
	sagemaker_updatePartnerAppCmd.Flags().Bool("no-enable-iam-session-based-identity", false, "When set to `TRUE`, the SageMaker Partner AI App sets the Amazon Web Services IAM session name or the authenticated IAM user as the identity of the SageMaker Partner AI App user.")
	sagemaker_updatePartnerAppCmd.Flags().String("tags", "", "Each tag consists of a key and an optional value.")
	sagemaker_updatePartnerAppCmd.Flags().String("tier", "", "Indicates the instance type and size of the cluster attached to the SageMaker Partner AI App.")
	sagemaker_updatePartnerAppCmd.MarkFlagRequired("arn")
	sagemaker_updatePartnerAppCmd.Flag("no-enable-auto-minor-version-upgrade").Hidden = true
	sagemaker_updatePartnerAppCmd.Flag("no-enable-iam-session-based-identity").Hidden = true
	sagemakerCmd.AddCommand(sagemaker_updatePartnerAppCmd)
}
