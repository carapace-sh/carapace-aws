package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_deleteComponentCmd = &cobra.Command{
	Use:   "delete-component",
	Short: "Deletes a version of a component from IoT Greengrass.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_deleteComponentCmd).Standalone()

	greengrassv2_deleteComponentCmd.Flags().String("arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the component version.")
	greengrassv2_deleteComponentCmd.MarkFlagRequired("arn")
	greengrassv2Cmd.AddCommand(greengrassv2_deleteComponentCmd)
}
