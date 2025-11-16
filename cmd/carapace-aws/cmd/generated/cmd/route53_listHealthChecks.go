package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listHealthChecksCmd = &cobra.Command{
	Use:   "list-health-checks",
	Short: "Retrieve a list of the health checks that are associated with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listHealthChecksCmd).Standalone()

	route53_listHealthChecksCmd.Flags().String("marker", "", "If the value of `IsTruncated` in the previous response was `true`, you have more health checks.")
	route53_listHealthChecksCmd.Flags().String("max-items", "", "The maximum number of health checks that you want `ListHealthChecks` to return in response to the current request.")
	route53Cmd.AddCommand(route53_listHealthChecksCmd)
}
