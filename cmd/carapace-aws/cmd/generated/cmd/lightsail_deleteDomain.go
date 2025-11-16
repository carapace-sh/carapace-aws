package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "Deletes the specified domain recordset and all of its domain records.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteDomainCmd).Standalone()

	lightsail_deleteDomainCmd.Flags().String("domain-name", "", "The specific domain name to delete.")
	lightsail_deleteDomainCmd.MarkFlagRequired("domain-name")
	lightsailCmd.AddCommand(lightsail_deleteDomainCmd)
}
