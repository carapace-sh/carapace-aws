package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_getIdentityPoolConfigurationCmd = &cobra.Command{
	Use:   "get-identity-pool-configuration",
	Short: "Gets the configuration settings of an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_getIdentityPoolConfigurationCmd).Standalone()

	cognitoSync_getIdentityPoolConfigurationCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_getIdentityPoolConfigurationCmd.MarkFlagRequired("identity-pool-id")
	cognitoSyncCmd.AddCommand(cognitoSync_getIdentityPoolConfigurationCmd)
}
