package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_batchGetFindingsCmd = &cobra.Command{
	Use:   "batch-get-findings",
	Short: "Returns a list of requested findings from standard scans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_batchGetFindingsCmd).Standalone()

	codeguruSecurity_batchGetFindingsCmd.Flags().String("finding-identifiers", "", "A list of finding identifiers.")
	codeguruSecurity_batchGetFindingsCmd.MarkFlagRequired("finding-identifiers")
	codeguruSecurityCmd.AddCommand(codeguruSecurity_batchGetFindingsCmd)
}
