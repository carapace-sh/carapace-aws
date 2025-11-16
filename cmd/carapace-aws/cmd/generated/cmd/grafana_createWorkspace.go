package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_createWorkspaceCmd = &cobra.Command{
	Use:   "create-workspace",
	Short: "Creates a *workspace*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_createWorkspaceCmd).Standalone()

	grafana_createWorkspaceCmd.Flags().String("account-access-type", "", "Specifies whether the workspace can access Amazon Web Services resources in this Amazon Web Services account only, or whether it can also access Amazon Web Services resources in other accounts in the same organization.")
	grafana_createWorkspaceCmd.Flags().String("authentication-providers", "", "Specifies whether this workspace uses SAML 2.0, IAM Identity Center, or both to authenticate users for using the Grafana console within a workspace.")
	grafana_createWorkspaceCmd.Flags().String("client-token", "", "A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.")
	grafana_createWorkspaceCmd.Flags().String("configuration", "", "The configuration string for the workspace that you create.")
	grafana_createWorkspaceCmd.Flags().String("grafana-version", "", "Specifies the version of Grafana to support in the new workspace.")
	grafana_createWorkspaceCmd.Flags().String("network-access-control", "", "Configuration for network access to your workspace.")
	grafana_createWorkspaceCmd.Flags().String("organization-role-name", "", "The name of an IAM role that already exists to use with Organizations to access Amazon Web Services data sources and notification channels in other accounts in an organization.")
	grafana_createWorkspaceCmd.Flags().String("permission-type", "", "When creating a workspace through the Amazon Web Services API, CLI or Amazon Web Services CloudFormation, you must manage IAM roles and provision the permissions that the workspace needs to use Amazon Web Services data sources and notification channels.")
	grafana_createWorkspaceCmd.Flags().String("stack-set-name", "", "The name of the CloudFormation stack set to use to generate IAM roles to be used for this workspace.")
	grafana_createWorkspaceCmd.Flags().String("tags", "", "The list of tags associated with the workspace.")
	grafana_createWorkspaceCmd.Flags().String("vpc-configuration", "", "The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.")
	grafana_createWorkspaceCmd.Flags().String("workspace-data-sources", "", "This parameter is for internal use only, and should not be used.")
	grafana_createWorkspaceCmd.Flags().String("workspace-description", "", "A description for the workspace.")
	grafana_createWorkspaceCmd.Flags().String("workspace-name", "", "The name for the workspace.")
	grafana_createWorkspaceCmd.Flags().String("workspace-notification-destinations", "", "Specify the Amazon Web Services notification channels that you plan to use in this workspace.")
	grafana_createWorkspaceCmd.Flags().String("workspace-organizational-units", "", "Specifies the organizational units that this workspace is allowed to use data sources from, if this workspace is in an account that is part of an organization.")
	grafana_createWorkspaceCmd.Flags().String("workspace-role-arn", "", "Specified the IAM role that grants permissions to the Amazon Web Services resources that the workspace will view data from, including both data sources and notification channels.")
	grafana_createWorkspaceCmd.MarkFlagRequired("account-access-type")
	grafana_createWorkspaceCmd.MarkFlagRequired("authentication-providers")
	grafana_createWorkspaceCmd.MarkFlagRequired("permission-type")
	grafanaCmd.AddCommand(grafana_createWorkspaceCmd)
}
