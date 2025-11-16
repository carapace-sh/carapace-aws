package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_describeRepositoryCreationTemplatesCmd = &cobra.Command{
	Use:   "describe-repository-creation-templates",
	Short: "Returns details about the repository creation templates in a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_describeRepositoryCreationTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_describeRepositoryCreationTemplatesCmd).Standalone()

		ecr_describeRepositoryCreationTemplatesCmd.Flags().String("max-results", "", "The maximum number of repository results returned by `DescribeRepositoryCreationTemplatesRequest` in paginated output.")
		ecr_describeRepositoryCreationTemplatesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeRepositoryCreationTemplates` request where `maxResults` was used and the results exceeded the value of that parameter.")
		ecr_describeRepositoryCreationTemplatesCmd.Flags().String("prefixes", "", "The repository namespace prefixes associated with the repository creation templates to describe.")
	})
	ecrCmd.AddCommand(ecr_describeRepositoryCreationTemplatesCmd)
}
