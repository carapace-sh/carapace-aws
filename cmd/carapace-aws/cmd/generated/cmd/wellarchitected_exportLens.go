package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_exportLensCmd = &cobra.Command{
	Use:   "export-lens",
	Short: "Export an existing lens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_exportLensCmd).Standalone()

	wellarchitected_exportLensCmd.Flags().String("lens-alias", "", "")
	wellarchitected_exportLensCmd.Flags().String("lens-version", "", "The lens version to be exported.")
	wellarchitected_exportLensCmd.MarkFlagRequired("lens-alias")
	wellarchitectedCmd.AddCommand(wellarchitected_exportLensCmd)
}
