package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getContactMethodsCmd = &cobra.Command{
	Use:   "get-contact-methods",
	Short: "Returns information about the configured contact methods.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getContactMethodsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getContactMethodsCmd).Standalone()

		lightsail_getContactMethodsCmd.Flags().String("protocols", "", "The protocols used to send notifications, such as `Email`, or `SMS` (text messaging).")
	})
	lightsailCmd.AddCommand(lightsail_getContactMethodsCmd)
}
