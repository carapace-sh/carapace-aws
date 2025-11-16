package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getConfiguredAudienceModelPolicyCmd = &cobra.Command{
	Use:   "get-configured-audience-model-policy",
	Short: "Returns information about a configured audience model policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getConfiguredAudienceModelPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_getConfiguredAudienceModelPolicyCmd).Standalone()

		cleanroomsml_getConfiguredAudienceModelPolicyCmd.Flags().String("configured-audience-model-arn", "", "The Amazon Resource Name (ARN) of the configured audience model that you are interested in.")
		cleanroomsml_getConfiguredAudienceModelPolicyCmd.MarkFlagRequired("configured-audience-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_getConfiguredAudienceModelPolicyCmd)
}
