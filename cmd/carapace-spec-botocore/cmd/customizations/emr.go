package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed emr/aws.emr.add-steps.yaml
var emr_addSteps []byte

//go:embed emr/aws.emr.create-cluster.yaml
var emr_createCluster []byte

//go:embed emr/aws.emr.create-default-roles.yaml
var emr_createDefaultRoles []byte

//go:embed emr/aws.emr.create-hbase-backup.yaml
var emr_createHbaseBackup []byte

//go:embed emr/aws.emr.disable-hbase-backups.yaml
var emr_disableHbaseBackups []byte

//go:embed emr/aws.emr.get.yaml
var emr_get []byte

//go:embed emr/aws.emr.install-applications.yaml
var emr_installApplications []byte

//go:embed emr/aws.emr.modify-cluster-attributes.yaml
var emr_modifyClusterAttributes []byte

//go:embed emr/aws.emr.put.yaml
var emr_put []byte

//go:embed emr/aws.emr.restore-from-hbase-backup.yaml
var emr_restoreHbaseBackup []byte

//go:embed emr/aws.emr.schedule-hbase-backup.yaml
var emr_scheduleHbaseBackup []byte

//go:embed emr/aws.emr.socks.yaml
var emr_socks []byte

//go:embed emr/aws.emr.ssh.yaml
var emr_ssh []byte

//go:embed emr/aws.emr.terminate-clusters.yaml
var emr_terminateClusters []byte

func init() {
	customizations["emr"] = func(cmd *command.Command) error {
		for _, spec := range [][]byte{
			emr_addSteps,
			emr_createCluster,
			emr_createDefaultRoles,
			emr_createHbaseBackup,
			emr_disableHbaseBackups,
			emr_get,
			emr_installApplications,
			emr_modifyClusterAttributes,
			emr_put,
			emr_restoreHbaseBackup,
			emr_scheduleHbaseBackup,
			emr_socks,
			emr_ssh,
			emr_terminateClusters,
		} {
			var specCommand command.Command
			if err := yaml.Unmarshal(spec, &specCommand); err != nil {
				return err
			}
			cmd.Commands = append(cmd.Commands, specCommand)
		}
		for i := range cmd.Commands {
			if cmd.Commands[i].Name != "wait" {
				continue
			}
			for j := range cmd.Commands[i].Commands {
				waitSubCmd := &cmd.Commands[i].Commands[j]
				if waitSubCmd.Name != "cluster-running" && waitSubCmd.Name != "cluster-terminated" {
					continue
				}
				waitSubCmd.AddFlag(command.Flag{
					Longhand:    "cli-input-json",
					Description: "Read arguments from the JSON string provided.",
					Value:       true,
				})
				waitSubCmd.AddFlag(command.Flag{
					Longhand:    "cli-input-yaml",
					Description: "Read arguments from the YAML string provided.",
					Value:       true,
				})
				waitSubCmd.AddFlag(command.Flag{
					Longhand:    "generate-cli-skeleton",
					Description: "Prints a JSON skeleton to standard output without sending an API request.",
				})
			}
		}
		return nil
	}

	customizations["emr.list-clusters"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "active",
			Description: "Shortcut options for –cluster-state",
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "failed",
			Description: "Shortcut options for –cluster-state",
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "terminated",
			Description: "Shortcut options for –cluster-state",
		})
		return nil
	}
	customizations["emr.describe-cluster"] = func(cmd *command.Command) error {
		delete(cmd.Flags, "--cli-input-json=")
		delete(cmd.Flags, "--cli-input-yaml=")
		delete(cmd.Flags, "--generate-cli-skeleton")
		return nil
	}
	customizations["emr.add-instance-groups"] = func(cmd *command.Command) error {
		delete(cmd.Flags, "--cli-input-json=")
		delete(cmd.Flags, "--cli-input-yaml=")
		delete(cmd.Flags, "--generate-cli-skeleton")
		return nil
	}
}
