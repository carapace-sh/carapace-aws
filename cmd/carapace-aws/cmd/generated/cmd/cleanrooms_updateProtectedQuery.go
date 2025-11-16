package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateProtectedQueryCmd = &cobra.Command{
	Use:   "update-protected-query",
	Short: "Updates the processing of a currently running query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateProtectedQueryCmd).Standalone()

	cleanrooms_updateProtectedQueryCmd.Flags().String("membership-identifier", "", "The identifier for a member of a protected query instance.")
	cleanrooms_updateProtectedQueryCmd.Flags().String("protected-query-identifier", "", "The identifier for a protected query instance.")
	cleanrooms_updateProtectedQueryCmd.Flags().String("target-status", "", "The target status of a query.")
	cleanrooms_updateProtectedQueryCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_updateProtectedQueryCmd.MarkFlagRequired("protected-query-identifier")
	cleanrooms_updateProtectedQueryCmd.MarkFlagRequired("target-status")
	cleanroomsCmd.AddCommand(cleanrooms_updateProtectedQueryCmd)
}
