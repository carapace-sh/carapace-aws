package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based policy to an OpenSearch Ingestion resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_putResourcePolicyCmd).Standalone()

	osis_putResourcePolicyCmd.Flags().String("policy", "", "The resource-based policy document in JSON format.")
	osis_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to attach the policy to.")
	osis_putResourcePolicyCmd.MarkFlagRequired("policy")
	osis_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	osisCmd.AddCommand(osis_putResourcePolicyCmd)
}
