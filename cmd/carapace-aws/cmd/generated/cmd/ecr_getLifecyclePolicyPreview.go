package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_getLifecyclePolicyPreviewCmd = &cobra.Command{
	Use:   "get-lifecycle-policy-preview",
	Short: "Retrieves the results of the lifecycle policy preview request for the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_getLifecyclePolicyPreviewCmd).Standalone()

	ecr_getLifecyclePolicyPreviewCmd.Flags().String("filter", "", "An optional parameter that filters results based on image tag status and all tags, if tagged.")
	ecr_getLifecyclePolicyPreviewCmd.Flags().String("image-ids", "", "The list of imageIDs to be included.")
	ecr_getLifecyclePolicyPreviewCmd.Flags().String("max-results", "", "The maximum number of repository results returned by `GetLifecyclePolicyPreviewRequest` in\u2028 paginated output.")
	ecr_getLifecyclePolicyPreviewCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated\u2028 `GetLifecyclePolicyPreviewRequest` request where `maxResults` was used and the\u2028 results exceeded the value of that parameter.")
	ecr_getLifecyclePolicyPreviewCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository.")
	ecr_getLifecyclePolicyPreviewCmd.Flags().String("repository-name", "", "The name of the repository.")
	ecr_getLifecyclePolicyPreviewCmd.MarkFlagRequired("repository-name")
	ecrCmd.AddCommand(ecr_getLifecyclePolicyPreviewCmd)
}
