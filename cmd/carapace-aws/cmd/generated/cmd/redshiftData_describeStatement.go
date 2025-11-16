package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_describeStatementCmd = &cobra.Command{
	Use:   "describe-statement",
	Short: "Describes the details about a specific instance when a query was run by the Amazon Redshift Data API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_describeStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftData_describeStatementCmd).Standalone()

		redshiftData_describeStatementCmd.Flags().String("id", "", "The identifier of the SQL statement to describe.")
		redshiftData_describeStatementCmd.MarkFlagRequired("id")
	})
	redshiftDataCmd.AddCommand(redshiftData_describeStatementCmd)
}
