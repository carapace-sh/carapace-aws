package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getClassificationExportConfigurationCmd = &cobra.Command{
	Use:   "get-classification-export-configuration",
	Short: "Retrieves the configuration settings for storing data classification results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getClassificationExportConfigurationCmd).Standalone()

	macie2Cmd.AddCommand(macie2_getClassificationExportConfigurationCmd)
}
