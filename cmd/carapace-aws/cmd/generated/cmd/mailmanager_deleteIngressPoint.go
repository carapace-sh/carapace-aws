package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_deleteIngressPointCmd = &cobra.Command{
	Use:   "delete-ingress-point",
	Short: "Delete an ingress endpoint resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_deleteIngressPointCmd).Standalone()

	mailmanager_deleteIngressPointCmd.Flags().String("ingress-point-id", "", "The identifier of the ingress endpoint resource that you want to delete.")
	mailmanager_deleteIngressPointCmd.MarkFlagRequired("ingress-point-id")
	mailmanagerCmd.AddCommand(mailmanager_deleteIngressPointCmd)
}
