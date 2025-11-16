package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listExportsCmd = &cobra.Command{
	Use:   "list-exports",
	Short: "List exports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listExportsCmd).Standalone()

	mgn_listExportsCmd.Flags().String("filters", "", "")
	mgn_listExportsCmd.Flags().String("max-results", "", "List export request max results.")
	mgn_listExportsCmd.Flags().String("next-token", "", "List export request next token.")
	mgnCmd.AddCommand(mgn_listExportsCmd)
}
