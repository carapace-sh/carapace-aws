package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_listResourceTypesCmd = &cobra.Command{
	Use:   "list-resource-types",
	Short: "Lists the resource types that can be shared by RAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_listResourceTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_listResourceTypesCmd).Standalone()

		ram_listResourceTypesCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_listResourceTypesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_listResourceTypesCmd.Flags().String("resource-region-scope", "", "Specifies that you want the results to include only resources that have the specified scope.")
	})
	ramCmd.AddCommand(ram_listResourceTypesCmd)
}
