package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getLifecyclePolicyCmd = &cobra.Command{
	Use:   "get-lifecycle-policy",
	Short: "Get details for the specified image lifecycle policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_getLifecyclePolicyCmd).Standalone()

		imagebuilder_getLifecyclePolicyCmd.Flags().String("lifecycle-policy-arn", "", "Specifies the Amazon Resource Name (ARN) of the image lifecycle policy resource to get.")
		imagebuilder_getLifecyclePolicyCmd.MarkFlagRequired("lifecycle-policy-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_getLifecyclePolicyCmd)
}
