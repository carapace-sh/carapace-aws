package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_acceptReservedInstancesExchangeQuoteCmd = &cobra.Command{
	Use:   "accept-reserved-instances-exchange-quote",
	Short: "Accepts the Convertible Reserved Instance exchange quote described in the [GetReservedInstancesExchangeQuote]() call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_acceptReservedInstancesExchangeQuoteCmd).Standalone()

	ec2_acceptReservedInstancesExchangeQuoteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptReservedInstancesExchangeQuoteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptReservedInstancesExchangeQuoteCmd.Flags().String("reserved-instance-ids", "", "The IDs of the Convertible Reserved Instances to exchange for another Convertible Reserved Instance of the same or higher value.")
	ec2_acceptReservedInstancesExchangeQuoteCmd.Flags().String("target-configurations", "", "The configuration of the target Convertible Reserved Instance to exchange for your current Convertible Reserved Instances.")
	ec2_acceptReservedInstancesExchangeQuoteCmd.Flag("no-dry-run").Hidden = true
	ec2_acceptReservedInstancesExchangeQuoteCmd.MarkFlagRequired("reserved-instance-ids")
	ec2Cmd.AddCommand(ec2_acceptReservedInstancesExchangeQuoteCmd)
}
