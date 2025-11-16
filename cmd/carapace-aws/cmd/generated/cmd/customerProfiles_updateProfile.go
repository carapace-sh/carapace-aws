package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_updateProfileCmd = &cobra.Command{
	Use:   "update-profile",
	Short: "Updates the properties of a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_updateProfileCmd).Standalone()

	customerProfiles_updateProfileCmd.Flags().String("account-number", "", "An account number that you have assigned to the customer.")
	customerProfiles_updateProfileCmd.Flags().String("additional-information", "", "Any additional information relevant to the customer’s profile.")
	customerProfiles_updateProfileCmd.Flags().String("address", "", "A generic address associated with the customer that is not mailing, shipping, or billing.")
	customerProfiles_updateProfileCmd.Flags().String("attributes", "", "A key value pair of attributes of a customer profile.")
	customerProfiles_updateProfileCmd.Flags().String("billing-address", "", "The customer’s billing address.")
	customerProfiles_updateProfileCmd.Flags().String("birth-date", "", "The customer’s birth date.")
	customerProfiles_updateProfileCmd.Flags().String("business-email-address", "", "The customer’s business email address.")
	customerProfiles_updateProfileCmd.Flags().String("business-name", "", "The name of the customer’s business.")
	customerProfiles_updateProfileCmd.Flags().String("business-phone-number", "", "The customer’s business phone number.")
	customerProfiles_updateProfileCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_updateProfileCmd.Flags().String("email-address", "", "The customer’s email address, which has not been specified as a personal or business address.")
	customerProfiles_updateProfileCmd.Flags().String("engagement-preferences", "", "Object that defines users preferred methods of engagement.")
	customerProfiles_updateProfileCmd.Flags().String("first-name", "", "The customer’s first name.")
	customerProfiles_updateProfileCmd.Flags().String("gender", "", "The gender with which the customer identifies.")
	customerProfiles_updateProfileCmd.Flags().String("gender-string", "", "An alternative to `Gender` which accepts any string as input.")
	customerProfiles_updateProfileCmd.Flags().String("home-phone-number", "", "The customer’s home phone number.")
	customerProfiles_updateProfileCmd.Flags().String("last-name", "", "The customer’s last name.")
	customerProfiles_updateProfileCmd.Flags().String("mailing-address", "", "The customer’s mailing address.")
	customerProfiles_updateProfileCmd.Flags().String("middle-name", "", "The customer’s middle name.")
	customerProfiles_updateProfileCmd.Flags().String("mobile-phone-number", "", "The customer’s mobile phone number.")
	customerProfiles_updateProfileCmd.Flags().String("party-type", "", "The type of profile used to describe the customer.")
	customerProfiles_updateProfileCmd.Flags().String("party-type-string", "", "An alternative to `PartyType` which accepts any string as input.")
	customerProfiles_updateProfileCmd.Flags().String("personal-email-address", "", "The customer’s personal email address.")
	customerProfiles_updateProfileCmd.Flags().String("phone-number", "", "The customer’s phone number, which has not been specified as a mobile, home, or business number.")
	customerProfiles_updateProfileCmd.Flags().String("profile-id", "", "The unique identifier of a customer profile.")
	customerProfiles_updateProfileCmd.Flags().String("profile-type", "", "Determines the type of the profile.")
	customerProfiles_updateProfileCmd.Flags().String("shipping-address", "", "The customer’s shipping address.")
	customerProfiles_updateProfileCmd.MarkFlagRequired("domain-name")
	customerProfiles_updateProfileCmd.MarkFlagRequired("profile-id")
	customerProfilesCmd.AddCommand(customerProfiles_updateProfileCmd)
}
