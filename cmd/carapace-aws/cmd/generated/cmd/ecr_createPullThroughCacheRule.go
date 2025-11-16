package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_createPullThroughCacheRuleCmd = &cobra.Command{
	Use:   "create-pull-through-cache-rule",
	Short: "Creates a pull through cache rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_createPullThroughCacheRuleCmd).Standalone()

	ecr_createPullThroughCacheRuleCmd.Flags().String("credential-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services Secrets Manager secret that identifies the credentials to authenticate to the upstream registry.")
	ecr_createPullThroughCacheRuleCmd.Flags().String("custom-role-arn", "", "Amazon Resource Name (ARN) of the IAM role to be assumed by Amazon ECR to authenticate to the ECR upstream registry.")
	ecr_createPullThroughCacheRuleCmd.Flags().String("ecr-repository-prefix", "", "The repository name prefix to use when caching images from the source registry.")
	ecr_createPullThroughCacheRuleCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry to create the pull through cache rule for.")
	ecr_createPullThroughCacheRuleCmd.Flags().String("upstream-registry", "", "The name of the upstream registry.")
	ecr_createPullThroughCacheRuleCmd.Flags().String("upstream-registry-url", "", "The registry URL of the upstream public registry to use as the source for the pull through cache rule.")
	ecr_createPullThroughCacheRuleCmd.Flags().String("upstream-repository-prefix", "", "The repository name prefix of the upstream registry to match with the upstream repository name.")
	ecr_createPullThroughCacheRuleCmd.MarkFlagRequired("ecr-repository-prefix")
	ecr_createPullThroughCacheRuleCmd.MarkFlagRequired("upstream-registry-url")
	ecrCmd.AddCommand(ecr_createPullThroughCacheRuleCmd)
}
