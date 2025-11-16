package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Returns the resource-based policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_getPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_getPolicyCmd).Standalone()

		entityresolution_getPolicyCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource for which the policy need to be returned.")
		entityresolution_getPolicyCmd.MarkFlagRequired("arn")
	})
	entityresolutionCmd.AddCommand(entityresolution_getPolicyCmd)
}
