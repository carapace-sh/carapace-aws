package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates a new search domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_createDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_createDomainCmd).Standalone()

		cloudsearch_createDomainCmd.Flags().String("domain-name", "", "A name for the domain you are creating.")
		cloudsearch_createDomainCmd.MarkFlagRequired("domain-name")
	})
	cloudsearchCmd.AddCommand(cloudsearch_createDomainCmd)
}
