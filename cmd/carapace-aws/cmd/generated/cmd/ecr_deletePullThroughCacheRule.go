package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_deletePullThroughCacheRuleCmd = &cobra.Command{
	Use:   "delete-pull-through-cache-rule",
	Short: "Deletes a pull through cache rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_deletePullThroughCacheRuleCmd).Standalone()

	ecr_deletePullThroughCacheRuleCmd.Flags().String("ecr-repository-prefix", "", "The Amazon ECR repository prefix associated with the pull through cache rule to delete.")
	ecr_deletePullThroughCacheRuleCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the pull through cache rule.")
	ecr_deletePullThroughCacheRuleCmd.MarkFlagRequired("ecr-repository-prefix")
	ecrCmd.AddCommand(ecr_deletePullThroughCacheRuleCmd)
}
