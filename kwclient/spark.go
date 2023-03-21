package kwclient

// 3013525718211495503
// 3013515008937035048
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type Graph struct {
	Prototype struct {
		Id                 string      `json:"id"`
		Type               string      `json:"type"`
		Name               string      `json:"name"`
		Icon               string      `json:"icon"`
		UserId             string      `json:"userId"`
		ProjectId          string      `json:"projectId"`
		ChannelTypeId      interface{} `json:"channelTypeId"`
		EditingBy          string      `json:"editingBy"`
		PrototypeDeletedAt interface{} `json:"prototypeDeletedAt"`
		CreatedAt          time.Time   `json:"createdAt"`
		UpdatedAt          time.Time   `json:"updatedAt"`
	} `json:"prototype"`
	Type struct {
		Id                  string `json:"id"`
		Icon                string `json:"icon"`
		Name                string `json:"name"`
		Type                string `json:"type"`
		StartNodeInstanceId string `json:"startNodeInstanceId"`
		Nodes               []struct {
			Field1 struct {
				Id          string `json:"id"`
				Name        string `json:"name"`
				Type        string `json:"type"`
				X           int    `json:"x"`
				Y           int    `json:"y"`
				Width       int    `json:"width"`
				Height      int    `json:"height"`
				PrototypeId string `json:"prototypeId"`
				DataSource  struct {
					DataReferenceType string        `json:"dataReferenceType"`
					Keys              []interface{} `json:"keys"`
				} `json:"dataSource"`
				DataDestination struct {
					DataReferenceType string        `json:"dataReferenceType"`
					Keys              []interface{} `json:"keys"`
				} `json:"dataDestination"`
				Args struct {
					Context   interface{} `json:"context"`
					Timestamp interface{} `json:"timestamp"`
				} `json:"args"`
				Deprecated bool `json:"deprecated"`
				Links      []struct {
					Id                 string  `json:"id"`
					NextNodeInstanceId string  `json:"nextNodeInstanceId"`
					Type               string  `json:"type"`
					Condition          *string `json:"condition"`
					Description        string  `json:"description"`
					Mode               string  `json:"mode"`
				} `json:"links"`
				Metrics []interface{} `json:"metrics"`
			} `json:"2852509980827845664"`
		} `json:"nodes"`
		Comments struct {
		} `json:"comments"`
		UserId     string `json:"userId"`
		ProjectId  string `json:"projectId"`
		Parameters struct {
		} `json:"parameters"`
		Tags []struct {
			Name      string `json:"name"`
			Color     string `json:"color"`
			IsDefault bool   `json:"isDefault"`
		} `json:"tags"`
	} `json:"type"`
}

type Script struct {
	Prototype struct {
		Id                 string      `json:"id"`
		Type               string      `json:"type"`
		Name               string      `json:"name"`
		Icon               string      `json:"icon"`
		UserId             string      `json:"userId"`
		ProjectId          string      `json:"projectId"`
		ChannelTypeId      interface{} `json:"channelTypeId"`
		EditingBy          string      `json:"editingBy"`
		PrototypeDeletedAt interface{} `json:"prototypeDeletedAt"`
		CreatedAt          time.Time   `json:"createdAt"`
		UpdatedAt          time.Time   `json:"updatedAt"`
	} `json:"prototype"`
	Type struct {
		Id           string      `json:"id"`
		Icon         string      `json:"icon"`
		Name         string      `json:"name"`
		Body         string      `json:"body"`
		ArgNames     []string    `json:"argNames"`
		ConnectionId interface{} `json:"connectionId"`
		ProjectId    string      `json:"projectId"`
		UserId       string      `json:"userId"`
		Type         string      `json:"type"`
	} `json:"type"`
}

type KWCode struct {
	Body string `json:"body"`
}

type KWJavaScriptPrototype struct {
	Prototype struct {
		Id                 string      `json:"id"`
		Type               string      `json:"type"`
		Name               string      `json:"name"`
		Icon               string      `json:"icon"`
		UserId             string      `json:"userId"`
		ProjectId          string      `json:"projectId"`
		ChannelTypeId      interface{} `json:"channelTypeId"`
		EditingBy          string      `json:"editingBy"`
		PrototypeDeletedAt interface{} `json:"prototypeDeletedAt"`
		CreatedAt          time.Time   `json:"createdAt"`
		UpdatedAt          time.Time   `json:"updatedAt"`
	} `json:"prototype"`
	Type struct {
		Icon         string        `json:"icon"`
		Body         string        `json:"body"`
		ArgNames     []interface{} `json:"argNames"`
		Id           string        `json:"id"`
		ProjectId    string        `json:"projectId"`
		UserId       string        `json:"userId"`
		Name         string        `json:"name"`
		ConnectionId interface{}   `json:"connectionId"`
		Type         string        `json:"type"`
	} `json:"type"`
}

type KWPrototype struct {
	Id                 string      `json:"id"`
	Type               string      `json:"type"`
	Name               string      `json:"name"`
	Icon               string      `json:"icon"`
	UserId             string      `json:"userId"`
	ProjectId          string      `json:"projectId"`
	ChannelTypeId      *string     `json:"channelTypeId"`
	EditingBy          *string     `json:"editingBy"`
	PrototypeDeletedAt interface{} `json:"prototypeDeletedAt"`
	CreatedAt          time.Time   `json:"createdAt"`
	UpdatedAt          time.Time   `json:"updatedAt"`
	Code               Script
	Graph              interface{}
	Tags               []struct {
		Name      string `json:"name"`
		Color     string `json:"color"`
		IsDefault bool   `json:"isDefault,omitempty"`
		Id        string `json:"id,omitempty"`
	} `json:"tags"`
	IconColor       string `json:"iconColor,omitempty"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	JourneyStepType string `json:"journeyStepType,omitempty"`
	Outcome         string `json:"outcome,omitempty"`
}

type KWProject struct {
	Project struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Schema      struct {
			JAM struct {
				LOG struct {
					Message struct {
						Node             string `json:"Node"`
						Output           string `json:"Output"`
						TransactionState string `json:"TransactionState"`
					} `json:"Message"`
					Category    string `json:"Category"`
					Severity    string `json:"Severity"`
					RequestId   string `json:"RequestId"`
					AmbientData struct {
						KWID  string `json:"KWID"`
						Error struct {
							Code    string `json:"Code"`
							Node    string `json:"Node"`
							Time    string `json:"Time"`
							Message string `json:"Message"`
						} `json:"Error"`
						Status       string `json:"Status"`
						GraphName    string `json:"GraphName"`
						HelpGuide    string `json:"HelpGuide"`
						ElapsedTime  string `json:"ElapsedTime"`
						ProjectName  string `json:"ProjectName"`
						PublishEvent struct {
							Topic     string `json:"Topic"`
							ObjectKey string `json:"ObjectKey"`
							Publisher string `json:"Publisher"`
						} `json:"PublishEvent"`
						TransactionId       string `json:"TransactionId"`
						RecordsChanged      string `json:"RecordsChanged"`
						TransactionEnd      string `json:"TransactionEnd"`
						EnvironmentName     string `json:"EnvironmentName"`
						TransactionType     string `json:"TransactionType"`
						TransactionStart    string `json:"TransactionStart"`
						ElapsedResponseTime string `json:"ElapsedResponseTime"`
					} `json:"AmbientData"`
					MessageType    string `json:"MessageType"`
					TenantClient   string `json:"TenantClient"`
					ConversationId string `json:"ConversationId"`
				} `json:"LOG"`
				Records struct {
					S3 struct {
						Bucket struct {
							Name string `json:"name"`
						} `json:"bucket"`
						Object struct {
							Key string `json:"key"`
						} `json:"object"`
					} `json:"s3"`
				} `json:"Records"`
			} `json:"JAM"`
			Auth struct {
				Host              string `json:"Host"`
				Accept            string `json:"Accept"`
				Connection        string `json:"Connection"`
				XAmzDate          string `json:"X-Amz-Date"`
				ContentType       string `json:"Content-Type"`
				Authorization     string `json:"Authorization"`
				CacheControl      string `json:"Cache-Control"`
				ContentLength     string `json:"Content-Length"`
				AcceptEncoding    string `json:"Accept-Encoding"`
				XAmzContentSha256 string `json:"X-Amz-Content-Sha256"`
			} `json:"auth"`
			Status  string `json:"status"`
			Ingress struct {
				Event              string `json:"Event"`
				Topic              string `json:"Topic"`
				UseCase            string `json:"UseCase"`
				EventType          string `json:"EventType"`
				SourceSystem       string `json:"SourceSystem"`
				LineOfBusiness     string `json:"LineOfBusiness"`
				PatientInformation struct {
					City              string `json:"City"`
					State             string `json:"State"`
					Gender            string `json:"Gender"`
					Zipcode           string `json:"Zipcode"`
					Address1          string `json:"Address1"`
					Address2          string `json:"Address2"`
					Language          string `json:"Language"`
					LastName          string `json:"LastName"`
					CellPhone         string `json:"CellPhone"`
					FirstName         string `json:"FirstName"`
					HomePhone         string `json:"HomePhone"`
					WorkPhone         string `json:"WorkPhone"`
					MiddleName        string `json:"MiddleName"`
					SuffixName        string `json:"SuffixName"`
					DateOfBirth       string `json:"DateOfBirth"`
					PhoneNumber       string `json:"PhoneNumber"`
					EmailAddress      string `json:"EmailAddress"`
					MemberNumber      string `json:"MemberNumber"`
					PhoneNumberType   string `json:"PhoneNumberType"`
					PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
				} `json:"PatientInformation"`
				ProviderInformation struct {
					SubmitterId    string `json:"SubmitterId"`
					DoctorLastName string `json:"DoctorLastName"`
					HOFriendlyName string `json:"HOFriendlyName"`
				} `json:"ProviderInformation"`
				AdditionalParameters     string `json:"AdditionalParameters"`
				CommunicationPreferences struct {
					SMSOptIn   string `json:"SMSOptIn"`
					EmailOptIn string `json:"EmailOptIn"`
					PhoneOptIn string `json:"PhoneOptIn"`
				} `json:"CommunicationPreferences"`
			} `json:"Ingress"`
			Channel string `json:"channel"`
			Encrypt struct {
				Key             string `json:"key"`
				PlainText       string `json:"plainText"`
				EncryptedOutput struct {
					IngressPayload     string `json:"IngressPayload"`
					PatientInformation struct {
						City              string `json:"City"`
						State             string `json:"State"`
						Gender            string `json:"Gender"`
						Zipcode           string `json:"Zipcode"`
						Address1          string `json:"Address1"`
						Address2          string `json:"Address2"`
						Language          string `json:"Language"`
						LastName          string `json:"LastName"`
						CellPhone         string `json:"CellPhone"`
						FirstName         string `json:"FirstName"`
						HomePhone         string `json:"HomePhone"`
						WorkPhone         string `json:"WorkPhone"`
						MiddleName        string `json:"MiddleName"`
						SuffixName        string `json:"SuffixName"`
						DateOfBirth       string `json:"DateOfBirth"`
						PhoneNumber       string `json:"PhoneNumber"`
						EmailAddress      string `json:"EmailAddress"`
						MemberNumber      string `json:"MemberNumber"`
						PhoneNumberType   string `json:"PhoneNumberType"`
						PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
					} `json:"PatientInformation"`
					ProviderInformation struct {
						SubmitterId    string `json:"SubmitterId"`
						DoctorLastName string `json:"DoctorLastName"`
						HOFriendlyName string `json:"HOFriendlyName"`
					} `json:"ProviderInformation"`
				} `json:"EncryptedOutput"`
			} `json:"encrypt"`
			Modules struct {
				ProjectHandlerListener struct {
					Input  string `json:"input"`
					Output string `json:"output"`
				} `json:"ProjectHandlerListener"`
			} `json:"modules"`
			Outreach struct {
				Event          string `json:"Event"`
				MemberNumber   string `json:"MemberNumber"`
				OutreachStatus string `json:"OutreachStatus"`
			} `json:"outreach"`
			Credential struct {
				Key    string `json:"key"`
				Uri    string `json:"uri"`
				Host   string `json:"host"`
				Region string `json:"region"`
				Secret string `json:"secret"`
			} `json:"Credential"`
			Encryption struct {
				StoreRequest struct {
					User        string `json:"user"`
					Project     string `json:"project"`
					Environment string `json:"environment"`
				} `json:"storeRequest"`
				StoreResponse struct {
					Body struct {
						Parameter struct {
							ARN   string `json:"ARN"`
							Name  string `json:"Name"`
							Type  string `json:"Type"`
							Value struct {
								Key    string `json:"key"`
								Vector string `json:"vector"`
							} `json:"Value"`
							Version          int    `json:"Version"`
							DataType         string `json:"DataType"`
							LastModifiedDate string `json:"LastModifiedDate"`
						} `json:"Parameter"`
						ResponseMetadata string `json:"ResponseMetadata"`
					} `json:"body"`
					Status struct {
						Code   int    `json:"code"`
						Reason string `json:"reason"`
					} `json:"status"`
					Headers string `json:"headers"`
				} `json:"storeResponse"`
				EncryptedOutput struct {
					IngressPayload     string `json:"IngressPayload"`
					PatientInformation struct {
						City              string `json:"City"`
						State             string `json:"State"`
						Gender            string `json:"Gender"`
						Zipcode           string `json:"Zipcode"`
						Address1          string `json:"Address1"`
						Address2          string `json:"Address2"`
						Language          string `json:"Language"`
						LastName          string `json:"LastName"`
						CellPhone         string `json:"CellPhone"`
						FirstName         string `json:"FirstName"`
						HomePhone         string `json:"HomePhone"`
						WorkPhone         string `json:"WorkPhone"`
						MiddleName        string `json:"MiddleName"`
						SuffixName        string `json:"SuffixName"`
						DateOfBirth       string `json:"DateOfBirth"`
						PhoneNumber       string `json:"PhoneNumber"`
						EmailAddress      string `json:"EmailAddress"`
						MemberNumber      string `json:"MemberNumber"`
						PhoneNumberType   string `json:"PhoneNumberType"`
						PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
					} `json:"PatientInformation"`
					ProviderInformation struct {
						SubmitterId    string `json:"SubmitterId"`
						DoctorLastName string `json:"DoctorLastName"`
						HOFriendlyName string `json:"HOFriendlyName"`
					} `json:"ProviderInformation"`
				} `json:"EncryptedOutput"`
			} `json:"Encryption"`
			CancelEvent struct {
				Event          string `json:"Event"`
				Topic          string `json:"Topic"`
				EventType      string `json:"EventType"`
				OutreachId     string `json:"outreachId"`
				MemberNumber   string `json:"MemberNumber"`
				SourceSystem   string `json:"SourceSystem"`
				LineOfBusiness string `json:"LineOfBusiness"`
			} `json:"cancelEvent"`
			Transaction struct {
				Scheduler struct {
					RunCount           string `json:"runCount"`
					GraphName          string `json:"graphName"`
					LastRunTs          string `json:"lastRunTs"`
					NextRunTs          string `json:"nextRunTs"`
					CurrentState       string `json:"currentState"`
					LastRunEndTs       string `json:"lastRunEndTs"`
					SecondLastRunTs    string `json:"secondLastRunTs"`
					SecondLastRunEndTs string `json:"secondLastRunEndTs"`
				} `json:"scheduler"`
				SchedulerTS string `json:"schedulerTS"`
			} `json:"transaction"`
			NoEngagement struct {
				Event                   string `json:"event"`
				Channel                 string `json:"channel"`
				UseCase                 string `json:"use_case"`
				MemberId                string `json:"member_id"`
				DateStamp               string `json:"date_stamp"`
				OutreachStatus          string `json:"outreach_status"`
				CommunicationIdentifier string `json:"communication_identifier"`
			} `json:"NoEngagement"`
			InboundEvent struct {
				Event           string `json:"Event"`
				Channel         string `json:"Channel"`
				UseCase         string `json:"UseCase"`
				Timestamp       string `json:"Timestamp"`
				MemberNumber    string `json:"MemberNumber"`
				EventProperties struct {
					Misc1 string `json:"Misc1"`
				} `json:"EventProperties"`
				CommunicationIdentifier string `json:"CommunicationIdentifier"`
			} `json:"inboundEvent"`
			SmsOutbound struct {
				Topic          string `json:"Topic"`
				Language       string `json:"Language"`
				PhoneNumber    string `json:"PhoneNumber"`
				SMSTemplate    string `json:"SMSTemplate"`
				MemberNumber   string `json:"MemberNumber"`
				SourceSystem   string `json:"SourceSystem"`
				DynamicContent struct {
					City              string `json:"City"`
					State             string `json:"State"`
					Gender            string `json:"Gender"`
					Zipcode           string `json:"Zipcode"`
					Address1          string `json:"Address1"`
					Address2          string `json:"Address2"`
					LastName          string `json:"LastName"`
					CellPhone         string `json:"CellPhone"`
					FirstName         string `json:"FirstName"`
					HomePhone         string `json:"HomePhone"`
					WorkPhone         string `json:"WorkPhone"`
					MiddleName        string `json:"MiddleName"`
					SuffixName        string `json:"SuffixName"`
					DateOfBirth       string `json:"DateOfBirth"`
					EmailAddress      string `json:"EmailAddress"`
					PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
				} `json:"DynamicContent"`
				PhoneNumberType         string `json:"PhoneNumberType"`
				AdditionalParameters    string `json:"AdditionalParameters"`
				CommunicationIdentifier string `json:"CommunicationIdentifier"`
			} `json:"sms_outbound"`
			JourneySteps struct {
				Steps string `json:"steps"`
			} `json:"journey-steps"`
			OutreachStatus string `json:"OutreachStatus"`
			DatabaseResult string `json:"databaseResult"`
			EmailOutbound  struct {
				Topic          string `json:"Topic"`
				Language       string `json:"Language"`
				EmailAddress   string `json:"EmailAddress"`
				MemberNumber   string `json:"MemberNumber"`
				SourceSystem   string `json:"SourceSystem"`
				EmailTemplate  string `json:"EmailTemplate"`
				DynamicContent struct {
					City              string `json:"City"`
					State             string `json:"State"`
					Gender            string `json:"Gender"`
					Zipcode           string `json:"Zipcode"`
					Address1          string `json:"Address1"`
					Address2          string `json:"Address2"`
					LastName          string `json:"LastName"`
					CellPhone         string `json:"CellPhone"`
					FirstName         string `json:"FirstName"`
					HomePhone         string `json:"HomePhone"`
					WorkPhone         string `json:"WorkPhone"`
					MiddleName        string `json:"MiddleName"`
					SuffixName        string `json:"SuffixName"`
					DateOfBirth       string `json:"DateOfBirth"`
					PhoneNumber       string `json:"PhoneNumber"`
					PhoneNumberType   string `json:"PhoneNumberType"`
					PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
				} `json:"DynamicContent"`
				AdditionalParameters    string `json:"AdditionalParameters"`
				CommunicationIdentifier string `json:"CommunicationIdentifier"`
			} `json:"email_outbound"`
			EncodedIngress   string `json:"encodedIngress"`
			OutreachRecord   string `json:"outreachRecord"`
			IngressUtilities struct {
				Channel                 string `json:"Channel"`
				DateStamp               string `json:"DateStamp"`
				OutreachId              string `json:"OutreachId"`
				CommunicationIdentifier string `json:"CommunicationIdentifier"`
			} `json:"IngressUtilities"`
			JavascriptResult   string `json:"JavascriptResult"`
			CredentialLiteral  string `json:"CredentialLiteral"`
			UniversalConnector struct {
				Payload struct {
				} `json:"payload "`
				UcSettings struct {
					Url     string `json:"url"`
					Headers struct {
						Host        string `json:"Host"`
						ContentType string `json:"Content-Type"`
					} `json:"headers"`
					Timeout string `json:"timeout"`
					AuthIds struct {
						JWT struct {
							Aud string `json:"aud"`
							Exp string `json:"exp"`
							Iss string `json:"iss"`
							Kid string `json:"kid"`
							Sub string `json:"sub"`
							Typ string `json:"typ"`
						} `json:"JWT"`
						Bearer        string `json:"Bearer"`
						UrlParameters string `json:"url_parameters"`
					} `json:"auth_ids"`
					AuthType      string `json:"auth_type"`
					ServiceType   string `json:"serviceType"`
					UrlParameters struct {
						UtmMedium string `json:"utm_medium"`
						UtmSource string `json:"utm_source"`
					} `json:"url_parameters"`
					RequestPayloadType  string `json:"requestPayloadType"`
					ResponsePayloadType string `json:"responsePayloadType"`
				} `json:"uc_settings"`
			} `json:"UniversalConnector"`
			ReturnFromGCACMValue string `json:"ReturnFromGCACMValue"`
			CancelEventUtilities struct {
				Channel                 string `json:"channel"`
				Payload                 string `json:"payload"`
				UseCase                 string `json:"use_case"`
				Datestamp               string `json:"datestamp"`
				CommunicationIdentifier string `json:"communication_identifier"`
			} `json:"cancelEventUtilities"`
			InboundEventUtilities struct {
				DbResult struct {
					UseCase string `json:"use_case"`
				} `json:"DbResult"`
			} `json:"inboundEventUtilities"`
			UniversalConnectorResponse struct {
				StatusCode struct {
					Code string `json:"code"`
				} `json:"statusCode"`
				ResponseBody struct {
					TokenizedURL string `json:"tokenizedURL"`
				} `json:"responseBody"`
				ResponseHeaders struct {
					XAmznRequestId             string `json:"x-amzn-RequestId"`
					XAmznRemappedAuthorization string `json:"x-amzn-Remapped-Authorization"`
				} `json:"responseHeaders"`
			} `json:"UniversalConnectorResponse"`
			KwServiceTemp20221104        string `json:"kw_service_temp_2022_11_04"`
			CredentialFromParameterStore struct {
				Status     string `json:"status"`
				Credential struct {
					Key struct {
						Crq       string `json:"crq"`
						Latest    string `json:"latest"`
						Secret    string `json:"secret"`
						Source    string `json:"source"`
						Version   string `json:"version"`
						AccessKey string `json:"accessKey"`
						CreatedBy string `json:"createdBy"`
						CreatedOn string `json:"createdOn"`
						ExpiresOn string `json:"expiresOn"`
					} `json:"key"`
					User   string `json:"user"`
					Source string `json:"source"`
				} `json:"credential"`
			} `json:"CredentialFromParameterStore"`
		} `json:"schema"`
		ProfileSchema struct {
			Attributes struct {
				OptOuts struct {
					SMS     string `json:"SMS"`
					Email   string `json:"Email"`
					UseCase string `json:"UseCase"`
				} `json:"opt_outs"`
			} `json:"Attributes"`
			Identifiers struct {
				KWID         string `json:"KWID"`
				MemberNumber string `json:"memberNumber"`
			} `json:"Identifiers"`
			Interactions string `json:"Interactions"`
			JourneySteps string `json:"JourneySteps"`
		} `json:"profileSchema"`
		DefaultEnvironmentId    string      `json:"defaultEnvironmentId"`
		ActiveEnvironmentId     string      `json:"activeEnvironmentId"`
		ProductionEnvironmentId string      `json:"productionEnvironmentId"`
		EmergencyContactListId  string      `json:"emergencyContactListId"`
		UserId                  string      `json:"userId"`
		OrganizationId          string      `json:"organizationId"`
		HasKDMNode              bool        `json:"hasKDMNode"`
		TemplateId              interface{} `json:"templateId"`
		RemoteHub               interface{} `json:"remoteHub"`
		TopLevelId              interface{} `json:"topLevelId"`
		CreatedAt               time.Time   `json:"createdAt"`
		UpdatedAt               time.Time   `json:"updatedAt"`
		Tags                    []struct {
			Id    string `json:"id"`
			Name  string `json:"name"`
			Color string `json:"color"`
		} `json:"tags"`
		Environments []struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			ProjectId string `json:"projectId"`
		} `json:"environments"`
		Favorites []interface{} `json:"favorites"`
		Metrics   struct {
		} `json:"metrics"`
		SplitTests struct {
		} `json:"splitTests"`
		PublicVariables struct {
		} `json:"publicVariables"`
		Versions    interface{} `json:"versions"`
		Permissions struct {
		} `json:"permissions"`
		Users  []interface{} `json:"users"`
		Groups []interface{} `json:"groups"`
	} `json:"project"`
	Prototypes []struct {
		Id                 string      `json:"id"`
		Type               string      `json:"type"`
		Name               string      `json:"name"`
		Icon               string      `json:"icon"`
		UserId             string      `json:"userId"`
		ProjectId          string      `json:"projectId"`
		ChannelTypeId      *string     `json:"channelTypeId"`
		EditingBy          *string     `json:"editingBy"`
		PrototypeDeletedAt interface{} `json:"prototypeDeletedAt"`
		CreatedAt          time.Time   `json:"createdAt"`
		UpdatedAt          time.Time   `json:"updatedAt"`
		Code               Script
		Graph              interface{}
		Tags               []struct {
			Name      string `json:"name"`
			Color     string `json:"color"`
			IsDefault bool   `json:"isDefault,omitempty"`
			Id        string `json:"id,omitempty"`
		} `json:"tags"`
		IconColor       string `json:"iconColor,omitempty"`
		BackgroundColor string `json:"backgroundColor,omitempty"`
		JourneyStepType string `json:"journeyStepType,omitempty"`
		Outcome         string `json:"outcome,omitempty"`
	} `json:"prototypes"`
	Types struct {
		Deleted bool   `json:"_deleted"`
		Id      string `json:"id"`
	} `json:"types"`
	Request string `json:"request"`
}

type KWConnections struct {
	Id                  string  `json:"id"`
	Type                string  `json:"type"`
	Name                string  `json:"name"`
	ProjectId           string  `json:"projectId"`
	Managed             bool    `json:"managed"`
	ManagedConnectionId *string `json:"managedConnectionId"`
	Connectors          []struct {
		Id                string      `json:"id"`
		DatabaseType      string      `json:"databaseType,omitempty"`
		Username          *string     `json:"username,omitempty"`
		Password          string      `json:"password,omitempty"`
		Hostname          *string     `json:"hostname,omitempty"`
		Port              *int        `json:"port,omitempty"`
		Database          *string     `json:"database,omitempty"`
		IsolationLevel    interface{} `json:"isolationLevel"`
		ConnectionId      string      `json:"connectionId"`
		EnvironmentId     string      `json:"environmentId"`
		ConnectionType    string      `json:"connection_type,omitempty"`
		Brokers           []string    `json:"brokers,omitempty"`
		Topics            []string    `json:"topics,omitempty"`
		ConnectionDetails struct {
			Algorithm             string `json:"algorithm"`
			Username              string `json:"username,omitempty"`
			Password              string `json:"password,omitempty"`
			ConnectorRequiredName string `json:"connector_required_name,omitempty"`
		} `json:"connection_details,omitempty"`
		CreatedAt   time.Time `json:"createdAt,omitempty"`
		UpdatedAt   time.Time `json:"updatedAt,omitempty"`
		Endpoint    string    `json:"endpoint,omitempty"`
		QueryParams struct {
		} `json:"queryParams,omitempty"`
		HeaderParams struct {
		} `json:"headerParams,omitempty"`
		MutualAuth          bool        `json:"mutualAuth,omitempty"`
		AwsSignatureEnabled bool        `json:"awsSignatureEnabled,omitempty"`
		SignatureService    string      `json:"signatureService,omitempty"`
		SignatureHost       string      `json:"signatureHost,omitempty"`
		SignatureRegion     interface{} `json:"signatureRegion"`
		AwsSecretAccessKey  string      `json:"awsSecretAccessKey,omitempty"`
		QueueType           string      `json:"queueType,omitempty"`
		ReadQueue           string      `json:"readQueue,omitempty"`
		WriteQueue          string      `json:"writeQueue,omitempty"`
		AwsAccessKeyId      *string     `json:"awsAccessKeyId,omitempty"`
		AwsRegion           *string     `json:"awsRegion,omitempty"`
		FileType            string      `json:"fileType,omitempty"`
		AwsBucket           *string     `json:"awsBucket,omitempty"`
	} `json:"connectors"`
	ManagedConnectionName string `json:"managedConnectionName,omitempty"`
}

type Fuse2 struct {
	Connections  []KWConnections
	Name         string
	JourneySteps []KWPrototype
	Graphs       []KWPrototype
	JavaScripts  []KWPrototype
}

type Connectors struct {
	Size  int `json:"size"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Data  []struct {
		Id    string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		StartItem struct {
			Links struct {
				Self string `json:"self"`
			} `json:"links"`
			Id       string `json:"id"`
			Position struct {
				X string `json:"x"`
				Y string `json:"y"`
			} `json:"position,omitempty"`
		} `json:"startItem"`
		EndItem struct {
			Links struct {
				Self string `json:"self"`
			} `json:"links"`
			Id       string `json:"id"`
			Position struct {
				X string `json:"x"`
				Y string `json:"y"`
			} `json:"position,omitempty"`
		} `json:"endItem"`
		CreatedAt time.Time `json:"createdAt"`
		CreatedBy struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"createdBy"`
		ModifiedAt time.Time `json:"modifiedAt"`
		ModifiedBy struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"modifiedBy"`
		Shape string `json:"shape"`
		Style struct {
			StartStrokeCap string `json:"startStrokeCap"`
			EndStrokeCap   string `json:"endStrokeCap"`
			StrokeWidth    string `json:"strokeWidth"`
			StrokeStyle    string `json:"strokeStyle"`
			StrokeColor    string `json:"strokeColor"`
		} `json:"style"`
		Type string `json:"type"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Type string `json:"type"`
}
type BoardItems struct {
	Size  int `json:"size"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Data  []struct {
		Id    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		CreatedAt time.Time `json:"createdAt"`
		CreatedBy struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"createdBy"`
		Data struct {
			Content string `json:"content"`
			Shape   string `json:"shape"`
		} `json:"data"`
		Geometry struct {
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"geometry"`
		ModifiedAt time.Time `json:"modifiedAt"`
		ModifiedBy struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"modifiedBy"`
		Position struct {
			X          float64 `json:"x"`
			Y          float64 `json:"y"`
			Origin     string  `json:"origin"`
			RelativeTo string  `json:"relativeTo"`
		} `json:"position"`
		Style struct {
			FillColor         string `json:"fillColor"`
			TextAlign         string `json:"textAlign"`
			TextAlignVertical string `json:"textAlignVertical"`
		} `json:"style"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Type string `json:"type"`
}

type ProjectId struct {
	Id string `json:"id"`
}

type Connection struct {
	Managed             bool        `json:"managed"`
	Id                  string      `json:"id"`
	Name                string      `json:"name"`
	Type                string      `json:"type"`
	ProjectId           string      `json:"projectId"`
	ManagedConnectionId interface{} `json:"managedConnectionId"`
	Connectors          []struct {
		Id            string `json:"id"`
		EnvironmentId string `json:"environmentId"`
		ConnectionId  string `json:"connectionId"`
	} `json:"connectors"`
}

type ProjectInfo struct {
	Id               string
	Name             string
	Logo             string
	ProjectId        string
	OrgId            string
	OrgName          string
	EnvironmentProd  string
	EnvironmentCTE   string
	ListenerStatuses map[string]ProjectStatus
}

type ProjectStatus struct {
	EnvironmentId   string `json:"environment_id"`
	VersionId       int    `json:"version_id"`
	NumListeners    int    `json:"num_listeners"`
	Running         bool   `json:"running"`
	StartingGraphId string `json:"starting_graph_id"`
	ProjectId       string `json:"project_id"`
	Id              string `json:"id"`
	PreviousState   string `json:"previous_state"`
	ListenerId      string `json:"listener_id"`
	OverallState    string `json:"overall_state"`
	NumRunning      int    `json:"num_running"`
	EngineGroup     string `json:"engine_group"`
}

type ProjectStatuses struct {
	Data []struct {
		EnvironmentId   string `json:"environment_id"`
		VersionId       int    `json:"version_id"`
		NumListeners    int    `json:"num_listeners"`
		Running         bool   `json:"running"`
		StartingGraphId string `json:"starting_graph_id"`
		ProjectId       string `json:"project_id"`
		Id              string `json:"id"`
		PreviousState   string `json:"previous_state"`
		ListenerId      string `json:"listener_id"`
		OverallState    string `json:"overall_state"`
		NumRunning      int    `json:"num_running"`
		EngineGroup     string `json:"engine_group"`
	} `json:"data"`
	Type string `json:"type"`
}

type Org struct {
	Name     string
	Projects map[string]ProjectInfo
}

func initializeProjectInfos() map[string]ProjectInfo {
	pps := make(map[string]ProjectInfo)

	var pp = ProjectInfo{}
	pp.ProjectId = "2354572809474147774"
	pp.Name = "Healthcare/Pharma"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Kitewheel Presales"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2372660393874032299"
	pp.EnvironmentCTE = "2354572809532868149"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2905614899210094321"
	pp.Name = "Enterprise"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Nationwide"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2911383154977670485"
	pp.EnvironmentCTE = "2911385156549870934"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2841118293232715377"
	pp.Name = "Pay_My_Bill"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Nationwide"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2841119466194994367"
	pp.EnvironmentCTE = "2841119416869979326"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2664772700065239848"
	pp.Name = "Atlantic Broadband Batch"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Atlantic Broadband"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2725602269181252424"
	pp.EnvironmentCTE = "2725602164969575239"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2657694252528308938"
	pp.Name = "Atlantic Broadband Main Controller"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Atlantic Broadband"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2670075426324091305"
	pp.EnvironmentCTE = "2670075249408348583"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2664772699788415782"
	pp.Name = "Atlantic Broadband ENI"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Atlantic Broadband"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2755385570838188990"
	pp.EnvironmentCTE = "2692317366989886098"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2664772699721306917"
	pp.Name = "Atlantic Broadband Vantage"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Atlantic Broadband"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2670076006362780077"
	pp.EnvironmentCTE = "2670075955318100396"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2372778672291319886"
	pp.Name = "Chase Collections"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Chase"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2383042848251187392"
	pp.EnvironmentCTE = "2394043022054265027"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2643800824597975515"
	pp.Name = "Scheduler"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Chase Fraud"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2766279367142937554"
	pp.EnvironmentCTE = "2718696616332366610"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2591767054894502124"
	pp.Name = "SMS Outbound"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Chase Fraud"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2765765576655377361"
	pp.EnvironmentCTE = "2695364350160147130"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2591767638708064497"
	pp.Name = "Chase Notify Outbound"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Chase Fraud"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2626567207840912037"
	pp.EnvironmentCTE = "2626567171920892580"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2582147883969810657"
	pp.Name = "Disputes"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Chase Fraud"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2592272672000578130"
	pp.EnvironmentCTE = "2592272489254752849"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2296784218418319373"
	pp.Name = "Chase Mortgage"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Chase Paid In Full"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2296784218493816853"
	pp.EnvironmentCTE = "2296784409494032406"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	// cofiroute
	pp = ProjectInfo{}
	pp.ProjectId = "2844692586223636850"
	pp.Name = "CTRMA Digital Billing"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Cofiroute"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2845386117275717742"
	pp.EnvironmentCTE = "2844692586273968196"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2942354845229388476"
	pp.Name = "Monitoring"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Cofiroute"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2942360037224092414"
	pp.EnvironmentCTE = "2974480745571816325"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	//exact sciences
	pp = ProjectInfo{}
	pp.ProjectId = "2911430083921842798"
	pp.Name = "Colorectal Cancer Screening"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Exact Sciences"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2946798377306428179"
	pp.EnvironmentCTE = "2946798323636114194"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2947517083233817293"
	pp.Name = "JAM Logging"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Exact Sciences"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2951238388080710432"
	pp.EnvironmentCTE = "2951238323672978207"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2905614899210094321"
	pp.Name = "Enterprise"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Nationwide"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2911383154977670485"
	pp.EnvironmentCTE = "2911385156549870934"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	pp = ProjectInfo{}
	pp.ProjectId = "2841118293232715377"
	pp.Name = "Pay My Bill"
	pp.OrgId = "2063393085278127258"
	pp.OrgName = "Nationwide"
	pp.Logo = "https://cdn.pixabay.com/photo/2012/04/25/08/47/medicine-41712_960_720.png"
	pp.EnvironmentProd = "2841119466194994367"
	pp.EnvironmentCTE = "2841119416869979326"
	pp.ListenerStatuses = make(map[string]ProjectStatus)
	pps[pp.ProjectId] = pp

	return pps
}

// https://hub-us.kitewheel.com/api/0/projects/3013515008937035048/journey-step-nodes

func getProjectStatuses(host string, sessionId string) ProjectStatuses {
	// https://hub-us.kitewheel.com/api/0/engine/graphs
	// projs := ProjectStatuses

	murl := "https://" + host + "/api/0/engine/graphs"

	req, _ := http.NewRequest("GET", murl, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://"+host)

	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	projs := ProjectStatuses{}
	json.Unmarshal([]byte(body), &projs)
	//println(body)
	return projs
}

func createMetric(projectId string, name string, desc string) {
	//https://hub-us.kitewheel.com/api/0/projects/3013515008937035048/metrics
	murl := "https://hub-us.kitewheel.com/api/0/projects/" + projectId + "/metrics"

	projBody := strings.NewReader("{\"projectId\":\"" + projectId + "\",\"name\":\"" + name + "\",\"type\":\"count\",\"description\":\"" + desc + "\",\"goal\":null,\"currencyMultiplier\":1,\"currencyType\":\"dollar\"}")

	req, _ := http.NewRequest("POST", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", "t9X7fU01-Rq7_FqkHivtTfghBdsjA6fRjdRo")
	req.Header.Add("origin", "https://hub-us.kitewheel.com")
	req.Header.Add("origin", "https://hub-us.kitewheel.com")
	req.Header.Add("cookie", "connect.sid=s%3A2SA_wMgG0k-OqVeNr8iZJTCdI_NKbpDI.jhy8MpvqrrCxGfS3m9Odok24Uogk1B1ekN8L8ssFfeE")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	println(body)
}

func createJourneyStep(projectId string, name string, desc string, outcome string, sessionId string, csrfToken string, host string) {
	murl := "https://" + host + "/api/0/projects/" + projectId + "/journey-step-nodes"

	println("Creating journey step: " + name)
	projBody := strings.NewReader("{\"outcome\":\"" + outcome + "\",\"outcomeLabel\":\"" + outcome + "\",\"name\":\"" + name + "\",\"description\":\"" + desc + "\",\"parameters\":{}}")

	req, _ := http.NewRequest("POST", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", csrfToken)
	req.Header.Add("referer", "https://"+host+"/mapper/"+projectId)
	req.Header.Add("authority", host)
	req.Header.Add("origin", "https://"+host)

	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	println(body)

}

func createGraph(projectId string, name string, sessionId string, csrfToken string, host string) {
	murl := "https://" + host + "/api/0/projects/" + projectId + "/node-prototypes"

	println("Creating graph: " + name)
	projBody := strings.NewReader("{\"tags\":[],\"type\":\"graph\",\"name\":\"" + name + "\"}")

	req, _ := http.NewRequest("POST", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", csrfToken)

	req.Header.Add("authority", host)
	req.Header.Add("origin", "https://"+host)
	//req.Header.Add("referer", "https://"+host+"/editor/"+projectId + "/logic/" + )
	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

}

func addNodeToGraph(projectId string, graphId string, prototypeId string, name string, x string, y string, sessionId string, csrfToken string, host string) {
	murl := "https://" + host + "/api/0/node-instances?type=graph"

	println("Creating graph: " + name)
	projBody := strings.NewReader("{\"type\":\"graph\",\"name\":\"Logging\",\"icon\":\"fas fa-cogs\",\"userId\":\"2943034479621181814\",\"projectId\":\"" + projectId + "\",\"channelTypeId\":null,\"editingBy\":null,\"prototypeDeletedAt\":null,\"createdAt\":\"2023-02-21T21:43:22.167Z\",\"updatedAt\":\"2023-02-21T21:43:22.167Z\",\"tags\":[{\"name\":\"graph\",\"color\":\"#a358b7\",\"isDefault\":true}],\"x\":" + x + ",\"y\":" + y + ",\"graphId\":\"" + graphId + "\",\"prototypeId\":\"" + prototypeId + "\"}")

	req, _ := http.NewRequest("POST", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", csrfToken)

	req.Header.Add("authority", host)
	req.Header.Add("origin", "https://"+host)
	//req.Header.Add("referer", "https://"+host+"/editor/"+projectId + "/logic/" + )
	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

}

func createJavaScript(projectId string, name string, code string, sessionId string, csrfToken string, host string) {
	murl := "https://" + host + "/api/0/projects/" + projectId + "/node-prototypes"

	println("Creating javascript: " + name)
	projBody := strings.NewReader("{\"tags\":[],\"type\":\"script\",\"name\":\"" + name + "\"}")

	req, _ := http.NewRequest("POST", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", csrfToken)

	req.Header.Add("authority", host)
	req.Header.Add("origin", "https://"+host)
	//req.Header.Add("referer", "https://"+host+"/editor/"+projectId + "/logic/" + )
	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	id := KWJavaScriptPrototype{}
	json.Unmarshal([]byte(body), &id)

	updateJavaScript(id.Prototype.Id, name, code, sessionId, csrfToken, host)
	println(body)

}

func updateJavaScript(prototypeId string, name string, code string, sessionId string, csrfToken string, host string) {
	murl := "https://" + host + "/api/0/node-prototypes/" + prototypeId

	println("Updating javascript: " + name)
	myBody := "{\"body\":\"" + code + "\"}"
	projBody := strings.NewReader(myBody)

	req, _ := http.NewRequest("PUT", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", csrfToken)

	req.Header.Add("authority", host)
	req.Header.Add("origin", "https://"+host)

	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	id := KWPrototype{}
	json.Unmarshal([]byte(body), &id)

	println(body)

}

func addJourneyNode(journeyId string, node string) {
	// https://hub-us.kitewheel.com/api/0/node-instances?type=journey
	murl := "https://hub-us.kitewheel.com/api/0/node-instances?type=journey"

	projBody := strings.NewReader("{\"x\":329,\"y\":355,\"type\":\"ghost\",\"backgroundColor\":\"#a81879\",\"iconColor\":\"#ffffff\",\"icon\":\"fas fa-users\",\"graphId\":\"" + journeyId + "\",\"prototypeId\":1}")

	req, _ := http.NewRequest("POST", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", "t9X7fU01-Rq7_FqkHivtTfghBdsjA6fRjdRo")
	req.Header.Add("origin", "https://hub-us.kitewheel.com")
	req.Header.Add("origin", "https://hub-us.kitewheel.com")
	req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
	req.Header.Add("cookie", "connect.sid=s%3A2SA_wMgG0k-OqVeNr8iZJTCdI_NKbpDI.jhy8MpvqrrCxGfS3m9Odok24Uogk1B1ekN8L8ssFfeE")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	println(body)
}

func createJourneyMap(projectId string, name string) {
	murl := "https://hub-us.kitewheel.com/api/0/projects/" + projectId + "/node-prototypes"

	projBody := strings.NewReader("{\"type\":\"journey\",\"name\":\"" + name + "\",\"description\":\"\",\"tags\":[]}")

	req, _ := http.NewRequest("POST", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", "t9X7fU01-Rq7_FqkHivtTfghBdsjA6fRjdRo")
	req.Header.Add("origin", "https://hub-us.kitewheel.com")
	req.Header.Add("origin", "https://hub-us.kitewheel.com")
	req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
	req.Header.Add("cookie", "connect.sid=s%3A2SA_wMgG0k-OqVeNr8iZJTCdI_NKbpDI.jhy8MpvqrrCxGfS3m9Odok24Uogk1B1ekN8L8ssFfeE")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	println(body)
}

func CreateProject(name string, description string, sessionId string, csrfToken string, host string) string {

	println("Creating new project " + name)
	murl := host + "/api/0/projects"

	projBody := strings.NewReader("{\"name\":\"" + name + "\",\"description\":\"" + description + "\",\"invalidInput\":[],\"selected\":{\"id\":\"1\",\"name\":\"Kitewheel\",\"hierarchy\":{\"organizationId\":\"1\",\"parents\":[],\"allDescendents\":[\"2124992990718985395\",\"1835831493516067943\",\"2322145224303838439\",\"2322736559620621545\",\"2613361665051198744\",\"2621964777013904666\",\"1835831154943460452\",\"1835831406266156134\",\"1835831655693026408\",\"1817405784846238814\",\"2510596089799771404\",\"2181460569423348927\",\"1835831301425333349\",\"2550829615333508370\",\"2809147778456880420\",\"1683579969785037891\",\"2600849731571680534\",\"2809147866595984677\",\"2809066276956144931\",\"1703186918570722374\",\"12\",\"2488028352359695624\",\"1852370781288268907\",\"2886747964884124970\",\"1816815216441164888\",\"2329800296575796458\",\"2283395180952290523\",\"2887399165124412715\",\"2342197734137660652\",\"3009768281960613083\",\"2996949223859029209\",\"2994054605316293848\",\"2550299151455749393\",\"2074209544312980637\",\"2796707919771469090\",\"1969339376450143365\",\"2733839075831186718\",\"2401011735247979775\",\"2284497976770430172\",\"2227141035505484998\",\"2116211057591583920\",\"2063393085278127258\",\"2774326590975247648\",\"2056157604056401047\",\"2896731500944819500\",\"2861956944996861224\",\"2235836434160288970\",\"2719311071516034333\",\"2616265122947007769\",\"2547285480680981774\",\"2388025253407229179\",\"2322145071706670310\",\"2517107438746862861\",\"2360512674527184111\",\"2829954022151030054\",\"2866433638868714793\",\"2937939208460829911\",\"2906359125514388781\",\"2654701749826225435\",\"2014696495852291215\",\"2263435265471153362\",\"2998946819976725722\",\"1835831776010830953\"],\"directDescendents\":[\"1703186918570722374\",\"2263435265471153362\",\"1817405784846238814\",\"12\",\"2488028352359695624\",\"2283395180952290523\",\"2510596089799771404\",\"2550829615333508370\",\"1683579969785037891\",\"2600849731571680534\",\"1852370781288268907\",\"2886747964884124970\",\"1816815216441164888\",\"2342197734137660652\",\"2550299151455749393\",\"2074209544312980637\",\"1969339376450143365\",\"2401011735247979775\",\"2284497976770430172\",\"2227141035505484998\",\"2116211057591583920\",\"2063393085278127258\",\"2056157604056401047\",\"2861956944996861224\",\"2719311071516034333\",\"2547285480680981774\",\"2829954022151030054\",\"2866433638868714793\",\"2937939208460829911\",\"2906359125514388781\",\"2654701749826225435\",\"2998946819976725722\"],\"parentNames\":[]}},\"organizationId\":\"1\"}")

	req, _ := http.NewRequest("POST", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", csrfToken)
	req.Header.Add("origin", csrfToken)

	req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	projectId := ProjectId{}
	json.Unmarshal([]byte(body), &projectId)

	println(body)
	println(projectId.Id)
	return projectId.Id
}

func cloneConnection(kwc KWConnections, projectId string, sessionId string, csrfToken string, host string) {

	for i, v := range kwc.Connectors {
		fmt.Println(i, v)
		println("Creating new SQS connection")
		murl := host + "/api/0/projects/" + projectId + "/connections"

		projBody := strings.NewReader("{\"name\":\"" + v.ConnectionDetails.ConnectorRequiredName + "\",\"type\":\"queue\"}")

		req, _ := http.NewRequest("POST", murl, projBody)

		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")
		req.Header.Add("csrf-token", csrfToken)
		req.Header.Add("origin", host)

		req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
		req.Header.Add("cookie", "connect.sid="+sessionId)
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		id := Connection{}
		json.Unmarshal([]byte(body), &id)

		// https://hub-us.kitewheel.com/api/0/connectors/3024357318918669839
		createConnectionDetails("xponent SQS", v.ConnectionDetails.Username, v.ConnectionDetails.Password, v.ReadQueue, v.WriteQueue, id.Connectors[0].Id, sessionId, csrfToken, host)

		//println(body)

	}

}

func CreateConnection(name string, key string, secret string, readQueue string, writeQueue string, projectId string, sessionId string, csrfToken string, host string) {

	println("Creating new SQS connection")
	murl := host + "/api/0/projects/" + projectId + "/connections"

	projBody := strings.NewReader("{\"name\":\"" + name + "\",\"type\":\"queue\"}")

	req, _ := http.NewRequest("POST", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", csrfToken)
	req.Header.Add("origin", host)

	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	id := Connection{}
	json.Unmarshal([]byte(body), &id)

	// https://hub-us.kitewheel.com/api/0/connectors/3024357318918669839
	createConnectionDetails(name, key, secret, readQueue, writeQueue, id.Connectors[0].Id, sessionId, csrfToken, host)

	//println(body)
}

func createConnectionDetails(name string, key string, secret string, readQueue string, writeQueue string, connectorId string, sessionId string, csrfToken string, host string) {

	println("Setting SQS details")
	murl := host + "/api/0/connectors/" + connectorId

	projBody := strings.NewReader("{\"readQueue\":\"" + readQueue + "\",\"writeQueue\":\"" + writeQueue + "\",\"awsAccessKeyId\":\"" + key + "\",\"awsSecretAccessKey\":\"" + secret + "\",\"awsRegion\":\"us-east-1\",\"queueType\":\"sqs\"}")

	req, _ := http.NewRequest("PUT", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("csrf-token", csrfToken)
	req.Header.Add("origin", host)

	req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//id := ProjectId{}
	//json.Unmarshal([]byte(body), &id)

	// https://hub-us.kitewheel.com/api/0/connectors/3024357318918669839

	println(body)
}

func getConnections(projectId string, host string, sessionId string) []KWConnections {
	println("Getting connections for project")
	murl := "https://" + host + "/api/0/projects/" + projectId + "/connections?version=null"

	projBody := strings.NewReader("{\"name\":\"a\",\"type\":\"queue\"}")

	req, _ := http.NewRequest("GET", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	//req.Header.Add("csrf-token", csrfToken)
	req.Header.Add("origin", host)

	req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//mys := string(body[:])

	id := []KWConnections{}
	json.Unmarshal(body, &id)

	// https://hub-us.kitewheel.com/api/0/connectors/3024357318918669839
	//createConnectionDetails("xponent SQS", key, secret, readQueue, writeQueue, id.Connectors[0].Id, sessionId, csrfToken, host)

	//println(mys)
	return id
}

func getGraph(scriptId string, host string, sessionId string) interface{} {
	println("Getting JavaScript...")
	murl := "https://" + host + "/api/0/node-prototypes/" + scriptId + "?type=graph"

	projBody := strings.NewReader("{\"name\":\"a\",\"type\":\"queue\"}")

	req, _ := http.NewRequest("GET", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	//req.Header.Add("csrf-token", csrfToken)
	req.Header.Add("origin", host)

	req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//mys := string(body[:])

	var id interface{}
	json.Unmarshal(body, &id)

	// https://hub-us.kitewheel.com/api/0/connectors/3024357318918669839
	//createConnectionDetails("xponent SQS", key, secret, readQueue, writeQueue, id.Connectors[0].Id, sessionId, csrfToken, host)

	//println(mys)
	return id
}

func getScript(scriptId string, host string, sessionId string) Script {
	println("Getting JavaScript...")
	murl := "https://" + host + "/api/0/node-prototypes/" + scriptId + "?type=graph"

	projBody := strings.NewReader("{\"name\":\"a\",\"type\":\"queue\"}")

	req, _ := http.NewRequest("GET", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	//req.Header.Add("csrf-token", csrfToken)
	req.Header.Add("origin", host)

	req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//mys := string(body[:])

	id := Script{}
	json.Unmarshal(body, &id)

	// https://hub-us.kitewheel.com/api/0/connectors/3024357318918669839
	//createConnectionDetails("xponent SQS", key, secret, readQueue, writeQueue, id.Connectors[0].Id, sessionId, csrfToken, host)

	//println(mys)
	return id
}

func getProject(projectId string, host string, sessionId string) KWProject {
	println("Getting project details for project")
	murl := "https://" + host + "/api/0/projects/" + projectId + "?type=refresh&prototype=open&mode=project"

	projBody := strings.NewReader("{\"name\":\"a\",\"type\":\"queue\"}")

	req, _ := http.NewRequest("GET", murl, projBody)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	//req.Header.Add("csrf-token", csrfToken)
	req.Header.Add("origin", host)

	req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
	req.Header.Add("cookie", "connect.sid="+sessionId)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//mys := string(body[:])

	id := KWProject{}
	json.Unmarshal(body, &id)

	// https://hub-us.kitewheel.com/api/0/connectors/3024357318918669839
	//createConnectionDetails("xponent SQS", key, secret, readQueue, writeQueue, id.Connectors[0].Id, sessionId, csrfToken, host)

	//println(mys)
	return id
}

func getBoardItems(boardId string) {
	// uXjVPzLdZ24=
	//murl := "https://api.miro.com/v2/boards/" + url.QueryEscape(boardId) + "/items?limit=10"

	murl := "https://api.miro.com/v2/boards/" + url.QueryEscape(boardId) + "/connectors?limit=10"

	req, _ := http.NewRequest("GET", murl, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	return
}

func main() {
	println("Xponent Development Kit (XDK) Beta 2023")
	//println("Telemetry")
	println(" ")
	//println("Checking customer projects...")
	println("")
	//fuse := fuse.Fuse{}
	//fuse.Parse(data)

	// 2106068342694479693
	// 2911430083921842798
	// 3024332494335644559 john test

	/*if os.Args[1] == "generate-fuse-from-project" {
		println("Creating Fuse file")
		generateFuseFromProject(os.Args[2], os.Args[3], "s%3AOk26kKbFTwA96v5xodwlz_oJEU5NDgoS.ll9u1vSd48M2qrgQcs9bYZUCvDm2WiEkQu66xiqQ4mU")
	}

	if os.Args[1] == "generate-project-from-fuse" {
		_, done := generateFromFuse(os.Args[2], os.Args[3], "s%3AOk26kKbFTwA96v5xodwlz_oJEU5NDgoS.ll9u1vSd48M2qrgQcs9bYZUCvDm2WiEkQu66xiqQ4mU", "IpRpdbyn-N9FYAnZRxwUoEnvCVvFIWXm-iy8")
		if done {
			return
		}
	}*/
	//generateFuseFromProject("exactssciences.fuse", "2911430083921842798", "s%3AOk26kKbFTwA96v5xodwlz_oJEU5NDgoS.ll9u1vSd48M2qrgQcs9bYZUCvDm2WiEkQu66xiqQ4mU")
	//generateFromFuse("exactsciences.fuse",)
	// "2911430083921842798"

	// c, _ := kwclient.NewClient("http://localhost:8001")

	//id := createProject("Vinay Go Test", "This is an automated Ignition Switch project. Built from Spark code!", "s:GJBSBi9LfBfgvpfu50n4er5Z6uJF3Ugo.oZznxtZcRZk7lOgMPcSD1gr4VXb7v2VtBOgEIrgxgpk", "sorUNEmR-lmVYfwjOZyWZGBsfQSZVV9fSTx8", "https://hub.csgjourney.com")
	// createConnection("Xponent Ingress", "xxx", "yyy", "Read queue", "write queue", id, "s:GJBSBi9LfBfgvpfu50n4er5Z6uJF3Ugo.oZznxtZcRZk7lOgMPcSD1gr4VXb7v2VtBOgEIrgxgpk", "sorUNEmR-lmVYfwjOZyWZGBsfQSZVV9fSTx8", "https://hub.csgjourney.com")

}

func generateFuseFromProject(fuseName string, projId string, sess string) {
	println("Generating Fuse file...")
	kwconns := getConnections(projId, "hub.csgjourney.com", sess)

	fuse := Fuse2{}
	fuse.Connections = kwconns
	fuse.Name = "Exact Sciences"

	kwproj := KWProject{}
	kwproj = getProject(projId, "hub.csgjourney.com", "s%3AOk26kKbFTwA96v5xodwlz_oJEU5NDgoS.ll9u1vSd48M2qrgQcs9bYZUCvDm2WiEkQu66xiqQ4mU")
	println("Project: " + fuse.Name)
	println(kwproj.Prototypes[0].Type)

	for _, v := range kwproj.Prototypes {
		if v.Type == "journeyStep" {
			println(v.Type)
			proto := KWPrototype{}
			proto = v
			fuse2 := append(fuse.JourneySteps, proto)
			fuse.JourneySteps = fuse2

		}
		if v.Type == "graph" {
			println(v.Type)
			proto := KWPrototype{}
			v.Graph = getGraph(v.Id, "hub.csgjourney.com", sess)
			proto = v
			fuse2 := append(fuse.Graphs, proto)
			fuse.Graphs = fuse2
		}
		if v.Type == "script" {
			println(v.Type)
			proto := KWPrototype{}
			v.Code = getScript(v.Id, "hub.csgjourney.com", sess)
			proto = v

			fuse2 := append(fuse.JavaScripts, proto)
			fuse.JavaScripts = fuse2
		}

	}

	file, err := os.OpenFile(fuseName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)

	err = enc.Encode(fuse)
	if err != nil {
		log.Fatalf("error encoding: %v", err)
	}
}

func generateFromFuse(fuseFileLoc string, intoProject string, sess string, csrf string) (error, bool) {
	println("Generating project from Fuse: " + fuseFileLoc)
	data, err := ioutil.ReadFile(fuseFileLoc)
	if err != nil {
		fmt.Println(err)
		return nil, true
	}
	//println(data)

	//enc := yaml.NewEncoder(data)
	host := "hub.csgjourney.com"
	destProj := intoProject //"3043337359372063662"
	sessionId := sess       //"s%3AOk26kKbFTwA96v5xodwlz_oJEU5NDgoS.ll9u1vSd48M2qrgQcs9bYZUCvDm2WiEkQu66xiqQ4mU"
	csrfToken := csrf       //"IpRpdbyn-N9FYAnZRxwUoEnvCVvFIWXm-iy8"
	fuseIn := Fuse2{}
	err = yaml.Unmarshal(data, &fuseIn)
	if err != nil {
		log.Fatalf("error encoding: %v", err)
	}
	println(fuseIn.Name)
	for _, o := range fuseIn.JavaScripts {
		println(o.Name)
		createJavaScript(destProj, o.Code.Prototype.Name, o.Code.Type.Body, sessionId, csrfToken, host)
		//createJourneyStep("3024332494335644559", o.Name, o.Name, o.Outcome, "s%3AOk26kKbFTwA96v5xodwlz_oJEU5NDgoS.ll9u1vSd48M2qrgQcs9bYZUCvDm2WiEkQu66xiqQ4mU", "DO2QUnJQ-eC7WRsE17qkSr4SVyc6osVtQMhI", "hub.csgjourney.com")
	}
	for _, o := range fuseIn.Graphs {
		println(o.Name)
		createGraph(destProj, o.Name, sessionId, csrfToken, host)
		//createJourneyStep("3024332494335644559", o.Name, o.Name, o.Outcome, "s%3AOk26kKbFTwA96v5xodwlz_oJEU5NDgoS.ll9u1vSd48M2qrgQcs9bYZUCvDm2WiEkQu66xiqQ4mU", "DO2QUnJQ-eC7WRsE17qkSr4SVyc6osVtQMhI", "hub.csgjourney.com")
	}
	for _, o := range fuseIn.JourneySteps {
		println(o.Name)
		createJourneyStep(destProj, o.Name, o.Name, o.Outcome, sessionId, csrfToken, host)
	}
	return err, false
}

/*println("Applying connections...")
for _, conn := range fuse.Resources.Connections {
	println("Connection")

	println(conn.Connection.Name)
	println(conn.Connection.Type)
	println("")
}

println("Applying JavaScript functions...")
for _, conn := range fuse.Resources.Javascripts {
	println("JavaScript")

	println(conn.Javascript.Name)
	println(conn.Javascript.Code)
	println("")
} */

/*orgs := make(map[string]Org)

projectInfos := initializeProjectInfos()
projects := getProjectStatuses("hub.csgjourney.com", "s%3Ak5pN7lr7xeh1e6t2i5UxuQ-wZZHmOu5Z.ky8PzA0CP6sxx8DYHyI3ob9vZRebVRjF7zJv4%2BzfrtA")
pi := ProjectInfo{}

for _, o := range projects.Data {
	//println(i)
	pi = projectInfos[o.ProjectId]
	if pi.ProjectId == "" {
		continue
	}
	pi.ListenerStatuses[o.Id] = o

	org := orgs[pi.OrgName]
	if org.Name == "" {
		org.Name = pi.OrgName
		org.Projects = make(map[string]ProjectInfo)
	}

	org.Projects[pi.ProjectId] = pi

	orgs[pi.OrgName] = org
	println(pi.OrgName)
	println(pi.Name)
	n := int64(o.NumListeners)
	switch o.EnvironmentId {
	case pi.EnvironmentProd:
		{
			println("Starting graph: https://hub.csgjourney.com/editor/" + pi.ProjectId + "/graph/" + o.StartingGraphId)

			println("# of listeners: " + strconv.FormatInt(n, 10))
			println("Production is " + o.OverallState)
		}
	case pi.EnvironmentCTE:
		{
			println("Starting graph: https://hub.csgjourney.com/editor/" + pi.ProjectId + "/graph/" + o.StartingGraphId)
			println("# of listeners: " + strconv.FormatInt(n, 10))
			println("CTE is " + o.OverallState)
		}
	}
	println("")*/

//}

// https://hub-us.kitewheel.com/editor/2734404962711766489/graph/2734425319011255851
/*
	projects = getProjectStatuses("hub-us.kitewheel.com", "s%3Ae6NHx6nWwAp4YtUwtzwFLo_X_5Lw7_hT.lexWz9Oc%2BNHgxiEb95Lw2uXPWZMGUEcKi0ue5JzpXvk")
	pi = ProjectInfo{}
	for _, o := range projects.Data {
		//println(i)
		pi = projectInfos[o.ProjectId]
		if pi.ProjectId == "" {
			continue
		}
		pi.ListenerStatuses[o.Id] = o

		org := orgs[pi.OrgName]
		if org.Name == "" {
			org.Name = pi.OrgName
			org.Projects = make(map[string]ProjectInfo)
		}

		org.Projects[pi.ProjectId] = pi

		orgs[pi.OrgName] = org
		println(pi.OrgName)
		println(pi.Name)
		n := int64(o.NumListeners)
		switch o.EnvironmentId {
		case pi.EnvironmentProd:
			{
				println("Starting graph: https://hub-us.kitewheel.com/editor/" + pi.ProjectId + "/graph/" + o.StartingGraphId)
				println("# of listeners: " + strconv.FormatInt(n, 10))
				println("Production is " + o.OverallState)
			}
		case pi.EnvironmentCTE:
			{
				println("Starting graph: https://hub-us.kitewheel.com/editor/" + pi.ProjectId + "/graph/" + o.StartingGraphId)
				println("# of listeners: " + strconv.FormatInt(n, 10))
				println("CTE is " + o.OverallState)
			}
		}
		println("")

	}*/

//createMetric("3013515008937035048", "John mtriccc", "testing metric")
//createJourneyStep("3013515008937035048", "John steperoo", "big step", "Positive")
//addJourneyNode("3013525718211495503", "hello")
//createJourneyMap("3013515008937035048", "brilliant")
//id := createProject("Ignition Switch", "This is an automated Ignition Switch project. Built from Spark code!", "s%3A6lNs2LQB6WeZ5gKUTswkrHEVbuyV_ynL.jObQTi6HsYwwZHfCcAIUqfqEauWeoGLFgHkvafhBi84", "LezAD4V1-VCrslMLU-w-9yd-VnjhYDz-E3i8", "https://hub-us.kitewheel.com")
//createConnection("Xponent Ingress", "xxx", "yyy", "Read queue", "write queue", id, "s%3A6lNs2LQB6WeZ5gKUTswkrHEVbuyV_ynL.jObQTi6HsYwwZHfCcAIUqfqEauWeoGLFgHkvafhBi84", "LezAD4V1-VCrslMLU-w-9yd-VnjhYDz-E3i8", "https://hub-us.kitewheel.com")
//getBoardItems("uXjVPzLdZ24=")
/*type Board = struct {
      Id          string `json:"id"`
      Type        string `json:"type"`
      Name        string `json:"name"`
      Description string `json:"description"`
      Links      struct {
        Self    string `json:"self"`
        Related string `json:"related"`
      } `json:"links"`
      CreatedAt time.Time `json:"createdAt"`
      CreatedBy struct {
        Id  string `json:"id"`
        Type string `json:"type"`
        Name string `json:"name"`
      } `json:"createdBy"`
      CurrentUserMembership struct {
        Id  string `json:"id"`
        Type string `json:"type"`
        Name string `json:"name"`
        Role string `json:"role"`
      } `json:"currentUserMembership"`
      ModifiedAt time.Time `json:"modifiedAt"`
      ModifiedBy struct {
        Id  string `json:"id"`
        Type string `json:"type"`
        Name string `json:"name"`
      } `json:"modifiedBy"`
      Owner struct {
        Id  string `json:"id"`
        Type string `json:"type"`
        Name string `json:"name"`
      } `json:"owner"`
      PermissionsPolicy struct {
        CollaborationToolsStartAccess string `json:"collaborationToolsStartAccess"`
        CopyAccess                    string `json:"copyAccess"`
        CopyAccessLevel              string `json:"copyAccessLevel"`
        SharingAccess                string `json:"sharingAccess"`
      } `json:"permissionsPolicy"`
      Policy struct {
        PermissionsPolicy struct {
            CollaborationToolsStartAccess string `json:"collaborationToolsStartAccess"`
            CopyAccess                    string `json:"copyAccess"`
            CopyAccessLevel              string `json:"copyAccessLevel"`
            SharingAccess                string `json:"sharingAccess"`
        } `json:"permissionsPolicy"`
        SharingPolicy struct {
            Access                            string `json:"access"`
            InviteToAccountAndBoardLinkAccess string `json:"inviteToAccountAndBoardLinkAccess"`
            OrganizationAccess                string `json:"organizationAccess"`
            TeamAccess                        string `json:"teamAccess"`
        } `json:"sharingPolicy"`
      } `json:"policy"`
      SharingPolicy struct {
        Access                            string `json:"access"`
        InviteToAccountAndBoardLinkAccess string `json:"inviteToAccountAndBoardLinkAccess"`
        OrganizationAccess                string `json:"organizationAccess"`
        TeamAccess                        string `json:"teamAccess"`
      } `json:"sharingPolicy"`
      Team struct {
        Id  string `json:"id"`
        Type string `json:"type"`
        Name string `json:"name"`
      } `json:"team"`
      ViewLink string `json:"viewLink"`
  }
  board := Board{}
  murl := "https://api.miro.com/v2/boards"

  payload := strings.NewReader("{\"name\":\"Untitled\",\"policy\":{\"permissionsPolicy\":{\"collaborationToolsStartAccess\":\"all_editors\",\"copyAccess\":\"anyone\",\"sharingAccess\":\"team_members_with_editing_rights\"},\"sharingPolicy\":{\"access\":\"private\",\"inviteToAccountAndBoardLinkAccess\":\"no_access\",\"organizationAccess\":\"private\",\"teamAccess\":\"private\"}}}")

  req, _ := http.NewRequest("POST", murl, payload)

  req.Header.Add("accept", "application/json")
  req.Header.Add("content-type", "application/json")
  req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
  json.Unmarshal([]byte(body), &board)

  println(board.Id)

  murl = "https://api.miro.com/v2/boards/" + url.QueryEscape(board.Id) + "/sticky_notes"

  payload = strings.NewReader("{\"data\":{\"content\":\"Hello\",\"shape\":\"square\"},\"position\":{\"origin\":\"center\",\"x\":0,\"y\":0}}")
  req, _ = http.NewRequest("POST", murl, payload)
  req.Header.Add("accept", "application/json")
  req.Header.Add("content-type", "application/json")
  req.Header.Add("authorization", "Bearer eyJtaXJvLm9yaWdpbiI6ImV1MDEifQ_lhUmrwicThOlAWFmhFbYipOFV-s")
  res, _ = http.DefaultClient.Do(req)
  body, _ = ioutil.ReadAll(res.Body)
  println(body)*/

/*
   {
     data: {
       content: "I'm a sticky note",
       shape: "square",
     },
     style: {
       fillColor: "light_yellow",
       textAlign: "center",
       textAlignVertical: "top",
     },
     position: {
       x: 0,
       y: 0,
     },
   }
*/

//}
