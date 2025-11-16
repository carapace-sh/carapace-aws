package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createContactMethodCmd = &cobra.Command{
	Use:   "create-contact-method",
	Short: "Creates an email or SMS text message contact method.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createContactMethodCmd).Standalone()

	lightsail_createContactMethodCmd.Flags().String("contact-endpoint", "", "The destination of the contact method, such as an email address or a mobile phone number.")
	lightsail_createContactMethodCmd.Flags().String("protocol", "", "The protocol of the contact method, such as `Email` or `SMS` (text messaging).")
	lightsail_createContactMethodCmd.MarkFlagRequired("contact-endpoint")
	lightsail_createContactMethodCmd.MarkFlagRequired("protocol")
	lightsailCmd.AddCommand(lightsail_createContactMethodCmd)
}
