package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_describeComponentCmd = &cobra.Command{
	Use:   "describe-component",
	Short: "Retrieves metadata for a version of a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_describeComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_describeComponentCmd).Standalone()

		greengrassv2_describeComponentCmd.Flags().String("arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the component version.")
		greengrassv2_describeComponentCmd.MarkFlagRequired("arn")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_describeComponentCmd)
}
