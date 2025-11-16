package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeRuntimeCmd = &cobra.Command{
	Use:   "personalize-runtime",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeRuntimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalizeRuntimeCmd).Standalone()

	})
	rootCmd.AddCommand(personalizeRuntimeCmd)
}
