package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminListDevicesCmd = &cobra.Command{
	Use:   "admin-list-devices",
	Short: "Lists a user's registered devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminListDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminListDevicesCmd).Standalone()

		cognitoIdp_adminListDevicesCmd.Flags().String("limit", "", "The maximum number of devices that you want Amazon Cognito to return in the response.")
		cognitoIdp_adminListDevicesCmd.Flags().String("pagination-token", "", "This API operation returns a limited number of results.")
		cognitoIdp_adminListDevicesCmd.Flags().String("user-pool-id", "", "The ID of the user pool where the device owner is a user.")
		cognitoIdp_adminListDevicesCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminListDevicesCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminListDevicesCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminListDevicesCmd)
}
