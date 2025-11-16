package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "Deletes the specified domain from Voice ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_deleteDomainCmd).Standalone()

	voiceId_deleteDomainCmd.Flags().String("domain-id", "", "The identifier of the domain you want to delete.")
	voiceId_deleteDomainCmd.MarkFlagRequired("domain-id")
	voiceIdCmd.AddCommand(voiceId_deleteDomainCmd)
}
