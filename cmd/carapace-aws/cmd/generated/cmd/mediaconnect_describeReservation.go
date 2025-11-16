package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_describeReservationCmd = &cobra.Command{
	Use:   "describe-reservation",
	Short: "Displays the details of a reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_describeReservationCmd).Standalone()

	mediaconnect_describeReservationCmd.Flags().String("reservation-arn", "", "The Amazon Resource Name (ARN) of the offering.")
	mediaconnect_describeReservationCmd.MarkFlagRequired("reservation-arn")
	mediaconnectCmd.AddCommand(mediaconnect_describeReservationCmd)
}
