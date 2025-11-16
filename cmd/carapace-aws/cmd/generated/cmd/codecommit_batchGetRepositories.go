package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_batchGetRepositoriesCmd = &cobra.Command{
	Use:   "batch-get-repositories",
	Short: "Returns information about one or more repositories.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_batchGetRepositoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_batchGetRepositoriesCmd).Standalone()

		codecommit_batchGetRepositoriesCmd.Flags().String("repository-names", "", "The names of the repositories to get information about.")
		codecommit_batchGetRepositoriesCmd.MarkFlagRequired("repository-names")
	})
	codecommitCmd.AddCommand(codecommit_batchGetRepositoriesCmd)
}
