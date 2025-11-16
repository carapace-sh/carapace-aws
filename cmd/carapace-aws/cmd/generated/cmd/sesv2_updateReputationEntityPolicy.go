package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_updateReputationEntityPolicyCmd = &cobra.Command{
	Use:   "update-reputation-entity-policy",
	Short: "Update the reputation management policy for a reputation entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_updateReputationEntityPolicyCmd).Standalone()

	sesv2_updateReputationEntityPolicyCmd.Flags().String("reputation-entity-policy", "", "The Amazon Resource Name (ARN) of the reputation management policy to apply to this entity.")
	sesv2_updateReputationEntityPolicyCmd.Flags().String("reputation-entity-reference", "", "The unique identifier for the reputation entity.")
	sesv2_updateReputationEntityPolicyCmd.Flags().String("reputation-entity-type", "", "The type of reputation entity.")
	sesv2_updateReputationEntityPolicyCmd.MarkFlagRequired("reputation-entity-policy")
	sesv2_updateReputationEntityPolicyCmd.MarkFlagRequired("reputation-entity-reference")
	sesv2_updateReputationEntityPolicyCmd.MarkFlagRequired("reputation-entity-type")
	sesv2Cmd.AddCommand(sesv2_updateReputationEntityPolicyCmd)
}
