package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves the resource-based policy attached to an OpenSearch Ingestion resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_getResourcePolicyCmd).Standalone()

	osis_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which to retrieve the policy.")
	osis_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	osisCmd.AddCommand(osis_getResourcePolicyCmd)
}
