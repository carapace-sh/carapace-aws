package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getContainerRecipePolicyCmd = &cobra.Command{
	Use:   "get-container-recipe-policy",
	Short: "Retrieves the policy for a container recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getContainerRecipePolicyCmd).Standalone()

	imagebuilder_getContainerRecipePolicyCmd.Flags().String("container-recipe-arn", "", "The Amazon Resource Name (ARN) of the container recipe for the policy being requested.")
	imagebuilder_getContainerRecipePolicyCmd.MarkFlagRequired("container-recipe-arn")
	imagebuilderCmd.AddCommand(imagebuilder_getContainerRecipePolicyCmd)
}
