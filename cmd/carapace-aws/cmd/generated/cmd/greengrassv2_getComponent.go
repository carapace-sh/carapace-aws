package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_getComponentCmd = &cobra.Command{
	Use:   "get-component",
	Short: "Gets the recipe for a version of a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_getComponentCmd).Standalone()

	greengrassv2_getComponentCmd.Flags().String("arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the component version.")
	greengrassv2_getComponentCmd.Flags().String("recipe-output-format", "", "The format of the recipe.")
	greengrassv2_getComponentCmd.MarkFlagRequired("arn")
	greengrassv2Cmd.AddCommand(greengrassv2_getComponentCmd)
}
