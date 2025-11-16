package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_describeDomainCmd = &cobra.Command{
	Use:   "describe-domain",
	Short: "Describes the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_describeDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_describeDomainCmd).Standalone()

		voiceId_describeDomainCmd.Flags().String("domain-id", "", "The identifier of the domain that you are describing.")
		voiceId_describeDomainCmd.MarkFlagRequired("domain-id")
	})
	voiceIdCmd.AddCommand(voiceId_describeDomainCmd)
}
