package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "Deletes a Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_deleteDomainCmd).Standalone()

	connectcases_deleteDomainCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_deleteDomainCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_deleteDomainCmd)
}
