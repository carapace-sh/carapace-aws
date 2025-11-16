package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_batchGetCommitsCmd = &cobra.Command{
	Use:   "batch-get-commits",
	Short: "Returns information about the contents of one or more commits in a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_batchGetCommitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_batchGetCommitsCmd).Standalone()

		codecommit_batchGetCommitsCmd.Flags().String("commit-ids", "", "The full commit IDs of the commits to get information about.")
		codecommit_batchGetCommitsCmd.Flags().String("repository-name", "", "The name of the repository that contains the commits.")
		codecommit_batchGetCommitsCmd.MarkFlagRequired("commit-ids")
		codecommit_batchGetCommitsCmd.MarkFlagRequired("repository-name")
	})
	codecommitCmd.AddCommand(codecommit_batchGetCommitsCmd)
}
