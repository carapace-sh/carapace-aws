package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listIndicesCmd = &cobra.Command{
	Use:   "list-indices",
	Short: "Lists the Amazon Kendra indexes that you created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listIndicesCmd).Standalone()

	kendra_listIndicesCmd.Flags().String("max-results", "", "The maximum number of indices to return.")
	kendra_listIndicesCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
	kendraCmd.AddCommand(kendra_listIndicesCmd)
}
