package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_setLogDeliveryConfigurationCmd = &cobra.Command{
	Use:   "set-log-delivery-configuration",
	Short: "Sets up or modifies the logging configuration of a user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_setLogDeliveryConfigurationCmd).Standalone()

	cognitoIdp_setLogDeliveryConfigurationCmd.Flags().String("log-configurations", "", "A collection of the logging configurations for a user pool.")
	cognitoIdp_setLogDeliveryConfigurationCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to configure logging.")
	cognitoIdp_setLogDeliveryConfigurationCmd.MarkFlagRequired("log-configurations")
	cognitoIdp_setLogDeliveryConfigurationCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_setLogDeliveryConfigurationCmd)
}
