package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteActivationCmd = &cobra.Command{
	Use:   "delete-activation",
	Short: "Deletes an activation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteActivationCmd).Standalone()

	ssm_deleteActivationCmd.Flags().String("activation-id", "", "The ID of the activation that you want to delete.")
	ssm_deleteActivationCmd.MarkFlagRequired("activation-id")
	ssmCmd.AddCommand(ssm_deleteActivationCmd)
}
