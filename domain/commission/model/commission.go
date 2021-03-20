package model

import (
	"pixstall-commission/domain/message/model"
	"time"
)

type Commission struct {
	ID               string `json:"id" bson:"id"`
	OpenCommissionID string `json:"openCommissionId" bson:"openCommissionId"`

	ArtistID             string  `json:"artistId" bson:"artistId"`
	ArtistName           string  `json:"artistName" bson:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath" bson:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId" bson:"requesterId"`
	RequesterName        string  `json:"requesterName" bson:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath" bson:"requesterProfilePath,omitempty"`

	Price                          Price    `json:"price" bson:"price"`
	DayNeed                        int      `json:"dayNeed" bson:"dayNeed"`
	Size                           *Size    `json:"size" bson:"size,omitempty"`
	Resolution                     *float64 `json:"resolution" bson:"resolution,omitempty"`
	ExportFormat                   *string  `json:"exportFormat" bson:"exportFormat,omitempty"`
	Desc                           string   `json:"desc" bson:"desc"`
	PaymentMethod                  string   `json:"paymentMethod" bson:"paymentMethod"`
	IsR18                          bool     `json:"isR18" bson:"isR18"`
	BePrivate                      bool     `json:"bePrivate" bson:"bePrivate"`
	Anonymous                      bool     `json:"anonymous" bson:"anonymous"`
	RefImagePaths                  []string `json:"refImagePaths" bson:"refImagePaths"`
	TimesAllowedDraftToChange      *int     `json:"timesAllowedDraftToChange" bson:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int     `json:"timesAllowedCompletionToChange" bson:"timesAllowedCompletionToChange"`
	DraftChangingRequestTime       int      `json:"draftChangingRequestTime" bson:"draftChangingRequestTime"`
	CompletionRevisionRequestTime  int      `json:"completionRevisionRequestTime" bson:"completionRevisionRequestTime"`

	Messages []model.Messaging `json:"messages" bson:"-"`

	CreateTime        time.Time              `json:"createTime" bson:"createTime"`
	CompleteTime      *time.Time             `json:"completeTime" bson:"completeTime,omitempty"`
	LastUpdateTime    time.Time              `json:"lastUpdateTime" bson:"lastUpdateTime"`
	ValidationHistory []CommissionValidation `bson:"validationHistory"`
	State             CommissionState        `json:"state" bson:"state"`
}

type CommissionState string

const (
	CommissionStatePendingValidation                       CommissionState = "PENDING_VALIDATION"
	CommissionStateInvalidatedDueToOpenCommission          CommissionState = "INVALIDATED_DUE_TO_OPEN_COMMISSION"
	CommissionStateInvalidatedDueToUsers                   CommissionState = "INVALIDATED_DUE_TO_USERS"
	CommissionStatePendingArtistApproval                   CommissionState = "PENDING_ARTIST_APPROVAL"
	CommissionStatePendingRequesterModificationValidation  CommissionState = "PENDING_REQUESTER_MODIFICATION_VALIDATION"
	CommissionStateInProgress                              CommissionState = "IN_PROGRESS"
	CommissionStatePendingRequesterAcceptance              CommissionState = "PENDING_REQUESTER_ACCEPTANCE"
	CommissionStateDeclinedByArtist                        CommissionState = "DECLINED_BY_ARTIST"
	CommissionStateCancelledByRequester                    CommissionState = "CANCELED_BY_REQUESTER"
	CommissionStatePendingUploadProduct                    CommissionState = "PENDING_UPLOAD_PRODUCT"
	CommissionStatePendingUploadProductDueToRevisionExceed CommissionState = "PENDING_UPLOAD_PRODUCT_DUE_TO_REVISION_EXCEED"
	CommissionStatePendingRequesterAcceptProduct           CommissionState = "PENDING_REQUESTER_ACCEPT_PRODUCT"
	CommissionStateCompleted                               CommissionState = "COMPLETED"
)

type CommissionDecision string

const (
	CommissionDecisionRequesterModify          CommissionDecision = "REQUESTER_MODIFY"
	CommissionDecisionArtistAccept             CommissionDecision = "ARTIST_ACCEPT"
	CommissionDecisionArtistDecline            CommissionDecision = "ARTIST_DECLINE"
	CommissionDecisionRequesterCancel          CommissionDecision = "REQUESTER_CANCEL"
	CommissionDecisionArtistUploadProofCopy    CommissionDecision = "ARTIST_UPLOAD_PROOF_COPY"
	CommissionDecisionRequesterAcceptProofCopy CommissionDecision = "REQUESTER_ACCEPT_PROOF_COPY"
	CommissionDecisionRequesterRequestRevision CommissionDecision = "REQUESTER_REQUEST_REVISION"
	CommissionDecisionArtistUploadProduct      CommissionDecision = "ARTIST_UPLOAD_PRODUCT"
	CommissionDecisionRequesterAcceptProduct   CommissionDecision = "REQUESTER_ACCEPT_PRODUCT"
)

type CommissionValidation string

const (
	CommissionValidationOpenCommission CommissionValidation = "COMM_VALIDATION_OPEN_COMM"
	CommissionValidationUsers          CommissionValidation = "COMM_VALIDATION_USERS"
)
