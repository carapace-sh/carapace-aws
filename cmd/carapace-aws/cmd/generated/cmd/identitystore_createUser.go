package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a user within the specified identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_createUserCmd).Standalone()

	identitystore_createUserCmd.Flags().String("addresses", "", "A list of `Address` objects containing addresses associated with the user.")
	identitystore_createUserCmd.Flags().String("birthdate", "", "The user's birthdate in YYYY-MM-DD format.")
	identitystore_createUserCmd.Flags().String("display-name", "", "A string containing the name of the user.")
	identitystore_createUserCmd.Flags().String("emails", "", "A list of `Email` objects containing email addresses associated with the user.")
	identitystore_createUserCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_createUserCmd.Flags().String("locale", "", "A string containing the geographical region or location of the user.")
	identitystore_createUserCmd.Flags().String("name", "", "An object containing the name of the user.")
	identitystore_createUserCmd.Flags().String("nick-name", "", "A string containing an alternate name for the user.")
	identitystore_createUserCmd.Flags().String("phone-numbers", "", "A list of `PhoneNumber` objects containing phone numbers associated with the user.")
	identitystore_createUserCmd.Flags().String("photos", "", "A list of photos associated with the user.")
	identitystore_createUserCmd.Flags().String("preferred-language", "", "A string containing the preferred language of the user.")
	identitystore_createUserCmd.Flags().String("profile-url", "", "A string containing a URL that might be associated with the user.")
	identitystore_createUserCmd.Flags().String("timezone", "", "A string containing the time zone of the user.")
	identitystore_createUserCmd.Flags().String("title", "", "A string containing the title of the user.")
	identitystore_createUserCmd.Flags().String("user-name", "", "A unique string used to identify the user.")
	identitystore_createUserCmd.Flags().String("user-type", "", "A string indicating the type of user.")
	identitystore_createUserCmd.Flags().String("website", "", "The user's personal website or blog URL.")
	identitystore_createUserCmd.MarkFlagRequired("identity-store-id")
	identitystoreCmd.AddCommand(identitystore_createUserCmd)
}
