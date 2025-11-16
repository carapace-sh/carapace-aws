package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getReservationCmd = &cobra.Command{
	Use:   "get-reservation",
	Short: "Gets an Amazon Redshift Serverless reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getReservationCmd).Standalone()

	redshiftServerless_getReservationCmd.Flags().String("reservation-id", "", "The ID of the reservation to retrieve.")
	redshiftServerless_getReservationCmd.MarkFlagRequired("reservation-id")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getReservationCmd)
}
