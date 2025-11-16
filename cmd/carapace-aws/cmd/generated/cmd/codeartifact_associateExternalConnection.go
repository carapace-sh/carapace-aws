package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_associateExternalConnectionCmd = &cobra.Command{
	Use:   "associate-external-connection",
	Short: "Adds an existing external connection to a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_associateExternalConnectionCmd).Standalone()

	codeartifact_associateExternalConnectionCmd.Flags().String("domain", "", "The name of the domain that contains the repository.")
	codeartifact_associateExternalConnectionCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_associateExternalConnectionCmd.Flags().String("external-connection", "", "The name of the external connection to add to the repository.")
	codeartifact_associateExternalConnectionCmd.Flags().String("repository", "", "The name of the repository to which the external connection is added.")
	codeartifact_associateExternalConnectionCmd.MarkFlagRequired("domain")
	codeartifact_associateExternalConnectionCmd.MarkFlagRequired("external-connection")
	codeartifact_associateExternalConnectionCmd.MarkFlagRequired("repository")
	codeartifactCmd.AddCommand(codeartifact_associateExternalConnectionCmd)
}
