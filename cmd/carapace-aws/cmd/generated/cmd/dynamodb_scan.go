package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "The `Scan` operation returns one or more items and item attributes by accessing every item in a table or a secondary index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_scanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_scanCmd).Standalone()

		dynamodb_scanCmd.Flags().String("attributes-to-get", "", "This is a legacy parameter.")
		dynamodb_scanCmd.Flags().String("conditional-operator", "", "This is a legacy parameter.")
		dynamodb_scanCmd.Flags().String("consistent-read", "", "A Boolean value that determines the read consistency model during the scan:")
		dynamodb_scanCmd.Flags().String("exclusive-start-key", "", "The primary key of the first item that this operation will evaluate.")
		dynamodb_scanCmd.Flags().String("expression-attribute-names", "", "One or more substitution tokens for attribute names in an expression.")
		dynamodb_scanCmd.Flags().String("expression-attribute-values", "", "One or more values that can be substituted in an expression.")
		dynamodb_scanCmd.Flags().String("filter-expression", "", "A string that contains conditions that DynamoDB applies after the `Scan` operation, but before the data is returned to you.")
		dynamodb_scanCmd.Flags().String("index-name", "", "The name of a secondary index to scan.")
		dynamodb_scanCmd.Flags().String("limit", "", "The maximum number of items to evaluate (not necessarily the number of matching items).")
		dynamodb_scanCmd.Flags().String("projection-expression", "", "A string that identifies one or more attributes to retrieve from the specified table or index.")
		dynamodb_scanCmd.Flags().String("return-consumed-capacity", "", "")
		dynamodb_scanCmd.Flags().String("scan-filter", "", "This is a legacy parameter.")
		dynamodb_scanCmd.Flags().String("segment", "", "For a parallel `Scan` request, `Segment` identifies an individual segment to be scanned by an application worker.")
		dynamodb_scanCmd.Flags().String("select", "", "The attributes to be returned in the result.")
		dynamodb_scanCmd.Flags().String("table-name", "", "The name of the table containing the requested items or if you provide `IndexName`, the name of the table to which that index belongs.")
		dynamodb_scanCmd.Flags().String("total-segments", "", "For a parallel `Scan` request, `TotalSegments` represents the total number of segments into which the `Scan` operation will be divided.")
		dynamodb_scanCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_scanCmd)
}
