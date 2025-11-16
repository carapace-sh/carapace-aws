package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_describeRepositoriesCmd = &cobra.Command{
	Use:   "describe-repositories",
	Short: "Describes image repositories in a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_describeRepositoriesCmd).Standalone()

	ecr_describeRepositoriesCmd.Flags().String("max-results", "", "The maximum number of repository results returned by `DescribeRepositories` in paginated output.")
	ecr_describeRepositoriesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeRepositories` request where `maxResults` was used and the results exceeded the value of that parameter.")
	ecr_describeRepositoriesCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repositories to be described.")
	ecr_describeRepositoriesCmd.Flags().String("repository-names", "", "A list of repositories to describe.")
	ecrCmd.AddCommand(ecr_describeRepositoriesCmd)
}
