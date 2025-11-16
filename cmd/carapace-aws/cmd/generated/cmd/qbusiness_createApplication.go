package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createApplicationCmd).Standalone()

	qbusiness_createApplicationCmd.Flags().String("attachments-configuration", "", "An option to allow end users to upload files directly during chat.")
	qbusiness_createApplicationCmd.Flags().String("client-ids-for-oidc", "", "The OIDC client ID for a Amazon Q Business application.")
	qbusiness_createApplicationCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create your Amazon Q Business application.")
	qbusiness_createApplicationCmd.Flags().String("description", "", "A description for the Amazon Q Business application.")
	qbusiness_createApplicationCmd.Flags().String("display-name", "", "A name for the Amazon Q Business application.")
	qbusiness_createApplicationCmd.Flags().String("encryption-configuration", "", "The identifier of the KMS key that is used to encrypt your data.")
	qbusiness_createApplicationCmd.Flags().String("iam-identity-provider-arn", "", "The Amazon Resource Name (ARN) of an identity provider being used by an Amazon Q Business application.")
	qbusiness_createApplicationCmd.Flags().String("identity-center-instance-arn", "", "The Amazon Resource Name (ARN) of the IAM Identity Center instance you are either creating for—or connecting to—your Amazon Q Business application.")
	qbusiness_createApplicationCmd.Flags().String("identity-type", "", "The authentication type being used by a Amazon Q Business application.")
	qbusiness_createApplicationCmd.Flags().String("personalization-configuration", "", "Configuration information about chat response personalization.")
	qbusiness_createApplicationCmd.Flags().String("q-apps-configuration", "", "An option to allow end users to create and use Amazon Q Apps in the web experience.")
	qbusiness_createApplicationCmd.Flags().String("quick-sight-configuration", "", "The Amazon QuickSight configuration for an Amazon Q Business application that uses QuickSight for authentication.")
	qbusiness_createApplicationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permissions to access your Amazon CloudWatch logs and metrics.")
	qbusiness_createApplicationCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize your Amazon Q Business application.")
	qbusiness_createApplicationCmd.MarkFlagRequired("display-name")
	qbusinessCmd.AddCommand(qbusiness_createApplicationCmd)
}
