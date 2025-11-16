package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listSensorStatisticsCmd = &cobra.Command{
	Use:   "list-sensor-statistics",
	Short: "Lists statistics about the data collected for each of the sensors that have been successfully ingested in the particular dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listSensorStatisticsCmd).Standalone()

	lookoutequipment_listSensorStatisticsCmd.Flags().String("dataset-name", "", "The name of the dataset associated with the list of Sensor Statistics.")
	lookoutequipment_listSensorStatisticsCmd.Flags().String("ingestion-job-id", "", "The ingestion job id associated with the list of Sensor Statistics.")
	lookoutequipment_listSensorStatisticsCmd.Flags().String("max-results", "", "Specifies the maximum number of sensors for which to retrieve statistics.")
	lookoutequipment_listSensorStatisticsCmd.Flags().String("next-token", "", "An opaque pagination token indicating where to continue the listing of sensor statistics.")
	lookoutequipment_listSensorStatisticsCmd.MarkFlagRequired("dataset-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_listSensorStatisticsCmd)
}
