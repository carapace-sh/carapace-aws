package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_describeFraudsterCmd = &cobra.Command{
	Use:   "describe-fraudster",
	Short: "Describes the specified fraudster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_describeFraudsterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_describeFraudsterCmd).Standalone()

		voiceId_describeFraudsterCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the fraudster.")
		voiceId_describeFraudsterCmd.Flags().String("fraudster-id", "", "The identifier of the fraudster you are describing.")
		voiceId_describeFraudsterCmd.MarkFlagRequired("domain-id")
		voiceId_describeFraudsterCmd.MarkFlagRequired("fraudster-id")
	})
	voiceIdCmd.AddCommand(voiceId_describeFraudsterCmd)
}
