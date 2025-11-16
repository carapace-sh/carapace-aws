package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putConfigurationSetSuppressionOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-suppression-options",
	Short: "Specify the account suppression list preferences for a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putConfigurationSetSuppressionOptionsCmd).Standalone()

	sesv2_putConfigurationSetSuppressionOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to change the suppression list preferences for.")
	sesv2_putConfigurationSetSuppressionOptionsCmd.Flags().String("suppressed-reasons", "", "A list that contains the reasons that email addresses are automatically added to the suppression list for your account.")
	sesv2_putConfigurationSetSuppressionOptionsCmd.MarkFlagRequired("configuration-set-name")
	sesv2Cmd.AddCommand(sesv2_putConfigurationSetSuppressionOptionsCmd)
}
