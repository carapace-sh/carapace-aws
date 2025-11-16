package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_resolveCaseCmd = &cobra.Command{
	Use:   "resolve-case",
	Short: "Resolves a support case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_resolveCaseCmd).Standalone()

	support_resolveCaseCmd.Flags().String("case-id", "", "The support case ID requested or returned in the call.")
	supportCmd.AddCommand(support_resolveCaseCmd)
}
