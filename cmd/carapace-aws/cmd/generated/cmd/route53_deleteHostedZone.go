package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deleteHostedZoneCmd = &cobra.Command{
	Use:   "delete-hosted-zone",
	Short: "Deletes a hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deleteHostedZoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_deleteHostedZoneCmd).Standalone()

		route53_deleteHostedZoneCmd.Flags().String("id", "", "The ID of the hosted zone you want to delete.")
		route53_deleteHostedZoneCmd.MarkFlagRequired("id")
	})
	route53Cmd.AddCommand(route53_deleteHostedZoneCmd)
}
