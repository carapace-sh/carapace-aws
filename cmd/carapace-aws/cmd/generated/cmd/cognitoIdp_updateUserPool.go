package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateUserPoolCmd = &cobra.Command{
	Use:   "update-user-pool",
	Short: "Updates the configuration of a user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateUserPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_updateUserPoolCmd).Standalone()

		cognitoIdp_updateUserPoolCmd.Flags().String("account-recovery-setting", "", "The available verified method a user can use to recover their password when they call `ForgotPassword`.")
		cognitoIdp_updateUserPoolCmd.Flags().String("admin-create-user-config", "", "The configuration for administrative creation of users.")
		cognitoIdp_updateUserPoolCmd.Flags().String("auto-verified-attributes", "", "The attributes that you want your user pool to automatically verify.")
		cognitoIdp_updateUserPoolCmd.Flags().String("deletion-protection", "", "When active, `DeletionProtection` prevents accidental deletion of your user pool.")
		cognitoIdp_updateUserPoolCmd.Flags().String("device-configuration", "", "The device-remembering configuration for a user pool.")
		cognitoIdp_updateUserPoolCmd.Flags().String("email-configuration", "", "The email configuration of your user pool.")
		cognitoIdp_updateUserPoolCmd.Flags().String("email-verification-message", "", "This parameter is no longer used.")
		cognitoIdp_updateUserPoolCmd.Flags().String("email-verification-subject", "", "This parameter is no longer used.")
		cognitoIdp_updateUserPoolCmd.Flags().String("lambda-config", "", "A collection of user pool Lambda triggers.")
		cognitoIdp_updateUserPoolCmd.Flags().String("mfa-configuration", "", "Sets multi-factor authentication (MFA) to be on, off, or optional.")
		cognitoIdp_updateUserPoolCmd.Flags().String("policies", "", "The password policy and sign-in policy in the user pool.")
		cognitoIdp_updateUserPoolCmd.Flags().String("pool-name", "", "The updated name of your user pool.")
		cognitoIdp_updateUserPoolCmd.Flags().String("sms-authentication-message", "", "The contents of the SMS message that your user pool sends to users in SMS authentication.")
		cognitoIdp_updateUserPoolCmd.Flags().String("sms-configuration", "", "The SMS configuration with the settings for your Amazon Cognito user pool to send SMS message with Amazon Simple Notification Service.")
		cognitoIdp_updateUserPoolCmd.Flags().String("sms-verification-message", "", "This parameter is no longer used.")
		cognitoIdp_updateUserPoolCmd.Flags().String("user-attribute-update-settings", "", "The settings for updates to user attributes.")
		cognitoIdp_updateUserPoolCmd.Flags().String("user-pool-add-ons", "", "Contains settings for activation of threat protection, including the operating mode and additional authentication types.")
		cognitoIdp_updateUserPoolCmd.Flags().String("user-pool-id", "", "The ID of the user pool you want to update.")
		cognitoIdp_updateUserPoolCmd.Flags().String("user-pool-tags", "", "The tag keys and values to assign to the user pool.")
		cognitoIdp_updateUserPoolCmd.Flags().String("user-pool-tier", "", "The user pool [feature plan](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-sign-in-feature-plans.html), or tier.")
		cognitoIdp_updateUserPoolCmd.Flags().String("verification-message-template", "", "The template for the verification message that your user pool delivers to users who set an email address or phone number attribute.")
		cognitoIdp_updateUserPoolCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_updateUserPoolCmd)
}
