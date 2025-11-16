package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeRiskConfigurationCmd = &cobra.Command{
	Use:   "describe-risk-configuration",
	Short: "Given an app client or user pool ID where threat protection is configured, describes the risk configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeRiskConfigurationCmd).Standalone()

	cognitoIdp_describeRiskConfigurationCmd.Flags().String("client-id", "", "The ID of the app client with the risk configuration that you want to inspect.")
	cognitoIdp_describeRiskConfigurationCmd.Flags().String("user-pool-id", "", "The ID of the user pool with the risk configuration that you want to inspect.")
	cognitoIdp_describeRiskConfigurationCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_describeRiskConfigurationCmd)
}
