package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_updateIngressPointCmd = &cobra.Command{
	Use:   "update-ingress-point",
	Short: "Update attributes of a provisioned ingress endpoint resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_updateIngressPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_updateIngressPointCmd).Standalone()

		mailmanager_updateIngressPointCmd.Flags().String("ingress-point-configuration", "", "If you choose an Authenticated ingress endpoint, you must configure either an SMTP password or a secret ARN.")
		mailmanager_updateIngressPointCmd.Flags().String("ingress-point-id", "", "The identifier for the ingress endpoint you want to update.")
		mailmanager_updateIngressPointCmd.Flags().String("ingress-point-name", "", "A user friendly name for the ingress endpoint resource.")
		mailmanager_updateIngressPointCmd.Flags().String("rule-set-id", "", "The identifier of an existing rule set that you attach to an ingress endpoint resource.")
		mailmanager_updateIngressPointCmd.Flags().String("status-to-update", "", "The update status of an ingress endpoint.")
		mailmanager_updateIngressPointCmd.Flags().String("traffic-policy-id", "", "The identifier of an existing traffic policy that you attach to an ingress endpoint resource.")
		mailmanager_updateIngressPointCmd.MarkFlagRequired("ingress-point-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_updateIngressPointCmd)
}
