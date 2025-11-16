package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getLensVersionDifferenceCmd = &cobra.Command{
	Use:   "get-lens-version-difference",
	Short: "Get lens version differences.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getLensVersionDifferenceCmd).Standalone()

	wellarchitected_getLensVersionDifferenceCmd.Flags().String("base-lens-version", "", "The base version of the lens.")
	wellarchitected_getLensVersionDifferenceCmd.Flags().String("lens-alias", "", "")
	wellarchitected_getLensVersionDifferenceCmd.Flags().String("target-lens-version", "", "The lens version to target a difference for.")
	wellarchitected_getLensVersionDifferenceCmd.MarkFlagRequired("lens-alias")
	wellarchitectedCmd.AddCommand(wellarchitected_getLensVersionDifferenceCmd)
}
