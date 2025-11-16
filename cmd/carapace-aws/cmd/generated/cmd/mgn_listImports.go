package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listImportsCmd = &cobra.Command{
	Use:   "list-imports",
	Short: "List imports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listImportsCmd).Standalone()

	mgn_listImportsCmd.Flags().String("filters", "", "List imports request filters.")
	mgn_listImportsCmd.Flags().String("max-results", "", "List imports request max results.")
	mgn_listImportsCmd.Flags().String("next-token", "", "List imports request next token.")
	mgnCmd.AddCommand(mgn_listImportsCmd)
}
