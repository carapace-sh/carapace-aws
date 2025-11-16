package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_listTransformersCmd = &cobra.Command{
	Use:   "list-transformers",
	Short: "Lists the available transformers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_listTransformersCmd).Standalone()

	b2bi_listTransformersCmd.Flags().String("max-results", "", "Specifies the number of items to return for the API response.")
	b2bi_listTransformersCmd.Flags().String("next-token", "", "When additional results are obtained from the command, a `NextToken` parameter is returned in the output.")
	b2biCmd.AddCommand(b2bi_listTransformersCmd)
}
