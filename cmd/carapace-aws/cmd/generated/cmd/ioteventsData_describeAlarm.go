package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_describeAlarmCmd = &cobra.Command{
	Use:   "describe-alarm",
	Short: "Retrieves information about an alarm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_describeAlarmCmd).Standalone()

	ioteventsData_describeAlarmCmd.Flags().String("alarm-model-name", "", "The name of the alarm model.")
	ioteventsData_describeAlarmCmd.Flags().String("key-value", "", "The value of the key used as a filter to select only the alarms associated with the [key](https://docs.aws.amazon.com/iotevents/latest/apireference/API_CreateAlarmModel.html#iotevents-CreateAlarmModel-request-key).")
	ioteventsData_describeAlarmCmd.MarkFlagRequired("alarm-model-name")
	ioteventsDataCmd.AddCommand(ioteventsData_describeAlarmCmd)
}
