package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getResourcesSummaryCmd = &cobra.Command{
	Use:   "get-resources-summary",
	Short: "Get counts of Proton resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getResourcesSummaryCmd).Standalone()

	protonCmd.AddCommand(proton_getResourcesSummaryCmd)
}
