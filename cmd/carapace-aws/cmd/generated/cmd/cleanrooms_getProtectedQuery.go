package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getProtectedQueryCmd = &cobra.Command{
	Use:   "get-protected-query",
	Short: "Returns query processing metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getProtectedQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getProtectedQueryCmd).Standalone()

		cleanrooms_getProtectedQueryCmd.Flags().String("membership-identifier", "", "The identifier for a membership in a protected query instance.")
		cleanrooms_getProtectedQueryCmd.Flags().String("protected-query-identifier", "", "The identifier for a protected query instance.")
		cleanrooms_getProtectedQueryCmd.MarkFlagRequired("membership-identifier")
		cleanrooms_getProtectedQueryCmd.MarkFlagRequired("protected-query-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getProtectedQueryCmd)
}
