package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteMitigationActionCmd = &cobra.Command{
	Use:   "delete-mitigation-action",
	Short: "Deletes a defined mitigation action from your Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteMitigationActionCmd).Standalone()

	iot_deleteMitigationActionCmd.Flags().String("action-name", "", "The name of the mitigation action that you want to delete.")
	iot_deleteMitigationActionCmd.MarkFlagRequired("action-name")
	iotCmd.AddCommand(iot_deleteMitigationActionCmd)
}
