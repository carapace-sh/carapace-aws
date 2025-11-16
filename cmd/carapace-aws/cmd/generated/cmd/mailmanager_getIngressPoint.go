package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getIngressPointCmd = &cobra.Command{
	Use:   "get-ingress-point",
	Short: "Fetch ingress endpoint resource attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getIngressPointCmd).Standalone()

	mailmanager_getIngressPointCmd.Flags().String("ingress-point-id", "", "The identifier of an ingress endpoint.")
	mailmanager_getIngressPointCmd.MarkFlagRequired("ingress-point-id")
	mailmanagerCmd.AddCommand(mailmanager_getIngressPointCmd)
}
