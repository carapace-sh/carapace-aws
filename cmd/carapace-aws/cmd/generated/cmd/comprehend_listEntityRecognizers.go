package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listEntityRecognizersCmd = &cobra.Command{
	Use:   "list-entity-recognizers",
	Short: "Gets a list of the properties of all entity recognizers that you created, including recognizers currently in training.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listEntityRecognizersCmd).Standalone()

	comprehend_listEntityRecognizersCmd.Flags().String("filter", "", "Filters the list of entities returned.")
	comprehend_listEntityRecognizersCmd.Flags().String("max-results", "", "The maximum number of results to return on each page.")
	comprehend_listEntityRecognizersCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendCmd.AddCommand(comprehend_listEntityRecognizersCmd)
}
