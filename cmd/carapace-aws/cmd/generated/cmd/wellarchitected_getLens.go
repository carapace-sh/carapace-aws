package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getLensCmd = &cobra.Command{
	Use:   "get-lens",
	Short: "Get an existing lens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getLensCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_getLensCmd).Standalone()

		wellarchitected_getLensCmd.Flags().String("lens-alias", "", "")
		wellarchitected_getLensCmd.Flags().String("lens-version", "", "The lens version to be retrieved.")
		wellarchitected_getLensCmd.MarkFlagRequired("lens-alias")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_getLensCmd)
}
