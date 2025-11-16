package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_listTargetResourceTypesCmd = &cobra.Command{
	Use:   "list-target-resource-types",
	Short: "Lists the target resource types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_listTargetResourceTypesCmd).Standalone()

	fis_listTargetResourceTypesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	fis_listTargetResourceTypesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	fisCmd.AddCommand(fis_listTargetResourceTypesCmd)
}
