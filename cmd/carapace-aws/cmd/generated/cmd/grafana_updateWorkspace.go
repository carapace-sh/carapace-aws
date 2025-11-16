package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_updateWorkspaceCmd = &cobra.Command{
	Use:   "update-workspace",
	Short: "Modifies an existing Amazon Managed Grafana workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_updateWorkspaceCmd).Standalone()

	grafana_updateWorkspaceCmd.Flags().String("account-access-type", "", "Specifies whether the workspace can access Amazon Web Services resources in this Amazon Web Services account only, or whether it can also access Amazon Web Services resources in other accounts in the same organization.")
	grafana_updateWorkspaceCmd.Flags().String("network-access-control", "", "The configuration settings for network access to your workspace.")
	grafana_updateWorkspaceCmd.Flags().Bool("no-remove-network-access-configuration", false, "Whether to remove the network access configuration from the workspace.")
	grafana_updateWorkspaceCmd.Flags().Bool("no-remove-vpc-configuration", false, "Whether to remove the VPC configuration from the workspace.")
	grafana_updateWorkspaceCmd.Flags().String("organization-role-name", "", "The name of an IAM role that already exists to use to access resources through Organizations.")
	grafana_updateWorkspaceCmd.Flags().String("permission-type", "", "Use this parameter if you want to change a workspace from `SERVICE_MANAGED` to `CUSTOMER_MANAGED`.")
	grafana_updateWorkspaceCmd.Flags().Bool("remove-network-access-configuration", false, "Whether to remove the network access configuration from the workspace.")
	grafana_updateWorkspaceCmd.Flags().Bool("remove-vpc-configuration", false, "Whether to remove the VPC configuration from the workspace.")
	grafana_updateWorkspaceCmd.Flags().String("stack-set-name", "", "The name of the CloudFormation stack set to use to generate IAM roles to be used for this workspace.")
	grafana_updateWorkspaceCmd.Flags().String("vpc-configuration", "", "The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.")
	grafana_updateWorkspaceCmd.Flags().String("workspace-data-sources", "", "This parameter is for internal use only, and should not be used.")
	grafana_updateWorkspaceCmd.Flags().String("workspace-description", "", "A description for the workspace.")
	grafana_updateWorkspaceCmd.Flags().String("workspace-id", "", "The ID of the workspace to update.")
	grafana_updateWorkspaceCmd.Flags().String("workspace-name", "", "A new name for the workspace to update.")
	grafana_updateWorkspaceCmd.Flags().String("workspace-notification-destinations", "", "Specify the Amazon Web Services notification channels that you plan to use in this workspace.")
	grafana_updateWorkspaceCmd.Flags().String("workspace-organizational-units", "", "Specifies the organizational units that this workspace is allowed to use data sources from, if this workspace is in an account that is part of an organization.")
	grafana_updateWorkspaceCmd.Flags().String("workspace-role-arn", "", "Specifies an IAM role that grants permissions to Amazon Web Services resources that the workspace accesses, such as data sources and notification channels.")
	grafana_updateWorkspaceCmd.Flag("no-remove-network-access-configuration").Hidden = true
	grafana_updateWorkspaceCmd.Flag("no-remove-vpc-configuration").Hidden = true
	grafana_updateWorkspaceCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_updateWorkspaceCmd)
}
