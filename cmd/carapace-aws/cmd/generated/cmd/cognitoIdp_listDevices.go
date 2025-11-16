package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "Lists the devices that Amazon Cognito has registered to the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listDevicesCmd).Standalone()

	cognitoIdp_listDevicesCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_listDevicesCmd.Flags().String("limit", "", "The maximum number of devices that you want Amazon Cognito to return in the response.")
	cognitoIdp_listDevicesCmd.Flags().String("pagination-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listDevicesCmd.MarkFlagRequired("access-token")
	cognitoIdpCmd.AddCommand(cognitoIdp_listDevicesCmd)
}
