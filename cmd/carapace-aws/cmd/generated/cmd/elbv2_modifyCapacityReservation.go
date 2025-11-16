package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_modifyCapacityReservationCmd = &cobra.Command{
	Use:   "modify-capacity-reservation",
	Short: "Modifies the capacity reservation of the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_modifyCapacityReservationCmd).Standalone()

	elbv2_modifyCapacityReservationCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
	elbv2_modifyCapacityReservationCmd.Flags().String("minimum-load-balancer-capacity", "", "The minimum load balancer capacity reserved.")
	elbv2_modifyCapacityReservationCmd.Flags().String("reset-capacity-reservation", "", "Resets the capacity reservation.")
	elbv2_modifyCapacityReservationCmd.MarkFlagRequired("load-balancer-arn")
	elbv2Cmd.AddCommand(elbv2_modifyCapacityReservationCmd)
}
