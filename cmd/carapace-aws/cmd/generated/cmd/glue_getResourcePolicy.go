package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves a specified resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getResourcePolicyCmd).Standalone()

	glue_getResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the Glue resource for which to retrieve the resource policy.")
	glueCmd.AddCommand(glue_getResourcePolicyCmd)
}
