package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Updates an Amazon Quick Sight user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateUserCmd).Standalone()

	quicksight_updateUserCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the user is in.")
	quicksight_updateUserCmd.Flags().String("custom-federation-provider-url", "", "The URL of the custom OpenID Connect (OIDC) provider that provides identity to let a user federate into Quick Sight with an associated Identity and Access Management(IAM) role.")
	quicksight_updateUserCmd.Flags().String("custom-permissions-name", "", "(Enterprise edition only) The name of the custom permissions profile that you want to assign to this user.")
	quicksight_updateUserCmd.Flags().String("email", "", "The email address of the user that you want to update.")
	quicksight_updateUserCmd.Flags().String("external-login-federation-provider-type", "", "The type of supported external login provider that provides identity to let a user federate into Quick Sight with an associated Identity and Access Management(IAM) role.")
	quicksight_updateUserCmd.Flags().String("external-login-id", "", "The identity ID for a user in the external login provider.")
	quicksight_updateUserCmd.Flags().String("namespace", "", "The namespace.")
	quicksight_updateUserCmd.Flags().Bool("no-unapply-custom-permissions", false, "A flag that you use to indicate that you want to remove all custom permissions from this user.")
	quicksight_updateUserCmd.Flags().String("role", "", "The Amazon Quick Sight role of the user.")
	quicksight_updateUserCmd.Flags().Bool("unapply-custom-permissions", false, "A flag that you use to indicate that you want to remove all custom permissions from this user.")
	quicksight_updateUserCmd.Flags().String("user-name", "", "The Amazon Quick Sight user name that you want to update.")
	quicksight_updateUserCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateUserCmd.MarkFlagRequired("email")
	quicksight_updateUserCmd.MarkFlagRequired("namespace")
	quicksight_updateUserCmd.Flag("no-unapply-custom-permissions").Hidden = true
	quicksight_updateUserCmd.MarkFlagRequired("role")
	quicksight_updateUserCmd.MarkFlagRequired("user-name")
	quicksightCmd.AddCommand(quicksight_updateUserCmd)
}
