package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "Permanently deletes a search domain and all of its data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_deleteDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_deleteDomainCmd).Standalone()

		cloudsearch_deleteDomainCmd.Flags().String("domain-name", "", "The name of the domain you want to permanently delete.")
		cloudsearch_deleteDomainCmd.MarkFlagRequired("domain-name")
	})
	cloudsearchCmd.AddCommand(cloudsearch_deleteDomainCmd)
}
