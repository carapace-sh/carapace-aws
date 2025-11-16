package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_updateProfileCmd = &cobra.Command{
	Use:   "update-profile",
	Short: "Updates the specified parameters for a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_updateProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_updateProfileCmd).Standalone()

		b2bi_updateProfileCmd.Flags().String("business-name", "", "Specifies the name for the business associated with this profile.")
		b2bi_updateProfileCmd.Flags().String("email", "", "Specifies the email address associated with this customer profile.")
		b2bi_updateProfileCmd.Flags().String("name", "", "The name of the profile, used to identify it.")
		b2bi_updateProfileCmd.Flags().String("phone", "", "Specifies the phone number associated with the profile.")
		b2bi_updateProfileCmd.Flags().String("profile-id", "", "Specifies the unique, system-generated identifier for the profile.")
		b2bi_updateProfileCmd.MarkFlagRequired("profile-id")
	})
	b2biCmd.AddCommand(b2bi_updateProfileCmd)
}
