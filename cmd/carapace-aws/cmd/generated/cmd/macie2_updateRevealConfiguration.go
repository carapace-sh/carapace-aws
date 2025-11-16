package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateRevealConfigurationCmd = &cobra.Command{
	Use:   "update-reveal-configuration",
	Short: "Updates the status and configuration settings for retrieving occurrences of sensitive data reported by findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateRevealConfigurationCmd).Standalone()

	macie2_updateRevealConfigurationCmd.Flags().String("configuration", "", "The KMS key to use to encrypt the sensitive data, and the status of the configuration for the Amazon Macie account.")
	macie2_updateRevealConfigurationCmd.Flags().String("retrieval-configuration", "", "The access method and settings to use when retrieving the sensitive data.")
	macie2_updateRevealConfigurationCmd.MarkFlagRequired("configuration")
	macie2Cmd.AddCommand(macie2_updateRevealConfigurationCmd)
}
