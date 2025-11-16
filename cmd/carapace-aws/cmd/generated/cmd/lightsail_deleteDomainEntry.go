package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteDomainEntryCmd = &cobra.Command{
	Use:   "delete-domain-entry",
	Short: "Deletes a specific domain entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteDomainEntryCmd).Standalone()

	lightsail_deleteDomainEntryCmd.Flags().String("domain-entry", "", "An array of key-value pairs containing information about your domain entries.")
	lightsail_deleteDomainEntryCmd.Flags().String("domain-name", "", "The name of the domain entry to delete.")
	lightsail_deleteDomainEntryCmd.MarkFlagRequired("domain-entry")
	lightsail_deleteDomainEntryCmd.MarkFlagRequired("domain-name")
	lightsailCmd.AddCommand(lightsail_deleteDomainEntryCmd)
}
