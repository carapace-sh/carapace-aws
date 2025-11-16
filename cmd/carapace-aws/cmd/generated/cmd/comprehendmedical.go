package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedicalCmd = &cobra.Command{
	Use:   "comprehendmedical",
	Short: "Amazon Comprehend Medical extracts structured information from unstructured clinical text.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedicalCmd).Standalone()

	rootCmd.AddCommand(comprehendmedicalCmd)
}
