package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_updatePullThroughCacheRuleCmd = &cobra.Command{
	Use:   "update-pull-through-cache-rule",
	Short: "Updates an existing pull through cache rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_updatePullThroughCacheRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_updatePullThroughCacheRuleCmd).Standalone()

		ecr_updatePullThroughCacheRuleCmd.Flags().String("credential-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services Secrets Manager secret that identifies the credentials to authenticate to the upstream registry.")
		ecr_updatePullThroughCacheRuleCmd.Flags().String("custom-role-arn", "", "Amazon Resource Name (ARN) of the IAM role to be assumed by Amazon ECR to authenticate to the ECR upstream registry.")
		ecr_updatePullThroughCacheRuleCmd.Flags().String("ecr-repository-prefix", "", "The repository name prefix to use when caching images from the source registry.")
		ecr_updatePullThroughCacheRuleCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry associated with the pull through cache rule.")
		ecr_updatePullThroughCacheRuleCmd.MarkFlagRequired("ecr-repository-prefix")
	})
	ecrCmd.AddCommand(ecr_updatePullThroughCacheRuleCmd)
}
