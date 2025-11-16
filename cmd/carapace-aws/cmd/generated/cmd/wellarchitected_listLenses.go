package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listLensesCmd = &cobra.Command{
	Use:   "list-lenses",
	Short: "List the available lenses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listLensesCmd).Standalone()

	wellarchitected_listLensesCmd.Flags().String("lens-name", "", "")
	wellarchitected_listLensesCmd.Flags().String("lens-status", "", "The status of lenses to be returned.")
	wellarchitected_listLensesCmd.Flags().String("lens-type", "", "The type of lenses to be returned.")
	wellarchitected_listLensesCmd.Flags().String("max-results", "", "")
	wellarchitected_listLensesCmd.Flags().String("next-token", "", "")
	wellarchitectedCmd.AddCommand(wellarchitected_listLensesCmd)
}
