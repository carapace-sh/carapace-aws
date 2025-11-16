package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the resource policy that Resource Access Manager uses to share your Incident Manager resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_deleteResourcePolicyCmd).Standalone()

	ssmIncidents_deleteResourcePolicyCmd.Flags().String("policy-id", "", "The ID of the resource policy you're deleting.")
	ssmIncidents_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource you're deleting the policy from.")
	ssmIncidents_deleteResourcePolicyCmd.MarkFlagRequired("policy-id")
	ssmIncidents_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_deleteResourcePolicyCmd)
}
