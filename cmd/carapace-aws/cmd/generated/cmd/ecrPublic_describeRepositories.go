package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_describeRepositoriesCmd = &cobra.Command{
	Use:   "describe-repositories",
	Short: "Describes repositories that are in a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_describeRepositoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_describeRepositoriesCmd).Standalone()

		ecrPublic_describeRepositoriesCmd.Flags().String("max-results", "", "The maximum number of repository results that's returned by `DescribeRepositories` in paginated output.")
		ecrPublic_describeRepositoriesCmd.Flags().String("next-token", "", "The `nextToken` value that's returned from a previous paginated `DescribeRepositories` request where `maxResults` was used and the results exceeded the value of that parameter.")
		ecrPublic_describeRepositoriesCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID that's associated with the registry that contains the repositories to be described.")
		ecrPublic_describeRepositoriesCmd.Flags().String("repository-names", "", "A list of repositories to describe.")
	})
	ecrPublicCmd.AddCommand(ecrPublic_describeRepositoriesCmd)
}
