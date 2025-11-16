package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeOpsItemsCmd = &cobra.Command{
	Use:   "describe-ops-items",
	Short: "Query a set of OpsItems.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeOpsItemsCmd).Standalone()

	ssm_describeOpsItemsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeOpsItemsCmd.Flags().String("next-token", "", "A token to start the list.")
	ssm_describeOpsItemsCmd.Flags().String("ops-item-filters", "", "One or more filters to limit the response.")
	ssmCmd.AddCommand(ssm_describeOpsItemsCmd)
}
