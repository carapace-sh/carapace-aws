package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_disassociateExternalConnectionCmd = &cobra.Command{
	Use:   "disassociate-external-connection",
	Short: "Removes an existing external connection from a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_disassociateExternalConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_disassociateExternalConnectionCmd).Standalone()

		codeartifact_disassociateExternalConnectionCmd.Flags().String("domain", "", "The name of the domain that contains the repository from which to remove the external repository.")
		codeartifact_disassociateExternalConnectionCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_disassociateExternalConnectionCmd.Flags().String("external-connection", "", "The name of the external connection to be removed from the repository.")
		codeartifact_disassociateExternalConnectionCmd.Flags().String("repository", "", "The name of the repository from which the external connection will be removed.")
		codeartifact_disassociateExternalConnectionCmd.MarkFlagRequired("domain")
		codeartifact_disassociateExternalConnectionCmd.MarkFlagRequired("external-connection")
		codeartifact_disassociateExternalConnectionCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_disassociateExternalConnectionCmd)
}
