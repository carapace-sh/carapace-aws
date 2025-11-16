package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_putClassificationExportConfigurationCmd = &cobra.Command{
	Use:   "put-classification-export-configuration",
	Short: "Adds or updates the configuration settings for storing data classification results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_putClassificationExportConfigurationCmd).Standalone()

	macie2_putClassificationExportConfigurationCmd.Flags().String("configuration", "", "The location to store data classification results in, and the encryption settings to use when storing results in that location.")
	macie2_putClassificationExportConfigurationCmd.MarkFlagRequired("configuration")
	macie2Cmd.AddCommand(macie2_putClassificationExportConfigurationCmd)
}
