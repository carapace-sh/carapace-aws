package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_describeIdentityUsageCmd = &cobra.Command{
	Use:   "describe-identity-usage",
	Short: "Gets usage information for an identity, including number of datasets and data usage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_describeIdentityUsageCmd).Standalone()

	cognitoSync_describeIdentityUsageCmd.Flags().String("identity-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_describeIdentityUsageCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_describeIdentityUsageCmd.MarkFlagRequired("identity-id")
	cognitoSync_describeIdentityUsageCmd.MarkFlagRequired("identity-pool-id")
	cognitoSyncCmd.AddCommand(cognitoSync_describeIdentityUsageCmd)
}
