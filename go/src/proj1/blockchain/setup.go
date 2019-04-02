package blockchain

import (
	"fmt"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"  // for ca Functioning  //for client level
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt" //Resource management of client
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry" //Packages to handle errors related to Hyperledger fabric
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp" // ca functioning for server level 
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config" // Blockchain core structure
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"  //Importing Fabric SDK
	"github.com/pkg/errors" // general exceptions
)

// FabricSetup implementation
type FabricSetup struct {
	ConfigFile      string
	OrgID           string
	OrdererID	string
	ChannelID       string
	ChainCodeID     string
	initialized     bool
	ChannelConfig   string
	ChaincodeGoPath string
	ChaincodePath   string
	OrgAdmin        string
	OrgName         string
	UserName        string
	admin           *resmgmt.Client  // pointer data types are mentioned in packages imported
	sdk             *fabsdk.FabricSDK // its a file pointer both 
}

// Initialize reads the configuration file and sets up the client, chain and event hub
func (setup *FabricSetup) Initialize() error { // Initialize() inbuilt func in FabricSDK
// Overriding the funtion Initialize()
	// Add parameters for the initialization
	if setup.initialized {
		return errors.New("sdk already initialized")
	}

	// Initialize the SDK with the configuration file
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile)) // sdk aur error me creating a new instance and configuration is taken from the specifed file
	if err != nil {
		return errors.WithMessage(err, "failed to create SDK")
	}
	setup.sdk = sdk
	fmt.Println("SDK created")

	// The resource management client is responsible for managing channels (create/update channel)
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName)) // Calling the function with user admin and orgsnisation name
//Creating Channel
	if err != nil {
		return errors.WithMessage(err, "failed to load Admin identity")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create channel management client from Admin identity")
	}
	setup.admin = resMgmtClient
	fmt.Println("Resource management client created")

	// The MSP client allow us to retrieve user information from their identity, like its signing identity which we will need to save the channel
	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(setup.OrgName))
// Context()- Additional Parameters
	if err != nil {
		return errors.WithMessage(err, "failed to create MSP client")
	}
	adminIdentity, err := mspClient.GetSigningIdentity(setup.OrgAdmin)
	if err != nil {
		return errors.WithMessage(err, "failed to get admin signing identity")
	}
	req := resmgmt.SaveChannelRequest{ChannelID: setup.ChannelID, ChannelConfigPath: setup.ChannelConfig, SigningIdentities: []msp.SigningIdentity{adminIdentity}}
	txID, err := setup.admin.SaveChannel(req, resmgmt.WithOrdererEndpoint(setup.OrdererID))
	if err != nil || txID.TransactionID == "" {
		return errors.WithMessage(err, "failed to save channel")
	}
	fmt.Println("Channel created")

	// Make admin user join the previously created channel
	if err = setup.admin.JoinChannel(setup.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(setup.OrdererID)); err != nil {
		return errors.WithMessage(err, "failed to make admin join channel")
	}
	fmt.Println("Channel joined")

	fmt.Println("Initialization Successful")
	setup.initialized = true
	return nil
}

func (setup *FabricSetup) CloseSDK() {
	setup.sdk.Close()
}

