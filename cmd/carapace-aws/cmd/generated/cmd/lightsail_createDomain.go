package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates a domain resource for the specified domain (example.com).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createDomainCmd).Standalone()

		lightsail_createDomainCmd.Flags().String("domain-name", "", "The domain name to manage (`example.com`).")
		lightsail_createDomainCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
		lightsail_createDomainCmd.MarkFlagRequired("domain-name")
	})
	lightsailCmd.AddCommand(lightsail_createDomainCmd)
}
