package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_getPackageVersionHistoryCmd = &cobra.Command{
	Use:   "get-package-version-history",
	Short: "Returns a list of versions of the package, along with their creation time and commit message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_getPackageVersionHistoryCmd).Standalone()

	es_getPackageVersionHistoryCmd.Flags().String("max-results", "", "Limits results to a maximum number of versions.")
	es_getPackageVersionHistoryCmd.Flags().String("next-token", "", "Used for pagination.")
	es_getPackageVersionHistoryCmd.Flags().String("package-id", "", "Returns an audit history of versions of the package.")
	es_getPackageVersionHistoryCmd.MarkFlagRequired("package-id")
	esCmd.AddCommand(es_getPackageVersionHistoryCmd)
}
