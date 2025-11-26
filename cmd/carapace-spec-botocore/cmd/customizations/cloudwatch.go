package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["cloudwatch.put-metric-data"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "unit",
			Description: "The unit of metric.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "metric-name",
			Description: "The name of the metric.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "dimensions",
			Description: "The dimensions of the metric.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "value",
			Description: "The value for the metric.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "timestamp",
			Description: "The time stamp used for the metric.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "statistic-values",
			Description: "A set of statistical values describing the metric.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "storage-resolution",
			Description: "The storage resolution of the metric.",
			Value:       true,
		})
		return nil
	}
}
