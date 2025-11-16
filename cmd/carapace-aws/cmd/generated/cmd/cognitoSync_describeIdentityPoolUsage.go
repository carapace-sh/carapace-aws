package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_describeIdentityPoolUsageCmd = &cobra.Command{
	Use:   "describe-identity-pool-usage",
	Short: "Gets usage details (for example, data storage) about a particular identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_describeIdentityPoolUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoSync_describeIdentityPoolUsageCmd).Standalone()

		cognitoSync_describeIdentityPoolUsageCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
		cognitoSync_describeIdentityPoolUsageCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoSyncCmd.AddCommand(cognitoSync_describeIdentityPoolUsageCmd)
}
