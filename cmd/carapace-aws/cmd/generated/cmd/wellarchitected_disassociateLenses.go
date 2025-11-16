package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_disassociateLensesCmd = &cobra.Command{
	Use:   "disassociate-lenses",
	Short: "Disassociate a lens from a workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_disassociateLensesCmd).Standalone()

	wellarchitected_disassociateLensesCmd.Flags().String("lens-aliases", "", "")
	wellarchitected_disassociateLensesCmd.Flags().String("workload-id", "", "")
	wellarchitected_disassociateLensesCmd.MarkFlagRequired("lens-aliases")
	wellarchitected_disassociateLensesCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_disassociateLensesCmd)
}
