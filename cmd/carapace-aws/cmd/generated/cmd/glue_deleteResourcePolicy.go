package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a specified policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteResourcePolicyCmd).Standalone()

		glue_deleteResourcePolicyCmd.Flags().String("policy-hash-condition", "", "The hash value returned when this policy was set.")
		glue_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the Glue resource for the resource policy to be deleted.")
	})
	glueCmd.AddCommand(glue_deleteResourcePolicyCmd)
}
