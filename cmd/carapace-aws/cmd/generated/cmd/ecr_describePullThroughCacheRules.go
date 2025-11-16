package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_describePullThroughCacheRulesCmd = &cobra.Command{
	Use:   "describe-pull-through-cache-rules",
	Short: "Returns the pull through cache rules for a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_describePullThroughCacheRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_describePullThroughCacheRulesCmd).Standalone()

		ecr_describePullThroughCacheRulesCmd.Flags().String("ecr-repository-prefixes", "", "The Amazon ECR repository prefixes associated with the pull through cache rules to return.")
		ecr_describePullThroughCacheRulesCmd.Flags().String("max-results", "", "The maximum number of pull through cache rules returned by `DescribePullThroughCacheRulesRequest` in paginated output.")
		ecr_describePullThroughCacheRulesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribePullThroughCacheRulesRequest` request where `maxResults` was used and the results exceeded the value of that parameter.")
		ecr_describePullThroughCacheRulesCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry to return the pull through cache rules for.")
	})
	ecrCmd.AddCommand(ecr_describePullThroughCacheRulesCmd)
}
