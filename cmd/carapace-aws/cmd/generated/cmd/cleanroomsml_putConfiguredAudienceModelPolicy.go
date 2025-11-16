package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_putConfiguredAudienceModelPolicyCmd = &cobra.Command{
	Use:   "put-configured-audience-model-policy",
	Short: "Create or update the resource policy for a configured audience model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_putConfiguredAudienceModelPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_putConfiguredAudienceModelPolicyCmd).Standalone()

		cleanroomsml_putConfiguredAudienceModelPolicyCmd.Flags().String("configured-audience-model-arn", "", "The Amazon Resource Name (ARN) of the configured audience model that the resource policy will govern.")
		cleanroomsml_putConfiguredAudienceModelPolicyCmd.Flags().String("configured-audience-model-policy", "", "The IAM resource policy.")
		cleanroomsml_putConfiguredAudienceModelPolicyCmd.Flags().String("policy-existence-condition", "", "Use this to prevent unexpected concurrent modification of the policy.")
		cleanroomsml_putConfiguredAudienceModelPolicyCmd.Flags().String("previous-policy-hash", "", "A cryptographic hash of the contents of the policy used to prevent unexpected concurrent modification of the policy.")
		cleanroomsml_putConfiguredAudienceModelPolicyCmd.MarkFlagRequired("configured-audience-model-arn")
		cleanroomsml_putConfiguredAudienceModelPolicyCmd.MarkFlagRequired("configured-audience-model-policy")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_putConfiguredAudienceModelPolicyCmd)
}
