package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listLensSharesCmd = &cobra.Command{
	Use:   "list-lens-shares",
	Short: "List the lens shares associated with the lens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listLensSharesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_listLensSharesCmd).Standalone()

		wellarchitected_listLensSharesCmd.Flags().String("lens-alias", "", "")
		wellarchitected_listLensSharesCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
		wellarchitected_listLensSharesCmd.Flags().String("next-token", "", "")
		wellarchitected_listLensSharesCmd.Flags().String("shared-with-prefix", "", "The Amazon Web Services account ID, organization ID, or organizational unit (OU) ID with which the lens is shared.")
		wellarchitected_listLensSharesCmd.Flags().String("status", "", "")
		wellarchitected_listLensSharesCmd.MarkFlagRequired("lens-alias")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_listLensSharesCmd)
}
