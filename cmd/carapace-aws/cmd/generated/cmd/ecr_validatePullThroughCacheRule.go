package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_validatePullThroughCacheRuleCmd = &cobra.Command{
	Use:   "validate-pull-through-cache-rule",
	Short: "Validates an existing pull through cache rule for an upstream registry that requires authentication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_validatePullThroughCacheRuleCmd).Standalone()

	ecr_validatePullThroughCacheRuleCmd.Flags().String("ecr-repository-prefix", "", "The repository name prefix associated with the pull through cache rule.")
	ecr_validatePullThroughCacheRuleCmd.Flags().String("registry-id", "", "The registry ID associated with the pull through cache rule.")
	ecr_validatePullThroughCacheRuleCmd.MarkFlagRequired("ecr-repository-prefix")
	ecrCmd.AddCommand(ecr_validatePullThroughCacheRuleCmd)
}
