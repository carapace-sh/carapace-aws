package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_setRiskConfigurationCmd = &cobra.Command{
	Use:   "set-risk-configuration",
	Short: "Configures threat protection for a user pool or app client.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_setRiskConfigurationCmd).Standalone()

	cognitoIdp_setRiskConfigurationCmd.Flags().String("account-takeover-risk-configuration", "", "The settings for automated responses and notification templates for adaptive authentication with threat protection.")
	cognitoIdp_setRiskConfigurationCmd.Flags().String("client-id", "", "The ID of the app client where you want to set a risk configuration.")
	cognitoIdp_setRiskConfigurationCmd.Flags().String("compromised-credentials-risk-configuration", "", "The configuration of automated reactions to detected compromised credentials.")
	cognitoIdp_setRiskConfigurationCmd.Flags().String("risk-exception-configuration", "", "A set of IP-address overrides to threat protection.")
	cognitoIdp_setRiskConfigurationCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to set a risk configuration.")
	cognitoIdp_setRiskConfigurationCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_setRiskConfigurationCmd)
}
