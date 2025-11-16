package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_describeDomainCmd = &cobra.Command{
	Use:   "describe-domain",
	Short: "Returns a [DomainDescription](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainDescription.html) object that contains information about the requested domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_describeDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_describeDomainCmd).Standalone()

		codeartifact_describeDomainCmd.Flags().String("domain", "", "A string that specifies the name of the requested domain.")
		codeartifact_describeDomainCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_describeDomainCmd.MarkFlagRequired("domain")
	})
	codeartifactCmd.AddCommand(codeartifact_describeDomainCmd)
}
