package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeCapacityReservationCmd = &cobra.Command{
	Use:   "describe-capacity-reservation",
	Short: "Describes the capacity reservation status for the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeCapacityReservationCmd).Standalone()

	elbv2_describeCapacityReservationCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
	elbv2_describeCapacityReservationCmd.MarkFlagRequired("load-balancer-arn")
	elbv2Cmd.AddCommand(elbv2_describeCapacityReservationCmd)
}
