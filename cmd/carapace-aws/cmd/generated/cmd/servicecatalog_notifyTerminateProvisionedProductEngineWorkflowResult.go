package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd = &cobra.Command{
	Use:   "notify-terminate-provisioned-product-engine-workflow-result",
	Short: "Notifies the result of the terminate engine execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd).Standalone()

		servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd.Flags().String("failure-reason", "", "The reason why the terminate engine execution failed.")
		servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd.Flags().String("idempotency-token", "", "The idempotency token that identifies the terminate engine execution.")
		servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd.Flags().String("record-id", "", "The identifier of the record.")
		servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd.Flags().String("status", "", "The status of the terminate engine execution.")
		servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd.Flags().String("workflow-token", "", "The encrypted contents of the terminate engine execution payload that Service Catalog sends after the Terraform product terminate workflow starts.")
		servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd.MarkFlagRequired("idempotency-token")
		servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd.MarkFlagRequired("record-id")
		servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd.MarkFlagRequired("status")
		servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd.MarkFlagRequired("workflow-token")
	})
	servicecatalogCmd.AddCommand(servicecatalog_notifyTerminateProvisionedProductEngineWorkflowResultCmd)
}
