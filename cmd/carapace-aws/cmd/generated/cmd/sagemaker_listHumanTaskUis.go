package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listHumanTaskUisCmd = &cobra.Command{
	Use:   "list-human-task-uis",
	Short: "Returns information about the human task user interfaces in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listHumanTaskUisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listHumanTaskUisCmd).Standalone()

		sagemaker_listHumanTaskUisCmd.Flags().String("creation-time-after", "", "A filter that returns only human task user interfaces with a creation time greater than or equal to the specified timestamp.")
		sagemaker_listHumanTaskUisCmd.Flags().String("creation-time-before", "", "A filter that returns only human task user interfaces that were created before the specified timestamp.")
		sagemaker_listHumanTaskUisCmd.Flags().String("max-results", "", "The total number of items to return.")
		sagemaker_listHumanTaskUisCmd.Flags().String("next-token", "", "A token to resume pagination.")
		sagemaker_listHumanTaskUisCmd.Flags().String("sort-order", "", "An optional value that specifies whether you want the results sorted in `Ascending` or `Descending` order.")
	})
	sagemakerCmd.AddCommand(sagemaker_listHumanTaskUisCmd)
}
