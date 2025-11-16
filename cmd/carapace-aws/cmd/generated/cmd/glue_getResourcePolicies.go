package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getResourcePoliciesCmd = &cobra.Command{
	Use:   "get-resource-policies",
	Short: "Retrieves the resource policies set on individual resources by Resource Access Manager during cross-account permission grants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getResourcePoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getResourcePoliciesCmd).Standalone()

		glue_getResourcePoliciesCmd.Flags().String("max-results", "", "The maximum size of a list to return.")
		glue_getResourcePoliciesCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation request.")
	})
	glueCmd.AddCommand(glue_getResourcePoliciesCmd)
}
