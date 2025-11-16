package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_listLayoutsCmd = &cobra.Command{
	Use:   "list-layouts",
	Short: "Lists all layouts in the given cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_listLayoutsCmd).Standalone()

	connectcases_listLayoutsCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_listLayoutsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connectcases_listLayoutsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcases_listLayoutsCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_listLayoutsCmd)
}
