package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_describeAlarmsCmd = &cobra.Command{
	Use:   "describe-alarms",
	Short: "Retrieves the specified alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_describeAlarmsCmd).Standalone()

	cloudwatch_describeAlarmsCmd.Flags().String("action-prefix", "", "Use this parameter to filter the results of the operation to only those alarms that use a certain alarm action.")
	cloudwatch_describeAlarmsCmd.Flags().String("alarm-name-prefix", "", "An alarm name prefix.")
	cloudwatch_describeAlarmsCmd.Flags().String("alarm-names", "", "The names of the alarms to retrieve information about.")
	cloudwatch_describeAlarmsCmd.Flags().String("alarm-types", "", "Use this parameter to specify whether you want the operation to return metric alarms or composite alarms.")
	cloudwatch_describeAlarmsCmd.Flags().String("children-of-alarm-name", "", "If you use this parameter and specify the name of a composite alarm, the operation returns information about the \"children\" alarms of the alarm you specify.")
	cloudwatch_describeAlarmsCmd.Flags().String("max-records", "", "The maximum number of alarm descriptions to retrieve.")
	cloudwatch_describeAlarmsCmd.Flags().String("next-token", "", "The token returned by a previous call to indicate that there is more data available.")
	cloudwatch_describeAlarmsCmd.Flags().String("parents-of-alarm-name", "", "If you use this parameter and specify the name of a metric or composite alarm, the operation returns information about the \"parent\" alarms of the alarm you specify.")
	cloudwatch_describeAlarmsCmd.Flags().String("state-value", "", "Specify this parameter to receive information only about alarms that are currently in the state that you specify.")
	cloudwatchCmd.AddCommand(cloudwatch_describeAlarmsCmd)
}
