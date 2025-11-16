package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a resource-based policy from an OpenSearch Ingestion resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_deleteResourcePolicyCmd).Standalone()

	osis_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which to delete the policy.")
	osis_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	osisCmd.AddCommand(osis_deleteResourcePolicyCmd)
}
