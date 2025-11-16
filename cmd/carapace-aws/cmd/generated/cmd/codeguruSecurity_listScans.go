package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_listScansCmd = &cobra.Command{
	Use:   "list-scans",
	Short: "Returns a list of all scans in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_listScansCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruSecurity_listScansCmd).Standalone()

		codeguruSecurity_listScansCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		codeguruSecurity_listScansCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	})
	codeguruSecurityCmd.AddCommand(codeguruSecurity_listScansCmd)
}
