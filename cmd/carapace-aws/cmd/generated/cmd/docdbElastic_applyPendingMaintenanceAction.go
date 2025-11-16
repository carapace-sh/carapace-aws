package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_applyPendingMaintenanceActionCmd = &cobra.Command{
	Use:   "apply-pending-maintenance-action",
	Short: "The type of pending maintenance action to be applied to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_applyPendingMaintenanceActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElastic_applyPendingMaintenanceActionCmd).Standalone()

		docdbElastic_applyPendingMaintenanceActionCmd.Flags().String("apply-action", "", "The pending maintenance action to apply to the resource.")
		docdbElastic_applyPendingMaintenanceActionCmd.Flags().String("apply-on", "", "A specific date to apply the pending maintenance action.")
		docdbElastic_applyPendingMaintenanceActionCmd.Flags().String("opt-in-type", "", "A value that specifies the type of opt-in request, or undoes an opt-in request.")
		docdbElastic_applyPendingMaintenanceActionCmd.Flags().String("resource-arn", "", "The Amazon DocumentDB Amazon Resource Name (ARN) of the resource to which the pending maintenance action applies.")
		docdbElastic_applyPendingMaintenanceActionCmd.MarkFlagRequired("apply-action")
		docdbElastic_applyPendingMaintenanceActionCmd.MarkFlagRequired("opt-in-type")
		docdbElastic_applyPendingMaintenanceActionCmd.MarkFlagRequired("resource-arn")
	})
	docdbElasticCmd.AddCommand(docdbElastic_applyPendingMaintenanceActionCmd)
}
