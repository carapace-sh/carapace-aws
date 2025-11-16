package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_updateDomainEntryCmd = &cobra.Command{
	Use:   "update-domain-entry",
	Short: "Updates a domain recordset after it is created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_updateDomainEntryCmd).Standalone()

	lightsail_updateDomainEntryCmd.Flags().String("domain-entry", "", "An array of key-value pairs containing information about the domain entry.")
	lightsail_updateDomainEntryCmd.Flags().String("domain-name", "", "The name of the domain recordset to update.")
	lightsail_updateDomainEntryCmd.MarkFlagRequired("domain-entry")
	lightsail_updateDomainEntryCmd.MarkFlagRequired("domain-name")
	lightsailCmd.AddCommand(lightsail_updateDomainEntryCmd)
}
