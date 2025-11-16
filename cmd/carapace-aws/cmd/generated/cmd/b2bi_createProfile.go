package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_createProfileCmd = &cobra.Command{
	Use:   "create-profile",
	Short: "Creates a customer profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_createProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_createProfileCmd).Standalone()

		b2bi_createProfileCmd.Flags().String("business-name", "", "Specifies the name for the business associated with this profile.")
		b2bi_createProfileCmd.Flags().String("client-token", "", "Reserved for future use.")
		b2bi_createProfileCmd.Flags().String("email", "", "Specifies the email address associated with this customer profile.")
		b2bi_createProfileCmd.Flags().String("logging", "", "Specifies whether or not logging is enabled for this profile.")
		b2bi_createProfileCmd.Flags().String("name", "", "Specifies the name of the profile.")
		b2bi_createProfileCmd.Flags().String("phone", "", "Specifies the phone number associated with the profile.")
		b2bi_createProfileCmd.Flags().String("tags", "", "Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type.")
		b2bi_createProfileCmd.MarkFlagRequired("business-name")
		b2bi_createProfileCmd.MarkFlagRequired("logging")
		b2bi_createProfileCmd.MarkFlagRequired("name")
		b2bi_createProfileCmd.MarkFlagRequired("phone")
	})
	b2biCmd.AddCommand(b2bi_createProfileCmd)
}
