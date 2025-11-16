package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteConfiguredAudienceModelPolicyCmd = &cobra.Command{
	Use:   "delete-configured-audience-model-policy",
	Short: "Deletes the specified configured audience model policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteConfiguredAudienceModelPolicyCmd).Standalone()

	cleanroomsml_deleteConfiguredAudienceModelPolicyCmd.Flags().String("configured-audience-model-arn", "", "The Amazon Resource Name (ARN) of the configured audience model policy that you want to delete.")
	cleanroomsml_deleteConfiguredAudienceModelPolicyCmd.MarkFlagRequired("configured-audience-model-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteConfiguredAudienceModelPolicyCmd)
}
