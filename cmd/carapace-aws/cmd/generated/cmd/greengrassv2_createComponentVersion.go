package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_createComponentVersionCmd = &cobra.Command{
	Use:   "create-component-version",
	Short: "Creates a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_createComponentVersionCmd).Standalone()

	greengrassv2_createComponentVersionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you can provide to ensure that the request is idempotent.")
	greengrassv2_createComponentVersionCmd.Flags().String("inline-recipe", "", "The recipe to use to create the component.")
	greengrassv2_createComponentVersionCmd.Flags().String("lambda-function", "", "The parameters to create a component from a Lambda function.")
	greengrassv2_createComponentVersionCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the resource.")
	greengrassv2Cmd.AddCommand(greengrassv2_createComponentVersionCmd)
}
