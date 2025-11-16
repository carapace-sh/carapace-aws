package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_describeFraudsterRegistrationJobCmd = &cobra.Command{
	Use:   "describe-fraudster-registration-job",
	Short: "Describes the specified fraudster registration job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_describeFraudsterRegistrationJobCmd).Standalone()

	voiceId_describeFraudsterRegistrationJobCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the fraudster registration job.")
	voiceId_describeFraudsterRegistrationJobCmd.Flags().String("job-id", "", "The identifier of the fraudster registration job you are describing.")
	voiceId_describeFraudsterRegistrationJobCmd.MarkFlagRequired("domain-id")
	voiceId_describeFraudsterRegistrationJobCmd.MarkFlagRequired("job-id")
	voiceIdCmd.AddCommand(voiceId_describeFraudsterRegistrationJobCmd)
}
