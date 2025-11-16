package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deleteHealthCheckCmd = &cobra.Command{
	Use:   "delete-health-check",
	Short: "Deletes a health check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deleteHealthCheckCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_deleteHealthCheckCmd).Standalone()

		route53_deleteHealthCheckCmd.Flags().String("health-check-id", "", "The ID of the health check that you want to delete.")
		route53_deleteHealthCheckCmd.MarkFlagRequired("health-check-id")
	})
	route53Cmd.AddCommand(route53_deleteHealthCheckCmd)
}
