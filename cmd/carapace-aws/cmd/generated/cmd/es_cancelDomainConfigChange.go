package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_cancelDomainConfigChangeCmd = &cobra.Command{
	Use:   "cancel-domain-config-change",
	Short: "Cancels a pending configuration change on an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_cancelDomainConfigChangeCmd).Standalone()

	es_cancelDomainConfigChangeCmd.Flags().String("domain-name", "", "Name of the OpenSearch Service domain configuration request to cancel.")
	es_cancelDomainConfigChangeCmd.Flags().String("dry-run", "", "When set to **True**, returns the list of change IDs and properties that will be cancelled without actually cancelling the change.")
	es_cancelDomainConfigChangeCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_cancelDomainConfigChangeCmd)
}
