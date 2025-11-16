package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_notifyProvisionProductEngineWorkflowResultCmd = &cobra.Command{
	Use:   "notify-provision-product-engine-workflow-result",
	Short: "Notifies the result of the provisioning engine execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_notifyProvisionProductEngineWorkflowResultCmd).Standalone()

	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.Flags().String("failure-reason", "", "The reason why the provisioning engine execution failed.")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.Flags().String("idempotency-token", "", "The idempotency token that identifies the provisioning engine execution.")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.Flags().String("outputs", "", "The output of the provisioning engine execution.")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.Flags().String("record-id", "", "The identifier of the record.")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.Flags().String("resource-identifier", "", "The ID for the provisioned product resources that are part of a resource group.")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.Flags().String("status", "", "The status of the provisioning engine execution.")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.Flags().String("workflow-token", "", "The encrypted contents of the provisioning engine execution payload that Service Catalog sends after the Terraform product provisioning workflow starts.")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.MarkFlagRequired("idempotency-token")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.MarkFlagRequired("record-id")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.MarkFlagRequired("status")
	servicecatalog_notifyProvisionProductEngineWorkflowResultCmd.MarkFlagRequired("workflow-token")
	servicecatalogCmd.AddCommand(servicecatalog_notifyProvisionProductEngineWorkflowResultCmd)
}
