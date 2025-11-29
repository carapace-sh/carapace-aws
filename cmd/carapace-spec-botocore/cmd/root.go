package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-aws/cmd/carapace-spec-botocore/cmd/customizations"
	"github.com/carapace-sh/carapace-spec/pkg/command"
	"github.com/neurosnap/sentences/english"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var rootCmd = &cobra.Command{
	Use:   "carapace-spec-botocore",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		command, err := parse(args[0])
		if err != nil {
			return err
		}
		customizations.CustomizeCommand("", command)

		if cmd.Flag("no-doc").Changed {
			stripDoc(command)
		}

		if cmd.Flag("stdout").Changed {
			m, err := yaml.Marshal(command)
			if err != nil {
				return err
			}
			fmt.Println("# yaml-language-server: $schema=https://carapace.sh/schemas/command.json")
			fmt.Println(string(m))
			return nil
		}

		dir, err := os.MkdirTemp("", "carapace-spec-botocore-*")
		if err != nil {
			return err
		}

		for _, subCommand := range command.Commands {
			m, err := yaml.Marshal(subCommand)
			if err != nil {
				return err
			}
			m = append([]byte("# yaml-language-server: $schema=https://carapace.sh/schemas/command.json\n"), m...)
			path := path.Join(dir, fmt.Sprintf("aws.%s.yaml", subCommand.Name))
			println(path)
			if err := os.WriteFile(path, m, os.ModePerm); err != nil {
				return err
			}
		}

		command.Commands = nil
		m, err := yaml.Marshal(command)
		if err != nil {
			return err
		}
		path := path.Join(dir, "aws.yaml")
		println(path)
		if err := os.WriteFile(path, m, os.ModePerm); err != nil {
			return err
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().Bool("no-doc", false, "strip documentation")
	rootCmd.Flags().Bool("stdout", false, "print to stdout")
}

func CamelCaseToDash(s string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	s = re.ReplaceAllString(s, "${1}-${2}")

	re = regexp.MustCompile("([A-Z])([A-Z])([a-z0-9])")
	s = re.ReplaceAllString(s, "${1}-${2}${3}")

	s = strings.ToLower(s)

	quickfix := map[string]string{
		"application-i-ds":                                   "application-ids",
		"asset-sh-a256":                                      "asset-sha256",
		"associated-license-asset-ruleset-ar-ns ":            "associated-license-asset-ruleset-arns",
		"associated-license-asset-ruleset-ar-ns":             "associated-license-asset-ruleset-arns",
		"associate-whats-app-business-account":               "associate-whatsapp-business-account",
		"autonomous-data-storage-size-in-t-bs":               "autonomous-data-storage-size-in-tbs",
		"callback-ur-ls":                                     "callback-urls",
		"create-cachedi-scsi-volume":                         "describe-stored-iscsi-volumes",
		"create-environment-e-c2":                            "create-environment-ec2",
		"create-storedi-scsi-volume":                         "create-stored-iscsi-volume",
		"create-whats-app-message-template":                  "create-whatsapp-message-template",
		"create-whats-app-message-template-from-library":     "create-whatsapp-message-template-from-library",
		"create-whats-app-message-template-media":            "create-whatsapp-message-template-media",
		"data-storage-size-in-t-bs":                          "data-storage-size-in-tbs",
		"db-node-storage-size-in-g-bs":                       "db-node-storage-size-in-gbs",
		"delete-whats-app-message-media":                     "delete-whatsapp-message-media",
		"delete-whats-app-message-template":                  "delete-whatsapp-message-template",
		"describe-ac-ls":                                     "describe-acls",
		"describe-cachedi-scsi-volumes":                      "describe-cached-iscsi-volumes",
		"describe-e-c2-instance-limits":                      "describe-ec2-instance-limits",
		"describe-ic-d10-cm-inference-job":                   "describe-icd10-cm-inference-job",
		"describe-storedi-scsi-volumes":                      "create-cached-iscsi-volume",
		"disassociate-whats-app-business-account":            "disassociate-whatsapp-business-account",
		"e-c2-family-filter":                                 "ec2-family-filter",
		"e-c2-inbound-permissions":                           "ec2-inbound-permissions",
		"e-c2-instance-id":                                   "ec2-instance-id",
		"e-c2-instance-type":                                 "ec2-instance-type",
		"e-c2-security-group-id":                             "ec2-security-group-id",
		"e-c2-security-group-name":                           "ec2-security-group-name",
		"e-c2-security-group-owner-id":                       "ec2-security-group-owner-id",
		"edn-s0-client-subnet-ip":                            "edns0-client-subnet-ip",
		"edn-s0-client-subnet-mask":                          "edns0-client-subnet-mask",
		"export-e-c2-instance-recommendations":               "export-ec2-instance-recommendations",
		"get-e-c2-instance-recommendations":                  "get-ec2-instance-recommendations",
		"get-e-c2-recommendation-projected-metrics":          "get-ec2-recommendation-projected-metrics",
		"get-linked-whats-app-business-account":              "get-linked-whatsapp-business-account",
		"get-linked-whats-app-business-account-phone-number": "get-linked-whatsapp-business-account-phone-number",
		"get-whats-app-message-media":                        "get-whatsapp-message-media",
		"get-whats-app-message-template":                     "get-whatsapp-message-template",
		"infer-ic-d10-cm":                                    "infer-icd10-cm",
		"launch-configuration-template-i-ds":                 "launch-configuration-template-ids",
		"list-hi-ts-for-qualification-type":                  "list-hits-for-qualification-type",
		"list-hi-ts":                                         "list-hits",
		"list-ic-d10-cm-inference-jobs":                      "list-icd10-cm-inference-jobs",
		"list-linked-whats-app-business-accounts":            "list-linked-whatsapp-business-accounts",
		"list-reviewable-hi-ts":                              "list-reviewable-hits",
		"list-web-ac-ls":                                     "list-web-acls",
		"list-whats-app-message-templates":                   "list-whatsapp-message-templates",
		"list-whats-app-template-library":                    "list-whatsapp-template-library",
		"logout-ur-ls":                                       "logout-urls",
		"memory-per-oracle-compute-unit-in-g-bs":             "memory-per-oracle-compute-unit-in-gbs",
		"memory-size-in-g-bs":                                "memory-size-in-gbs",
		"no-retrieve-a-zs":                                   "no-retrieve-azs",
		"notification-ar-ns":                                 "notification-arns",
		"open-id-connect-provider-ar-ns":                     "open-id-connect-provider-arns",
		"pool-ar-ns":                                         "pool-arns",
		"post-whats-app-message-media":                       "post-whatsapp-message-media",
		"provider-ar-ns":                                     "provider-arns",
		"put-whats-app-business-account-event-destinations":  "put-whatsapp-business-account-event-destinations",
		"recovery-instance-i-ds":                             "recovery-instance-ids",
		"replication-configuration-template-i-ds":            "replication-configuration-template-ids",
		"replication-servers-security-groups-i-ds":           "replication-servers-security-groups-ids",
		"retrieve-a-zs":                                      "retrieve-azs",
		"saml-provider-ar-ns":                                "saml-provider-arns",
		"send-whats-app-message":                             "send-whatsapp-message",
		"source-server-i-d":                                  "source-server-ids",
		"start-ic-d10-cm-inference-job":                      "start-icd10-cm-inference-job",
		"stop-ic-d10-cm-inference-job":                       "stop-icd10-cm-inference-job",
		"swap-environment-cnam-es":                           "swap-environment-cnames",
		"tape-ar-ns":                                         "tape-arns",
		"target-group-ar-ns":                                 "target-group-arns",
		"update-whats-app-message-template":                  "update-whatsapp-message-template",
		"volume-ar-ns":                                       "volume-arns",
		"vtl-device-ar-ns":                                   "vtl-device-arns",
	}
	if replacement, ok := quickfix[s]; ok {
		return replacement
	}
	return s
}

type Service struct {
	Version  string `json:"version,omitempty"`
	Metadata struct {
		APIVersion          string   `json:"apiVersion,omitempty"`
		EndpointPrefix      string   `json:"endpointPrefix,omitempty"`
		Protocol            string   `json:"protocol,omitempty"`
		Protocols           []string `json:"protocols,omitempty"`
		ServiceAbbreviation string   `json:"serviceAbbreviation,omitempty"`
		ServiceFullName     string   `json:"serviceFullName,omitempty"`
		ServiceID           string   `json:"serviceId,omitempty"`
		SignatureVersion    string   `json:"signatureVersion,omitempty"`
		UID                 string   `json:"uid,omitempty"`
		XMLNamespace        string   `json:"xmlNamespace,omitempty"`
		Auth                []string `json:"auth,omitempty"`
	} `json:"metadata"`
	Operations    map[string]Operation
	Shapes        map[string]Shape `json:"shapes"`
	Documentation string           `json:"documentation,omitempty"`
}

type Shape struct {
	Type     string            `json:"type,omitempty"`
	Enum     []string          `json:"enum,omitempty"`
	Member   Member            `json:"member"`
	Members  map[string]Member `json:"members"`
	Required []string          `json:"required,omitempty"`
	Payload  string            `json:"payload,omitempty"`
}

type Member struct {
	Shape         string `json:"shape,omitempty"`
	Documentation string `json:"documentation,omitempty"`
	LocationName  string `json:"locationName,omitempty"`
}

type Operation struct {
	Name string `json:"name"`
	HTTP struct {
		Method     string `json:"method"`
		RequestURI string `json:"requestUri"`
	} `json:"http"`
	Input struct {
		Shape string `json:"shape"`
	} `json:"input"`
	Output struct {
		Shape string `json:"shape"`
	} `json:"output"`
	Documentation string `json:"documentation"`
}

func parse(path string) (*command.Command, error) {
	services, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	cmd := &command.Command{
		Name:        "aws",
		Description: "The AWS Command Line Interface is a unified tool to manage your AWS services.",
	}

	for _, serviceDir := range services {
		if !serviceDir.IsDir() {
			continue
		}
		path := filepath.Join(path, serviceDir.Name())
		versions, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		cmd.Commands = append(cmd.Commands,
			parseService(serviceDir.Name(), filepath.Join(path, versions[len(versions)-1].Name())),
		)
	}
	return cmd, nil
}

type Paginator struct {
	LimitKey *string `json:"limit_key,omitempty"`
}

func parsePaginators(folder string) (map[string]Paginator, error) {
	content, err := os.ReadFile(filepath.Join(folder, "paginators-1.json"))
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]Paginator), nil
		}
		return nil, err
	}

	var paginators struct {
		Pagination map[string]Paginator `json:"pagination"`
	}
	if err := json.Unmarshal(content, &paginators); err != nil {
		return nil, err
	}
	return paginators.Pagination, nil
}

func parseService(name, folder string) command.Command {
	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err.Error())
	}

	content, err := os.ReadFile(filepath.Join(folder, "service-2.json"))
	if err != nil {
		panic(err.Error())
	}

	var service Service
	if err := json.Unmarshal(content, &service); err != nil {
		panic(err.Error())
	}

	paginators, err := parsePaginators(folder)
	if err != nil {
		panic(err.Error())
	}

	cmd := command.Command{Name: name}
	cmd.Description = service.Metadata.ServiceFullName
	cmd.Documentation.Command, _ = htmltomarkdown.ConvertString(service.Documentation)

	for name, operation := range service.Operations {
		if customizations.Removal(cmd.Name, CamelCaseToDash(name)) {
			continue
		}

		opdoc, _ := htmltomarkdown.ConvertString(operation.Documentation)
		subCmd := command.Command{
			Name: CamelCaseToDash(name),
		}
		if tokens := tokenizer.Tokenize(opdoc); len(tokens) > 0 {
			subCmd.Description = strings.Split(tokens[0].Text, "\n")[0]
		}
		subCmd.Documentation.Command = opdoc
		subCmd.Documentation.Flag = make(map[string]string)

		// custom flags added by awscli
		switch {
		case isStreaming(service, operation):
			subCmd.Completion.Positional = [][]string{{"$files"}}
		default:
			addCustomFlags(&subCmd, paginators, name)
		}

		if shape, ok := service.Shapes[operation.Input.Shape]; ok {
			switch shape.Type {
			case "structure":
				// TODO shape.Member
				for name, member := range shape.Members {
					required := slices.Contains(shape.Required, name)
					memberdoc, _ := htmltomarkdown.ConvertString(member.Documentation)

					flagName := renameFlag(cmd.Name, subCmd.Name, CamelCaseToDash(name))
					subCmd.Documentation.Flag[flagName] = memberdoc
					if tokens := tokenizer.Tokenize(memberdoc); len(tokens) > 0 {
						memberdoc = strings.Split(tokens[0].Text, "\n")[0]
					}

					boolFlag := isBoolean(member.Shape)
					if !boolFlag {
						if shape, ok := service.Shapes[member.Shape]; ok && isBoolean(shape.Type) {
							boolFlag = true
						}
					}

					subCmd.AddFlag(command.Flag{
						Longhand:    flagName,
						Description: memberdoc,
						Value:       !boolFlag,
						Required:    required,
					})
					if boolFlag {
						negatedFlagName := "no-" + flagName
						negatedFlagName = renameFlag(cmd.Name, subCmd.Name, negatedFlagName)

						subCmd.AddFlag(command.Flag{
							Longhand:    negatedFlagName,
							Description: memberdoc,
							Required:    required,
						})
					}

					if memberShape, ok := service.Shapes[member.Shape]; ok &&
						strings.ToLower(memberShape.Type) == "string" && len(memberShape.Enum) != 0 {
						if subCmd.Completion.Flag == nil {
							subCmd.Completion.Flag = map[string][]string{}
						}
						subCmd.Completion.Flag[CamelCaseToDash(name)] = memberShape.Enum
					}
				}
			default:
				// TODO others
			}
		}
		customizations.CustomizeCommand(fmt.Sprintf("%s.%s", cmd.Name, subCmd.Name), &subCmd)
		cmd.Commands = append(cmd.Commands, subCmd)
	}

	waiters, err := parseWaiters(folder)
	if err != nil {
		panic(err.Error())
	}
	if len(waiters) > 0 {
		waitCmd := command.Command{
			Name:        "wait",
			Description: "Wait  until  a particular condition is satisfied.",
		}

		for name, waiter := range waiters {
			operationName := CamelCaseToDash(waiter.Operation)
			for _, command := range cmd.Commands {
				if command.Name == operationName {
					command.Name = CamelCaseToDash(name)
					waitCmd.Commands = append(waitCmd.Commands, command)
					break
				}
			}
			// TODO sort in carapace-spec
			slices.SortFunc(waitCmd.Commands, func(a, b command.Command) int { return strings.Compare(a.Name, b.Name) })
		}
		cmd.Commands = append(cmd.Commands, waitCmd)
	}

	customizations.CustomizeCommand(cmd.Name, &cmd)

	// TODO sort in carapace-spec
	slices.SortFunc(cmd.Commands, func(a, b command.Command) int { return strings.Compare(a.Name, b.Name) })
	return cmd
}

func isStreaming(service Service, operation Operation) bool {
	if shape, ok := service.Shapes[operation.Output.Shape]; ok {
		if shape.Payload != "" {
			if payloadShape, ok := shape.Members[shape.Payload]; ok {
				if strings.ToLower(payloadShape.Shape) == "blob" {
					return true
				}
				if payloadShape, ok := service.Shapes[payloadShape.Shape]; ok {
					if strings.ToLower(payloadShape.Type) == "blob" {
						return true
					}
				}
			}
		}
	}
	return false
}

func addCustomFlags(subCmd *command.Command, paginators map[string]Paginator, name string) {
	subCmd.AddFlag(command.Flag{Longhand: "cli-input-json", Description: "Read arguments from the JSON string provided.", Value: true})
	subCmd.AddFlag(command.Flag{Longhand: "cli-input-yaml", Description: "Read arguments from the YAML string provided.", Value: true})
	subCmd.AddFlag(command.Flag{Longhand: "generate-cli-skeleton", Description: "Prints a JSON skeleton to standard output without sending an API request."})

	subCmd.Documentation.Flag["cli-input-json"] = `Reads arguments from the JSON string provided.
The JSON string follows the  format  provided  by --generate-cli-skeleton.
If other arguments are provided on the command line,  those  values  will override the JSON-provided values.
It is not possible to pass arbitrary binary values using a JSON-provided value as the string will be taken literally.
This may  not  be  specified  along with --cli-input-yaml.`
	subCmd.Documentation.Flag["cli-input-yaml"] = `Reads arguments from the JSON string provided.
The JSON string follows the  format  provided  by --generate-cli-skeleton.
If other arguments are provided on the command line,  those  values  will override the JSON-provided values.
It is not possible to pass arbitrary binary values using a JSON-provided value as the string will be taken literally.
This may  not  be  specified  along with --cli-input-yaml.`
	subCmd.Documentation.Flag["generate-cli-skeleton"] = `If other arguments are provided on the command line,  those  values  will override the JSON-provided values. It is not possible to pass arbitrary binary values using a JSON-provided value as the string will be taken literally.
This may  not  be  specified  along with --cli-input-yaml.`

	if paginator, ok := paginators[name]; ok {
		subCmd.AddFlag(command.Flag{Longhand: "max-items", Description: "The  total number of items to return in the command's output.", Value: true})
		subCmd.AddFlag(command.Flag{Longhand: "starting-token", Description: "A token to specify where to start paginating.", Value: true})

		subCmd.Documentation.Flag["max-items"] = `The total number of items to return in the command's output.
If the total number of items available is more than the value specified, a NextToken is provided in the command's output.
To resume pagination, provide the NextToken value in the starting-token argument of a subsequent  command.
Do not use the NextToken response element directly outside of the AWS CLI.

For usage examples, see Pagination in the AWS Command Line Interface User Guide.`
		subCmd.Documentation.Flag["starting-token"] = `A token to specify where to start paginating.
This is the NextToken from a previously truncated response.

For usage examples, see Pagination in the AWS Command Line Interface User Guide.`

		if paginator.LimitKey != nil {
			subCmd.AddFlag(command.Flag{Longhand: "page-size", Description: "The size of each page to get in the AWS service call.", Value: true})

			subCmd.Documentation.Flag["page-size"] = `The size of each page to get in the AWS service call.
This does not affect the number of items returned in the command's output.
Setting a smaller page size results in more calls to the AWS service, retrieving fewer items in each call.
This can help prevent the AWS service calls from timing out.

For usage examples, see Pagination in the AWS Command Line Interface User Guide.`
		}
	}
}

func isBoolean(s string) bool {
	return (strings.HasSuffix(strings.ToLower(s), "boolean") && !strings.HasPrefix(strings.ToLower(s), "map")) || strings.HasSuffix(s, "AttributeBooleanValue")
}

type Waiter struct {
	Description string `json:"description"`
	Operation   string `json:"operation"`
}

func parseWaiters(folder string) (map[string]Waiter, error) {
	content, err := os.ReadFile(filepath.Join(folder, "waiters-2.json"))
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]Waiter), nil
		}
		return nil, err
	}

	var waiters struct {
		Waiters map[string]Waiter `json:"waiters"`
	}
	if err := json.Unmarshal(content, &waiters); err != nil {
		return nil, err
	}
	return waiters.Waiters, nil
}

func stripDoc(command *command.Command) {
	command.Documentation.Command = ""
	command.Documentation.Flag = nil
	for index := range command.Commands {
		stripDoc(&command.Commands[index])
	}
}
