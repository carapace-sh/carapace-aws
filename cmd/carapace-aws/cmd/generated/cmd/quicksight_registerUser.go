package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_registerUserCmd = &cobra.Command{
	Use:   "register-user",
	Short: "Creates an Amazon Quick Sight user whose identity is associated with the Identity and Access Management (IAM) identity or role specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_registerUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_registerUserCmd).Standalone()

		quicksight_registerUserCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the user is in.")
		quicksight_registerUserCmd.Flags().String("custom-federation-provider-url", "", "The URL of the custom OpenID Connect (OIDC) provider that provides identity to let a user federate into Quick Sight with an associated Identity and Access Management(IAM) role.")
		quicksight_registerUserCmd.Flags().String("custom-permissions-name", "", "(Enterprise edition only) The name of the custom permissions profile that you want to assign to this user.")
		quicksight_registerUserCmd.Flags().String("email", "", "The email address of the user that you want to register.")
		quicksight_registerUserCmd.Flags().String("external-login-federation-provider-type", "", "The type of supported external login provider that provides identity to let a user federate into Amazon Quick Sight with an associated Identity and Access Management(IAM) role.")
		quicksight_registerUserCmd.Flags().String("external-login-id", "", "The identity ID for a user in the external login provider.")
		quicksight_registerUserCmd.Flags().String("iam-arn", "", "The ARN of the IAM user or role that you are registering with Amazon Quick Sight.")
		quicksight_registerUserCmd.Flags().String("identity-type", "", "The identity type that your Quick Sight account uses to manage the identity of users.")
		quicksight_registerUserCmd.Flags().String("namespace", "", "The namespace.")
		quicksight_registerUserCmd.Flags().String("session-name", "", "You need to use this parameter only when you register one or more users using an assumed IAM role.")
		quicksight_registerUserCmd.Flags().String("tags", "", "The tags to associate with the user.")
		quicksight_registerUserCmd.Flags().String("user-name", "", "The Amazon Quick Sight user name that you want to create for the user you are registering.")
		quicksight_registerUserCmd.Flags().String("user-role", "", "The Amazon Quick Sight role for the user.")
		quicksight_registerUserCmd.MarkFlagRequired("aws-account-id")
		quicksight_registerUserCmd.MarkFlagRequired("email")
		quicksight_registerUserCmd.MarkFlagRequired("identity-type")
		quicksight_registerUserCmd.MarkFlagRequired("namespace")
		quicksight_registerUserCmd.MarkFlagRequired("user-role")
	})
	quicksightCmd.AddCommand(quicksight_registerUserCmd)
}
