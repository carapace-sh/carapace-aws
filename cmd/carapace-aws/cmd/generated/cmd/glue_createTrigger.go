package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createTriggerCmd = &cobra.Command{
	Use:   "create-trigger",
	Short: "Creates a new trigger.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createTriggerCmd).Standalone()

	glue_createTriggerCmd.Flags().String("actions", "", "The actions initiated by this trigger when it fires.")
	glue_createTriggerCmd.Flags().String("description", "", "A description of the new trigger.")
	glue_createTriggerCmd.Flags().String("event-batching-condition", "", "Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.")
	glue_createTriggerCmd.Flags().String("name", "", "The name of the trigger.")
	glue_createTriggerCmd.Flags().String("predicate", "", "A predicate to specify when the new trigger should fire.")
	glue_createTriggerCmd.Flags().String("schedule", "", "A `cron` expression used to specify the schedule (see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ?")
	glue_createTriggerCmd.Flags().String("start-on-creation", "", "Set to `true` to start `SCHEDULED` and `CONDITIONAL` triggers when created.")
	glue_createTriggerCmd.Flags().String("tags", "", "The tags to use with this trigger.")
	glue_createTriggerCmd.Flags().String("type", "", "The type of the new trigger.")
	glue_createTriggerCmd.Flags().String("workflow-name", "", "The name of the workflow associated with the trigger.")
	glue_createTriggerCmd.MarkFlagRequired("actions")
	glue_createTriggerCmd.MarkFlagRequired("name")
	glue_createTriggerCmd.MarkFlagRequired("type")
	glueCmd.AddCommand(glue_createTriggerCmd)
}
