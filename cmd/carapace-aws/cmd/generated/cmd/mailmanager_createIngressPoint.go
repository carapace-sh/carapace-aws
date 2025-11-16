package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_createIngressPointCmd = &cobra.Command{
	Use:   "create-ingress-point",
	Short: "Provision a new ingress endpoint resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_createIngressPointCmd).Standalone()

	mailmanager_createIngressPointCmd.Flags().String("client-token", "", "A unique token that Amazon SES uses to recognize subsequent retries of the same request.")
	mailmanager_createIngressPointCmd.Flags().String("ingress-point-configuration", "", "If you choose an Authenticated ingress endpoint, you must configure either an SMTP password or a secret ARN.")
	mailmanager_createIngressPointCmd.Flags().String("ingress-point-name", "", "A user friendly name for an ingress endpoint resource.")
	mailmanager_createIngressPointCmd.Flags().String("network-configuration", "", "Specifies the network configuration for the ingress point.")
	mailmanager_createIngressPointCmd.Flags().String("rule-set-id", "", "The identifier of an existing rule set that you attach to an ingress endpoint resource.")
	mailmanager_createIngressPointCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for the resource.")
	mailmanager_createIngressPointCmd.Flags().String("traffic-policy-id", "", "The identifier of an existing traffic policy that you attach to an ingress endpoint resource.")
	mailmanager_createIngressPointCmd.Flags().String("type", "", "The type of the ingress endpoint to create.")
	mailmanager_createIngressPointCmd.MarkFlagRequired("ingress-point-name")
	mailmanager_createIngressPointCmd.MarkFlagRequired("rule-set-id")
	mailmanager_createIngressPointCmd.MarkFlagRequired("traffic-policy-id")
	mailmanager_createIngressPointCmd.MarkFlagRequired("type")
	mailmanagerCmd.AddCommand(mailmanager_createIngressPointCmd)
}
