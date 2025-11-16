package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_listAdapterVersionsCmd = &cobra.Command{
	Use:   "list-adapter-versions",
	Short: "List all version of an adapter that meet the specified filtration criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_listAdapterVersionsCmd).Standalone()

	textract_listAdapterVersionsCmd.Flags().String("adapter-id", "", "A string containing a unique ID for the adapter to match for when listing adapter versions.")
	textract_listAdapterVersionsCmd.Flags().String("after-creation-time", "", "Specifies the lower bound for the ListAdapterVersions operation.")
	textract_listAdapterVersionsCmd.Flags().String("before-creation-time", "", "Specifies the upper bound for the ListAdapterVersions operation.")
	textract_listAdapterVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return when listing adapter versions.")
	textract_listAdapterVersionsCmd.Flags().String("next-token", "", "Identifies the next page of results to return when listing adapter versions.")
	textractCmd.AddCommand(textract_listAdapterVersionsCmd)
}
