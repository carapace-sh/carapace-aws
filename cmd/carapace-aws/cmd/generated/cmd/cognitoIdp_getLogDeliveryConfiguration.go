package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getLogDeliveryConfigurationCmd = &cobra.Command{
	Use:   "get-log-delivery-configuration",
	Short: "Given a user pool ID, returns the logging configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getLogDeliveryConfigurationCmd).Standalone()

	cognitoIdp_getLogDeliveryConfigurationCmd.Flags().String("user-pool-id", "", "The ID of the user pool that has the logging configuration that you want to view.")
	cognitoIdp_getLogDeliveryConfigurationCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_getLogDeliveryConfigurationCmd)
}
