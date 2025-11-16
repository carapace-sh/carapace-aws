package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastoreCmd = &cobra.Command{
	Use:   "mediastore",
	Short: "An AWS Elemental MediaStore container is a namespace that holds folders and objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastoreCmd).Standalone()

	})
	rootCmd.AddCommand(mediastoreCmd)
}
