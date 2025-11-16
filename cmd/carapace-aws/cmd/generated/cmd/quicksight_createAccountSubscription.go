package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createAccountSubscriptionCmd = &cobra.Command{
	Use:   "create-account-subscription",
	Short: "Creates an Amazon Quick Sight account, or subscribes to Amazon Quick Sight Q.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createAccountSubscriptionCmd).Standalone()

	quicksight_createAccountSubscriptionCmd.Flags().String("account-name", "", "The name of your Amazon Quick Sight account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("active-directory-name", "", "The name of your Active Directory.")
	quicksight_createAccountSubscriptionCmd.Flags().String("admin-group", "", "The admin group associated with your Active Directory or IAM Identity Center account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("admin-pro-group", "", "The admin pro group associated with your Active Directory or IAM Identity Center account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("authentication-method", "", "The method that you want to use to authenticate your Quick Sight account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("author-group", "", "The author group associated with your Active Directory or IAM Identity Center account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("author-pro-group", "", "The author pro group associated with your Active Directory or IAM Identity Center account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID of the account that you're using to create your Quick Sight account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("contact-number", "", "A 10-digit phone number for the author of the Amazon Quick Sight account to use for future communications.")
	quicksight_createAccountSubscriptionCmd.Flags().String("directory-id", "", "The ID of the Active Directory that is associated with your Quick Sight account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("edition", "", "The edition of Amazon Quick Sight that you want your account to have.")
	quicksight_createAccountSubscriptionCmd.Flags().String("email-address", "", "The email address of the author of the Amazon Quick Sight account to use for future communications.")
	quicksight_createAccountSubscriptionCmd.Flags().String("first-name", "", "The first name of the author of the Amazon Quick Sight account to use for future communications.")
	quicksight_createAccountSubscriptionCmd.Flags().String("iamidentity-center-instance-arn", "", "The Amazon Resource Name (ARN) for the IAM Identity Center instance.")
	quicksight_createAccountSubscriptionCmd.Flags().String("last-name", "", "The last name of the author of the Amazon Quick Sight account to use for future communications.")
	quicksight_createAccountSubscriptionCmd.Flags().String("notification-email", "", "The email address that you want Quick Sight to send notifications to regarding your Quick Sight account or Quick Sight subscription.")
	quicksight_createAccountSubscriptionCmd.Flags().String("reader-group", "", "The reader group associated with your Active Directory or IAM Identity Center account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("reader-pro-group", "", "The reader pro group associated with your Active Directory or IAM Identity Center account.")
	quicksight_createAccountSubscriptionCmd.Flags().String("realm", "", "The realm of the Active Directory that is associated with your Quick Sight account.")
	quicksight_createAccountSubscriptionCmd.MarkFlagRequired("account-name")
	quicksight_createAccountSubscriptionCmd.MarkFlagRequired("authentication-method")
	quicksight_createAccountSubscriptionCmd.MarkFlagRequired("aws-account-id")
	quicksight_createAccountSubscriptionCmd.MarkFlagRequired("notification-email")
	quicksightCmd.AddCommand(quicksight_createAccountSubscriptionCmd)
}
