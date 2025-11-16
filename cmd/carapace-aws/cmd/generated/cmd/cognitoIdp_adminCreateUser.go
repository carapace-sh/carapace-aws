package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminCreateUserCmd = &cobra.Command{
	Use:   "admin-create-user",
	Short: "Creates a new user in the specified user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminCreateUserCmd).Standalone()

	cognitoIdp_adminCreateUserCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
	cognitoIdp_adminCreateUserCmd.Flags().String("desired-delivery-mediums", "", "Specify `EMAIL` if email will be used to send the welcome message.")
	cognitoIdp_adminCreateUserCmd.Flags().String("force-alias-creation", "", "This parameter is used only if the `phone_number_verified` or `email_verified` attribute is set to `True`.")
	cognitoIdp_adminCreateUserCmd.Flags().String("message-action", "", "Set to `RESEND` to resend the invitation message to a user that already exists, and to reset the temporary-password duration with a new temporary password.")
	cognitoIdp_adminCreateUserCmd.Flags().String("temporary-password", "", "The user's temporary password.")
	cognitoIdp_adminCreateUserCmd.Flags().String("user-attributes", "", "An array of name-value pairs that contain user attributes and attribute values to be set for the user to be created.")
	cognitoIdp_adminCreateUserCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to create a user.")
	cognitoIdp_adminCreateUserCmd.Flags().String("username", "", "The value that you want to set as the username sign-in attribute.")
	cognitoIdp_adminCreateUserCmd.Flags().String("validation-data", "", "Temporary user attributes that contribute to the outcomes of your pre sign-up Lambda trigger.")
	cognitoIdp_adminCreateUserCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminCreateUserCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminCreateUserCmd)
}
