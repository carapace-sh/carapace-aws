package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_createDomainCmd).Standalone()

	codeartifact_createDomainCmd.Flags().String("domain", "", "The name of the domain to create.")
	codeartifact_createDomainCmd.Flags().String("encryption-key", "", "The encryption key for the domain.")
	codeartifact_createDomainCmd.Flags().String("tags", "", "One or more tag key-value pairs for the domain.")
	codeartifact_createDomainCmd.MarkFlagRequired("domain")
	codeartifactCmd.AddCommand(codeartifact_createDomainCmd)
}
