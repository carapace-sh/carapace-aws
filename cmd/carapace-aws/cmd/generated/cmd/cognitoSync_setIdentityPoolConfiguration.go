package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_setIdentityPoolConfigurationCmd = &cobra.Command{
	Use:   "set-identity-pool-configuration",
	Short: "Sets the necessary configuration for push sync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_setIdentityPoolConfigurationCmd).Standalone()

	cognitoSync_setIdentityPoolConfigurationCmd.Flags().String("cognito-streams", "", "Options to apply to this identity pool for Amazon Cognito streams.")
	cognitoSync_setIdentityPoolConfigurationCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_setIdentityPoolConfigurationCmd.Flags().String("push-sync", "", "Options to apply to this identity pool for push synchronization.")
	cognitoSync_setIdentityPoolConfigurationCmd.MarkFlagRequired("identity-pool-id")
	cognitoSyncCmd.AddCommand(cognitoSync_setIdentityPoolConfigurationCmd)
}
