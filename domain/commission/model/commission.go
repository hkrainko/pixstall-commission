package model

import (
	model2 "pixstall-commission/domain/file/model"
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

	Price                          Price        `json:"price" bson:"price"`
	DayNeed                        int          `json:"dayNeed" bson:"dayNeed"`
	Size                           *model2.Size `json:"size" bson:"size,omitempty"`
	Resolution                     *float64     `json:"resolution" bson:"resolution,omitempty"`
	ExportFormat                   *string      `json:"exportFormat" bson:"exportFormat,omitempty"`
	Desc                           string       `json:"desc" bson:"desc"`
	PaymentMethod                  string       `json:"paymentMethod" bson:"paymentMethod"`
	IsR18                          bool         `json:"isR18" bson:"isR18"`
	BePrivate                      bool         `json:"bePrivate" bson:"bePrivate"`
	Anonymous                      bool         `json:"anonymous" bson:"anonymous"`
	RefImagePaths                  []string     `json:"refImagePaths" bson:"refImagePaths"`
	TimesAllowedDraftToChange      *int         `json:"timesAllowedDraftToChange" bson:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int         `json:"timesAllowedCompletionToChange" bson:"timesAllowedCompletionToChange"`
	DraftChangingRequestTime       int          `json:"draftChangingRequestTime" bson:"draftChangingRequestTime"`
	ProofCopyRevisionRequestTime   int          `json:"proofCopyRevisionRequestTime" bson:"proofCopyRevisionRequestTime"`

	Messages []model.Messaging `json:"messages" bson:"-"`

	ProofCopyImagePaths []string `json:"proofCopyImagePaths" bson:"proofCopyImagePaths"`

	StartTime     *time.Time    `json:"startTime,omitempty" bson:"startTime,omitempty"`
	CompletedTime *time.Time    `json:"completedTime,omitempty" bson:"completedTime,omitempty"`
	DisplayImage  *DisplayImage `json:"displayImage,omitempty" bson:"displayImage,omitempty"`

	CompletionFilePath *string `json:"completionFilePath,omitempty" bson:"completionFilePath,omitempty"`
	Rating             *int    `json:"rating,omitempty" bson:"rating,omitempty"`
	Comment            *string `json:"comment,omitempty" bson:"comment,omitempty"`

	CreateTime        time.Time              `json:"createTime" bson:"createTime"`
	LastUpdateTime    time.Time              `json:"lastUpdateTime" bson:"lastUpdateTime"`
	ValidationHistory []CommissionValidation `bson:"validationHistory"`
	State             CommissionState        `json:"state" bson:"state"`
}

type DisplayImage struct {
	Path        string      `json:"path" bson:"path"`
	Volume      int64       `json:"volume" bson:"volume"`
	Size        model2.Size `json:"size" bson:"size"`
	ContentType string      `json:"contentType" bson:"contentType"`
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
