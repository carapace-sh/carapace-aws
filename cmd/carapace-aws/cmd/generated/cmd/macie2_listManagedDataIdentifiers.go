package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listManagedDataIdentifiersCmd = &cobra.Command{
	Use:   "list-managed-data-identifiers",
	Short: "Retrieves information about all the managed data identifiers that Amazon Macie currently provides.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listManagedDataIdentifiersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_listManagedDataIdentifiersCmd).Standalone()

		macie2_listManagedDataIdentifiersCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	})
	macie2Cmd.AddCommand(macie2_listManagedDataIdentifiersCmd)
}
