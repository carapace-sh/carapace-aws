package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd = &cobra.Command{
	Use:   "notify-update-provisioned-product-engine-workflow-result",
	Short: "Notifies the result of the update engine execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd).Standalone()

		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.Flags().String("failure-reason", "", "The reason why the update engine execution failed.")
		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.Flags().String("idempotency-token", "", "The idempotency token that identifies the update engine execution.")
		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.Flags().String("outputs", "", "The output of the update engine execution.")
		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.Flags().String("record-id", "", "The identifier of the record.")
		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.Flags().String("status", "", "The status of the update engine execution.")
		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.Flags().String("workflow-token", "", "The encrypted contents of the update engine execution payload that Service Catalog sends after the Terraform product update workflow starts.")
		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.MarkFlagRequired("idempotency-token")
		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.MarkFlagRequired("record-id")
		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.MarkFlagRequired("status")
		servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd.MarkFlagRequired("workflow-token")
	})
	servicecatalogCmd.AddCommand(servicecatalog_notifyUpdateProvisionedProductEngineWorkflowResultCmd)
}
