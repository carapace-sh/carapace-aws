package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_createPartnershipCmd = &cobra.Command{
	Use:   "create-partnership",
	Short: "Creates a partnership between a customer and a trading partner, based on the supplied parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_createPartnershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_createPartnershipCmd).Standalone()

		b2bi_createPartnershipCmd.Flags().String("capabilities", "", "Specifies a list of the capabilities associated with this partnership.")
		b2bi_createPartnershipCmd.Flags().String("capability-options", "", "Specify the structure that contains the details for the associated capabilities.")
		b2bi_createPartnershipCmd.Flags().String("client-token", "", "Reserved for future use.")
		b2bi_createPartnershipCmd.Flags().String("email", "", "Specifies the email address associated with this trading partner.")
		b2bi_createPartnershipCmd.Flags().String("name", "", "Specifies a descriptive name for the partnership.")
		b2bi_createPartnershipCmd.Flags().String("phone", "", "Specifies the phone number associated with the partnership.")
		b2bi_createPartnershipCmd.Flags().String("profile-id", "", "Specifies the unique, system-generated identifier for the profile connected to this partnership.")
		b2bi_createPartnershipCmd.Flags().String("tags", "", "Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type.")
		b2bi_createPartnershipCmd.MarkFlagRequired("capabilities")
		b2bi_createPartnershipCmd.MarkFlagRequired("email")
		b2bi_createPartnershipCmd.MarkFlagRequired("name")
		b2bi_createPartnershipCmd.MarkFlagRequired("profile-id")
	})
	b2biCmd.AddCommand(b2bi_createPartnershipCmd)
}
