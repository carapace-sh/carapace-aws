package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_createUserPoolCmd = &cobra.Command{
	Use:   "create-user-pool",
	Short: "Creates a new Amazon Cognito user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_createUserPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_createUserPoolCmd).Standalone()

		cognitoIdp_createUserPoolCmd.Flags().String("account-recovery-setting", "", "The available verified method a user can use to recover their password when they call `ForgotPassword`.")
		cognitoIdp_createUserPoolCmd.Flags().String("admin-create-user-config", "", "The configuration for administrative creation of users.")
		cognitoIdp_createUserPoolCmd.Flags().String("alias-attributes", "", "Attributes supported as an alias for this user pool.")
		cognitoIdp_createUserPoolCmd.Flags().String("auto-verified-attributes", "", "The attributes that you want your user pool to automatically verify.")
		cognitoIdp_createUserPoolCmd.Flags().String("deletion-protection", "", "When active, `DeletionProtection` prevents accidental deletion of your user pool.")
		cognitoIdp_createUserPoolCmd.Flags().String("device-configuration", "", "The device-remembering configuration for a user pool.")
		cognitoIdp_createUserPoolCmd.Flags().String("email-configuration", "", "The email configuration of your user pool.")
		cognitoIdp_createUserPoolCmd.Flags().String("email-verification-message", "", "This parameter is no longer used.")
		cognitoIdp_createUserPoolCmd.Flags().String("email-verification-subject", "", "This parameter is no longer used.")
		cognitoIdp_createUserPoolCmd.Flags().String("lambda-config", "", "A collection of user pool Lambda triggers.")
		cognitoIdp_createUserPoolCmd.Flags().String("mfa-configuration", "", "Sets multi-factor authentication (MFA) to be on, off, or optional.")
		cognitoIdp_createUserPoolCmd.Flags().String("policies", "", "The password policy and sign-in policy in the user pool.")
		cognitoIdp_createUserPoolCmd.Flags().String("pool-name", "", "A friendly name for your user pool.")
		cognitoIdp_createUserPoolCmd.Flags().String("schema", "", "An array of attributes for the new user pool.")
		cognitoIdp_createUserPoolCmd.Flags().String("sms-authentication-message", "", "The contents of the SMS message that your user pool sends to users in SMS OTP and MFA authentication.")
		cognitoIdp_createUserPoolCmd.Flags().String("sms-configuration", "", "The settings for your Amazon Cognito user pool to send SMS messages with Amazon Simple Notification Service.")
		cognitoIdp_createUserPoolCmd.Flags().String("sms-verification-message", "", "This parameter is no longer used.")
		cognitoIdp_createUserPoolCmd.Flags().String("user-attribute-update-settings", "", "The settings for updates to user attributes.")
		cognitoIdp_createUserPoolCmd.Flags().String("user-pool-add-ons", "", "Contains settings for activation of threat protection, including the operating mode and additional authentication types.")
		cognitoIdp_createUserPoolCmd.Flags().String("user-pool-tags", "", "The tag keys and values to assign to the user pool.")
		cognitoIdp_createUserPoolCmd.Flags().String("user-pool-tier", "", "The user pool [feature plan](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-sign-in-feature-plans.html), or tier.")
		cognitoIdp_createUserPoolCmd.Flags().String("username-attributes", "", "Specifies whether a user can use an email address or phone number as a username when they sign up.")
		cognitoIdp_createUserPoolCmd.Flags().String("username-configuration", "", "Sets the case sensitivity option for sign-in usernames.")
		cognitoIdp_createUserPoolCmd.Flags().String("verification-message-template", "", "The template for the verification message that your user pool delivers to users who set an email address or phone number attribute.")
		cognitoIdp_createUserPoolCmd.MarkFlagRequired("pool-name")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_createUserPoolCmd)
}
