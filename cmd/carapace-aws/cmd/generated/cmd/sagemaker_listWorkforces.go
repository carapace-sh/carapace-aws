package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listWorkforcesCmd = &cobra.Command{
	Use:   "list-workforces",
	Short: "Use this operation to list all private and vendor workforces in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listWorkforcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listWorkforcesCmd).Standalone()

		sagemaker_listWorkforcesCmd.Flags().String("max-results", "", "The maximum number of workforces returned in the response.")
		sagemaker_listWorkforcesCmd.Flags().String("name-contains", "", "A filter you can use to search for workforces using part of the workforce name.")
		sagemaker_listWorkforcesCmd.Flags().String("next-token", "", "A token to resume pagination.")
		sagemaker_listWorkforcesCmd.Flags().String("sort-by", "", "Sort workforces using the workforce name or creation date.")
		sagemaker_listWorkforcesCmd.Flags().String("sort-order", "", "Sort workforces in ascending or descending order.")
	})
	sagemakerCmd.AddCommand(sagemaker_listWorkforcesCmd)
}
