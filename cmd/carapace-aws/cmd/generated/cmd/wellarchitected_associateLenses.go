package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_associateLensesCmd = &cobra.Command{
	Use:   "associate-lenses",
	Short: "Associate a lens to a workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_associateLensesCmd).Standalone()

	wellarchitected_associateLensesCmd.Flags().String("lens-aliases", "", "")
	wellarchitected_associateLensesCmd.Flags().String("workload-id", "", "")
	wellarchitected_associateLensesCmd.MarkFlagRequired("lens-aliases")
	wellarchitected_associateLensesCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_associateLensesCmd)
}
