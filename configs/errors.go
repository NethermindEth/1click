/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package configs

// All the strings that are needed for error logging.
const (
	InstallingDependenciesError       = "something went wrong while installing dependencies. %s"
	ShowingInstructionsError          = "something went wrong while showing the instructions for installing %s"
	ScriptPathError                   = "failed to get path for instructions file. Error: %v"
	ReadingInstructionError           = "failed to read instructions from file %s"
	OSNotSupportedError               = "installation not supported for %s"
	ProvideClientsError               = "please provide both execution client and consensus client"
	IncorrectClientError              = "incorrect %s client name \"%s\". Please provide correct client name. Use 'clients' command to see the list of supported clients"
	NoClientsFoundError               = "no %s clients found. Please check your configuration file"
	CreatingFileError                 = "failed to create file %s. Error: %v"
	OpeningFileError                  = "failed to open file %s. Error: %v"
	ClosingFileError                  = "failed to close file %s"
	GeneratingScriptsError            = "generating docker-compose files for execution client %s, consensus client %s and validator client %s failed. Error: %s"
	ClientNotSupportedError           = "client %s is not supported. Please use 'clients' command to see the list of supported clients"
	GetRawTemplatesError              = "failed to get raw templates for %s"
	LoadingTemplatesError             = "error loading templates: %s"
	PrintingFileError                 = "something went wrong printing file %s. Error: %v"
	RunningCMDError                   = "failed to execute '%s' command. Error: %v"
	DependencyNotInstalledError       = "dependency %s is not installed on host machine"
	DependenciesMissingError          = "required dependencies are missing. Please install them and try again. Dependencies can be installed using 'sedge cli' command"
	DockerEngineOffError              = "it seems docker engine is not running. Please start it and try again. Error: %v"
	DockerComposeOffError             = "it seems docker compose plugin is not installed. Please install it and try again. Error: %v"
	DockerComposeScriptNotFoundError  = "docker-compose script not found at %s. Please run 'sedge cli' command to generate it or check the script path. By default, it should be located at %s"
	ScriptIsNotRunningError           = "services of docker-compose script provided are not running. Error: %v"
	GettingLogsError                  = "failed to get logs for services %s. Error: %v"
	DockerComposePsReturnedEmptyError = "'docker compose ps --services' returned empty string"
	PromptFailedError                 = "prompt failed %v"
	GeneratingKeystoreError           = "failed to generate keystore folder. Error: %v"
	ResizingPtyError                  = "error resizing pty: %v"
	InvalidVolumePathError            = "invalid path provided: %s. If you intended to pass a host directory, use absolute path"
	ZipError                          = "all lists must have the same size"
	CommandError                      = "command '%s' throws error: %v"
	RunClientsFlagAmbiguousError      = "provided values %v for run-clients flag are ambiguous. If you want to use the option 'all', then do not insert aditional options"
	InstallNotSupported               = "install support for %s is not available for %s. Please install it and try again"
	DistroInfoError                   = "failed to get linux distribution info. Error: %v"
	MonitoringToolInitError           = "monitoring tool initialization failed. Error: %v"
	TrackSyncError                    = "endpoint %s returned an error. Error: %v"
	RunClientsError                   = "provided values %v for run-clients are incorrect. Posible correct values are %v"
	EmptyClientMapError               = "is not possible to select a random element from an empty collection"
	NoSupportedClientsError           = "collection of clients given for random choice doesn't have any supported client. Check the target network (flag --network). Use 'clients' command to see the list of supported clients for every supported network"
	NetworkValidationFailedError      = "'network' flag validation failed. Error: %v"
	UnknownNetworkError               = "unknown network \"%s\". Please provide correct network name. Use 'networks' command to see the list of supported networks"
	GenerateJWTSecretError            = "JWT secret generation failed. Error: %v"
	GetPWDError                       = "something failed trying to get current working directory. Error: %v"
	InvalidFeeRecipientError          = "provided Fee Recipient is not a valid Ethereum address"
	KeystorePasswordError             = "keystore password must have more than 8 characters"
	KeystorePasswordRetryError        = "the two entered keystore passwords do not match"
	CreatingKeystorePasswordError     = "failed to create keystore password file. Error: %v"
	GetContainerIPError               = "failed to get docker internal IP address of docker-compose service %s. Error: %v"
	TekuDatadirError                  = "failed to create teku %s datadir with 777 permissions. Error :%v"
	UnableToTrackSyncError            = "it seems both execution and consensus containers are down. It was not possible to get internal IP address for the monitoring tool to be able to track sync progress. Check the logs using 'sedge logs' to find the problem"
	NoOutputDockerInspectError        = "no output from docker inspect"
	IPNotFoundError                   = "could not find IP address"
	PortOccupationError               = "port occupation check failed. Error: %v"
	DefaultPortEmptyError             = "default %s can not be empty"
)
