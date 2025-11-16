package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getChangeCmd = &cobra.Command{
	Use:   "get-change",
	Short: "Returns the current status of a change batch request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getChangeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getChangeCmd).Standalone()

		route53_getChangeCmd.Flags().String("id", "", "The ID of the change batch request.")
		route53_getChangeCmd.MarkFlagRequired("id")
	})
	route53Cmd.AddCommand(route53_getChangeCmd)
}
