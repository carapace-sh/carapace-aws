package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createDomainEntryCmd = &cobra.Command{
	Use:   "create-domain-entry",
	Short: "Creates one of the following domain name system (DNS) records in a domain DNS zone: Address (A), canonical name (CNAME), mail exchanger (MX), name server (NS), start of authority (SOA), service locator (SRV), or text (TXT).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createDomainEntryCmd).Standalone()

	lightsail_createDomainEntryCmd.Flags().String("domain-entry", "", "An array of key-value pairs containing information about the domain entry request.")
	lightsail_createDomainEntryCmd.Flags().String("domain-name", "", "The domain name (`example.com`) for which you want to create the domain entry.")
	lightsail_createDomainEntryCmd.MarkFlagRequired("domain-entry")
	lightsail_createDomainEntryCmd.MarkFlagRequired("domain-name")
	lightsailCmd.AddCommand(lightsail_createDomainEntryCmd)
}
