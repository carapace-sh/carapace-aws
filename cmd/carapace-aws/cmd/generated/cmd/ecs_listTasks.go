package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listTasksCmd = &cobra.Command{
	Use:   "list-tasks",
	Short: "Returns a list of tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listTasksCmd).Standalone()

	ecs_listTasksCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster to use when filtering the `ListTasks` results.")
	ecs_listTasksCmd.Flags().String("container-instance", "", "The container instance ID or full ARN of the container instance to use when filtering the `ListTasks` results.")
	ecs_listTasksCmd.Flags().String("desired-status", "", "The task desired status to use when filtering the `ListTasks` results.")
	ecs_listTasksCmd.Flags().String("family", "", "The name of the task definition family to use when filtering the `ListTasks` results.")
	ecs_listTasksCmd.Flags().String("launch-type", "", "The launch type to use when filtering the `ListTasks` results.")
	ecs_listTasksCmd.Flags().String("max-results", "", "The maximum number of task results that `ListTasks` returned in paginated output.")
	ecs_listTasksCmd.Flags().String("next-token", "", "The `nextToken` value returned from a `ListTasks` request indicating that more results are available to fulfill the request and further calls will be needed.")
	ecs_listTasksCmd.Flags().String("service-name", "", "The name of the service to use when filtering the `ListTasks` results.")
	ecs_listTasksCmd.Flags().String("started-by", "", "The `startedBy` value to filter the task results with.")
	ecsCmd.AddCommand(ecs_listTasksCmd)
}
