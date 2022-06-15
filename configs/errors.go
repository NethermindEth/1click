package configs

// All the strings that are needed for error logging.
const (
	InstallingDependenciesError       = "something went wrong while installing dependencies. %s"
	ShowingInstructionsError          = "something went wrong while showing the instructions for installing %s"
	ScriptPathError                   = "failed to get path for instructions file. Error: %s"
	ReadingInstructionError           = "failed to read instructions from file %s"
	OSNotSupportedError               = "installation not supported for %s"
	ProvideClientsError               = "please provide both execution client and consensus client"
	IncorrectClientError              = "incorrect %s client name \"%s\". Please provide correct client name. Use 'clients' command to see the list of supported clients"
	NoClientsFoundError               = "no %s clients found. Please check your configuration file"
	CreatingFileError                 = "failed to create file %s. Error: %s"
	OpeningFileError                  = "failed to open file %s. Error: %s"
	ClosingFileError                  = "failed to close file %s"
	GeneratingScriptsError            = "generating docker-compose files for execution client %s, consensus client %s and validator client %s failed. Error: %s"
	ClientNotSupportedError           = "client %s is not supported. Please use 'clients' command to see the list of supported clients"
	GetRawTemplatesError              = "failed to get raw templates for %s"
	LoadingTemplatesError             = "error loading templates: %s"
	PrintingFileError                 = "something went wrong printing file %s. Error: %s"
	RunningCMDError                   = "failed to execute '%s' command. Error: %s"
	DependencyNotInstalledError       = "dependency %s is not installed on host machine"
	DependenciesMissingError          = "required dependencies are missing. Please install them and try again. Dependencies can be installed using '1click cli' command"
	DockerEngineOffError              = "it seems docker engine is not running. Please start it and try again. Error: %s"
	DockerComposeScriptNotFoundError  = "docker-compose script not found at %s. Please run '1click cli' command to generate it or check the script path. By default, it should be located at %s"
	ScriptIsNotRunningError           = "services of docker-compose script provided are not running. Error: %s"
	GettingLogsError                  = "failed to get logs for services %s. Error: %s"
	DockerComposePsReturnedEmptyError = "'docker-compose ps --services' returned empty string"
	PromptFailedError                 = "prompt failed %s"
	GeneratingKeystoreError           = "failed to generate keystore folder. Error: %s"
	ResizingPtyError                  = "error resizing pty: %s"
	InvalidVolumePathError            = "invalid path provided: %s. If you intended to pass a host directory, use absolute path"
	ZipError                          = "all lists must have the same size"
	CommandError                      = "command '%s' throws error: %s"
	RunClientsFlagAmbiguousError      = "provided values %v for run-clients flag are ambiguous. If you want to use the option 'all', then do not insert aditional options"
	InstallNotSupported               = "install support for %s is not available for %s. Please install it and try again"
	DistroInfoError                   = "failed to get linux distribution info. Error: %s"
	MonitoringToolInitError           = "monitoring tool initialization failed. Error: %s"
	TrackSyncError                    = "endpoint %s returned an error. Error: %v"
	RunClientsError                   = "provided values %v for run-clients are incorrect. Posible correct values are %v"
	EmptyClientMapError               = "is not possible to select a random element from an empty collection"
	NoSupportedClientsError           = "collection of clients given for random choice doesn't have any supported client. Check the target network (flag --network). Use 'clients' command to see the list of supported clients for every supported network"
	NetworkValidationFailedError      = "'network' flag validation failed. Error: %v"
	UnknownNetworkError               = "unknown network \"%s\". Please provide correct network name. Use 'networks' command to see the list of supported networks"
)
