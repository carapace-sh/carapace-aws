package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getComponentPolicyCmd = &cobra.Command{
	Use:   "get-component-policy",
	Short: "Gets a component policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getComponentPolicyCmd).Standalone()

	imagebuilder_getComponentPolicyCmd.Flags().String("component-arn", "", "The Amazon Resource Name (ARN) of the component whose policy you want to retrieve.")
	imagebuilder_getComponentPolicyCmd.MarkFlagRequired("component-arn")
	imagebuilderCmd.AddCommand(imagebuilder_getComponentPolicyCmd)
}
