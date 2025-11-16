package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "Deletes a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_deleteDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_deleteDomainCmd).Standalone()

		codeartifact_deleteDomainCmd.Flags().String("domain", "", "The name of the domain to delete.")
		codeartifact_deleteDomainCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_deleteDomainCmd.MarkFlagRequired("domain")
	})
	codeartifactCmd.AddCommand(codeartifact_deleteDomainCmd)
}
