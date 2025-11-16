package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getReservedInstancesExchangeQuoteCmd = &cobra.Command{
	Use:   "get-reserved-instances-exchange-quote",
	Short: "Returns a quote and exchange information for exchanging one or more specified Convertible Reserved Instances for a new Convertible Reserved Instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getReservedInstancesExchangeQuoteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getReservedInstancesExchangeQuoteCmd).Standalone()

		ec2_getReservedInstancesExchangeQuoteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getReservedInstancesExchangeQuoteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getReservedInstancesExchangeQuoteCmd.Flags().String("reserved-instance-ids", "", "The IDs of the Convertible Reserved Instances to exchange.")
		ec2_getReservedInstancesExchangeQuoteCmd.Flags().String("target-configurations", "", "The configuration of the target Convertible Reserved Instance to exchange for your current Convertible Reserved Instances.")
		ec2_getReservedInstancesExchangeQuoteCmd.Flag("no-dry-run").Hidden = true
		ec2_getReservedInstancesExchangeQuoteCmd.MarkFlagRequired("reserved-instance-ids")
	})
	ec2Cmd.AddCommand(ec2_getReservedInstancesExchangeQuoteCmd)
}
