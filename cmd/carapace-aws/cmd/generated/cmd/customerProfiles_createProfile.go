package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createProfileCmd = &cobra.Command{
	Use:   "create-profile",
	Short: "Creates a standard profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createProfileCmd).Standalone()

	customerProfiles_createProfileCmd.Flags().String("account-number", "", "An account number that you have assigned to the customer.")
	customerProfiles_createProfileCmd.Flags().String("additional-information", "", "Any additional information relevant to the customer’s profile.")
	customerProfiles_createProfileCmd.Flags().String("address", "", "A generic address associated with the customer that is not mailing, shipping, or billing.")
	customerProfiles_createProfileCmd.Flags().String("attributes", "", "A key value pair of attributes of a customer profile.")
	customerProfiles_createProfileCmd.Flags().String("billing-address", "", "The customer’s billing address.")
	customerProfiles_createProfileCmd.Flags().String("birth-date", "", "The customer’s birth date.")
	customerProfiles_createProfileCmd.Flags().String("business-email-address", "", "The customer’s business email address.")
	customerProfiles_createProfileCmd.Flags().String("business-name", "", "The name of the customer’s business.")
	customerProfiles_createProfileCmd.Flags().String("business-phone-number", "", "The customer’s business phone number.")
	customerProfiles_createProfileCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_createProfileCmd.Flags().String("email-address", "", "The customer’s email address, which has not been specified as a personal or business address.")
	customerProfiles_createProfileCmd.Flags().String("engagement-preferences", "", "Object that defines the preferred methods of engagement, per channel.")
	customerProfiles_createProfileCmd.Flags().String("first-name", "", "The customer’s first name.")
	customerProfiles_createProfileCmd.Flags().String("gender", "", "The gender with which the customer identifies.")
	customerProfiles_createProfileCmd.Flags().String("gender-string", "", "An alternative to `Gender` which accepts any string as input.")
	customerProfiles_createProfileCmd.Flags().String("home-phone-number", "", "The customer’s home phone number.")
	customerProfiles_createProfileCmd.Flags().String("last-name", "", "The customer’s last name.")
	customerProfiles_createProfileCmd.Flags().String("mailing-address", "", "The customer’s mailing address.")
	customerProfiles_createProfileCmd.Flags().String("middle-name", "", "The customer’s middle name.")
	customerProfiles_createProfileCmd.Flags().String("mobile-phone-number", "", "The customer’s mobile phone number.")
	customerProfiles_createProfileCmd.Flags().String("party-type", "", "The type of profile used to describe the customer.")
	customerProfiles_createProfileCmd.Flags().String("party-type-string", "", "An alternative to `PartyType` which accepts any string as input.")
	customerProfiles_createProfileCmd.Flags().String("personal-email-address", "", "The customer’s personal email address.")
	customerProfiles_createProfileCmd.Flags().String("phone-number", "", "The customer’s phone number, which has not been specified as a mobile, home, or business number.")
	customerProfiles_createProfileCmd.Flags().String("profile-type", "", "The type of the profile.")
	customerProfiles_createProfileCmd.Flags().String("shipping-address", "", "The customer’s shipping address.")
	customerProfiles_createProfileCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_createProfileCmd)
}
